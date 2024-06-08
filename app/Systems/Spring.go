package Systems

type Spring struct {
	pointA, pointB *Body
	rate           float64
}

func NewSpring(rate float64, pointA, pointB *Body) Spring {
	return Spring{
		pointA: pointA,
		pointB: pointB,
		rate:   rate,
	}
}

func (s Spring) ApplyForce() {
	s.pointA.ApplyForce(s.pointB.Position.Substract(s.pointA.Position).Scale(s.rate))
	s.pointB.ApplyForce(s.pointA.Position.Substract(s.pointB.Position).Scale(s.rate))
}
