package simulationOptimised

import (
	"fmt"

	"github.com/RugiSerl/physics/app/Systems"
	"github.com/RugiSerl/physics/app/math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TreeContent struct {
	Empty       bool
	Initialized bool
	body        *Systems.Body
	MassCenter  math.Vec2
	Mass        float64
	Region      math.Rect
}
type QuadTree struct {
	Data     TreeContent
	Children []QuadTree
}

func NewQuadTree(region math.Rect) QuadTree {
	return QuadTree{
		Data: TreeContent{
			Empty:       true,
			body:        nil,
			Initialized: true,
			Region:      region,
		},
		Children: make([]QuadTree, 4),
	}
}

func fitsIn(v math.Vec2, r math.Rect) bool {
	return r.PointCollision(v)
}

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

func selectSubtree(v math.Vec2, r math.Rect) (int, error) {
	switch {
	case fitsIn(v, subRegion(r, 0)):
		return 0, nil
	case fitsIn(v, subRegion(r, 1)):
		return 1, nil
	case fitsIn(v, subRegion(r, 2)):
		return 2, nil
	case fitsIn(v, subRegion(r, 3)):
		return 3, nil
	default:
		fmt.Println(r.PointCollision(v), " and ", fitsIn(v, subRegion(r, 0)) || fitsIn(v, subRegion(r, 1)) || fitsIn(v, subRegion(r, 2)) || fitsIn(v, subRegion(r, 3)))

		fmt.Println("SHIT THE BODY IS OUT OF BOUNDS, pos", v, " and region :", r)
		panic("out of bounds")

	}
}

func (q *QuadTree) Insert(b *Systems.Body) {
	fmt.Println("inserting.. pos :", b.Position, " and region : ", q.Data.Region)
	if q.Data.Empty { // empty tree
		fmt.Println("Empty")

		q.Data.body = b // no problem here
		q.Data.Empty = false
	} else if q.Data.body != nil { // leaf of the tree
		fmt.Println("Leaf")
		if b.Position == q.Data.body.Position { // make sure the bodies are not at the same exact position to avoid stack overflow
			fmt.Println("a")
			b.Position = b.Position.Add(math.Vec2{X: 0, Y: 1e-1}) // just add a little bit

		}
		// divide the quadtree to contain both bodies
		sub1, _ := selectSubtree(q.Data.body.Position, q.Data.Region)
		sub2, _ := selectSubtree(b.Position, q.Data.Region)

		if sub1 == sub2 { // don't initializate two times if it is in the same children
			q.Children[sub1] = NewQuadTree(subRegion(q.Data.Region, sub1))
			q.Children[sub1].Insert(q.Data.body)
			q.Children[sub1].Insert(b)

		} else {
			fmt.Println("sub region :", sub1)
			q.Children[sub1] = NewQuadTree(subRegion(q.Data.Region, sub1))
			q.Children[sub1].Insert(q.Data.body)

			q.Children[sub2] = NewQuadTree(subRegion(q.Data.Region, sub2))
			q.Children[sub2].Insert(b)
		}
		// this means the node is no longer a leaf
		q.Data.body = nil

	} else { // normal node
		fmt.Println("Node")
		sub, _ := selectSubtree(b.Position, q.Data.Region)
		if !q.Children[sub].Data.Initialized {
			q.Children[sub] = NewQuadTree(subRegion(q.Data.Region, sub))
		}
		q.Children[sub].Insert(b)
	}
}

func (q QuadTree) PrintTree() {
	fmt.Print(q.Data)
	if !q.Data.Empty {
		for _, e := range q.Children {
			e.PrintTree()
		}
	}
}

func (q QuadTree) ShowRegion() {
	q.Data.Region.Draw(rl.Blue)
}
