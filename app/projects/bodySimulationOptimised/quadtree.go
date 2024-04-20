package simulationOptimised

import (
	"fmt"
	"slices"

	"github.com/RugiSerl/physics/app/Systems"
	"github.com/RugiSerl/physics/app/math"
	m "github.com/RugiSerl/physics/app/math"
	"github.com/RugiSerl/physics/app/physicUnit"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	APPROXIMATION_RATE = 1.0
)

type TreeType int

const (
	Node  TreeType = iota // The tree contains subregions with bodies, but no body
	Leaf                  // The tree contains only one body but no bodies in subregions
	Empty                 // The body does not contain anything
)

type TreeContent struct {
	treeType    TreeType
	Initialized bool               // keep track if the tree has been created in NewQuadTree()
	body        *Systems.Body      // body pointer, if the tree is a leaf
	ForceOnBody physicUnit.Force2D // forced applied on the body
	MassCenter  math.Vec2          // Mass center of the children, for approximation
	Mass        float64            //total mass of childrens
	Region      math.Rect          // region the tree is covering
}

type QuadTree struct {
	Data     TreeContent
	Children []*QuadTree
}

// Initialize an empty tree
func NewQuadTree(region math.Rect) *QuadTree {
	return &QuadTree{
		Data: TreeContent{
			treeType:    Empty,
			body:        nil,
			Initialized: true,
			Region:      region,
		},
		Children: make([]*QuadTree, 4),
	}
}

//----------------------------------------------------------------------------------------
// INSERT FUNCTIONS ----------------------------------------------------------------------

// Insert a new element in the tree
func (q *QuadTree) Insert(b *Systems.Body) {
	if !fitsIn(b.Position, q.Data.Region) { // don't bother when the body isn't in the tree region
		return
	}
	switch q.Data.treeType {
	case Empty:

		q.Data.body = b // no problem here
		q.Data.treeType = Leaf
	case Leaf:

		if b.Position == q.Data.body.Position { // we don't need that case
			return
		}
		// divide the quadtree to contain both bodies
		sub1 := selectSubtree(q.Data.body.Position, q.Data.Region)
		sub2 := selectSubtree(b.Position, q.Data.Region)

		if sub1 == sub2 { // don't initializate two times if it is in the same children
			q.Children[sub1] = NewQuadTree(subRegion(q.Data.Region, sub1))
			q.Children[sub1].Insert(q.Data.body)
			q.Children[sub1].Insert(b)

		} else { // same thing as previous case, but initializing two different childrens
			q.Children[sub1] = NewQuadTree(subRegion(q.Data.Region, sub1))
			q.Children[sub1].Insert(q.Data.body)

			q.Children[sub2] = NewQuadTree(subRegion(q.Data.Region, sub2))
			q.Children[sub2].Insert(b)
		}

		q.Data.treeType = Node
		q.Data.body = nil // get rid of the body (wait this is weird out of context)
	case Node:
		sub := selectSubtree(b.Position, q.Data.Region)
		if q.Children[sub] != nil {
			q.Children[sub] = NewQuadTree(subRegion(q.Data.Region, sub))
			q.Children[sub].Insert(b)
		}

	}
}

// return whether the position is contained in the rectangle
func fitsIn(v math.Vec2, r math.Rect) bool {
	return r.PointCollision(v)
}

// Compute the rectangle of the subregion, from 0 to 3, from left to right and top to bottom
func subRegion(region math.Rect, index int) math.Rect {
	switch index {
	case 0:
		return math.Rect{Position: region.Position, Size: region.Size.Scale(0.5)}
	case 1:
		return math.Rect{Position: region.Position.Add(math.Vec2{X: 0, Y: region.Size.Y / 2}), Size: region.Size.Scale(0.5)}
	case 2:
		return math.Rect{Position: region.Position.Add(math.Vec2{X: region.Size.X / 2, Y: 0}), Size: region.Size.Scale(0.5)}
	case 3:
		return math.Rect{Position: region.Position.Add(math.Vec2{X: region.Size.X / 2, Y: region.Size.Y / 2}), Size: region.Size.Scale(0.5)}
	default: // this should NEVER happen
		fmt.Println("No THERE IS A PROBLEM WITH INDEX :", index)
		return math.Rect{}
	}
}

// Choose the subtree in which the position fits
func selectSubtree(v math.Vec2, r math.Rect) int {
	switch {
	case fitsIn(v, subRegion(r, 0)):
		return 0
	case fitsIn(v, subRegion(r, 1)):
		return 1
	case fitsIn(v, subRegion(r, 2)):
		return 2
	case fitsIn(v, subRegion(r, 3)):
		return 3
	default:
		fmt.Println(r.PointCollision(v), " and ", fitsIn(v, subRegion(r, 0)) || fitsIn(v, subRegion(r, 1)) || fitsIn(v, subRegion(r, 2)) || fitsIn(v, subRegion(r, 3)))

		fmt.Println("SHIT THE BODY IS OUT OF BOUNDS, pos", v, " and region :", r)
		panic("out of bounds")
	}
}

//----------------------------------------------------------------------------------------
// UPDATE STATS --------------------------------------------------------------------------

// Approximation of the center of mass
func (q *QuadTree) UpdateMass() {
	switch q.Data.treeType {
	case Empty: // Don't do anything
		return
	case Leaf: // Base case
		q.Data.Mass = q.Data.body.Mass
		q.Data.MassCenter = q.Data.body.Position

	case Node: // Weighted average of the tree's childrens
		massCenter := math.NewVec2(0, 0)
		totalMass := float64(0) // sum of the weight in the average

		for _, child := range q.Children {
			massCenter = massCenter.Add(child.Data.MassCenter).Scale(child.Data.Mass)
			totalMass += child.Data.Mass
		}

		q.Data.Mass = totalMass
		if totalMass != 0 {
			q.Data.MassCenter = massCenter.Scale(1 / totalMass)
		}
	}
}

//----------------------------------------------------------------------------------------
// UPDATE FORCES -------------------------------------------------------------------------

// Update the force without approximation
func (q *QuadTree) UpdateForcesNormal(otherBodies []*Systems.Body) {
	if len(otherBodies) == 0 { // Case if the tree is the parent of the subtree being calculated normally

		otherBodies = q.ListChildren()

	}

	switch q.Data.treeType {
	case Empty:
		return
	case Leaf:
		q.Data.ForceOnBody = UpdateBodyForces(q.Data.body, otherBodies)

	case Node:
		for _, child := range q.Children {
			if child != nil {
				child.UpdateForcesNormal(otherBodies)

			}
		}
	}

}

// Update forces of a single body
func UpdateBodyForces(b *Systems.Body, otherBodies []*Systems.Body) physicUnit.Force2D {
	var force physicUnit.Force2D = m.NewVec2(0, 0)

	for _, otherBody := range otherBodies {
		if otherBody != b {
			vector := otherBody.Position.Substract(b.Position).Normalize()
			r := vector.GetNorm()
			attraction := vector.Scale(physicUnit.G * otherBody.Mass * b.Mass / r * r)

			force = force.Add(attraction)
		}
	}

	return force
}

// Give the list of all the recursive childrens of the tree
func (q *QuadTree) ListChildren() []*Systems.Body {
	switch q.Data.treeType {
	case Empty:
		return []*Systems.Body{}
	case Leaf:
		return []*Systems.Body{q.Data.body}
	case Node:
		slice := []*Systems.Body{}
		for i, child := range q.Children {
			if child != nil {
				fmt.Println(i, "th child: ", child)
				fmt.Println("part 1: ", slice, "part 2: ", child.ListChildren())
				slice = slices.Concat(slice, child.ListChildren())

			}
		}
		fmt.Println("final slice: ", slice)

		return slice
	}

	return []*Systems.Body{}
}

//----------------------------------------------------------------------------------------
// UPDATE POSITIONS ----------------------------------------------------------------------

func (q *QuadTree) UpdatePositions(root *QuadTree) {
	switch q.Data.treeType {
	case Empty:
		return
	case Leaf:

		q.Data.body.UpdatePosition(q.Data.ForceOnBody)

		// Now the body is probably no longer in the good region :(
		// This is why we re-insert it from the root of the tree and make sure this becomes empty
		temp := q.Data.body // we use temp to prevent edge case if we set to nil after it was re-inserted
		q.Data.body = nil
		q.Data.treeType = Empty

		root.Insert(temp)

	case Node:
		for _, child := range q.Children {
			if child != nil {
				child.UpdatePositions(root)

			}
		}
	}
}

//----------------------------------------------------------------------------------------
// USEFUL FUNCTIONS ----------------------------------------------------------------------

// Print the tree in console
func (q *QuadTree) PrintTree() {
	fmt.Print(q.Data)
	if q.Data.treeType != Empty {
		for _, e := range q.Children {
			e.PrintTree()
		}
	}
}

// Print the tree on screen
func (q *QuadTree) ShowRegion(n uint8) {
	q.Data.Region.Draw(rl.NewColor(0, 0, 255-n, 128))
	for _, child := range q.Children {
		if child != nil {
			child.ShowRegion(n + 20)

		}
	}
}
