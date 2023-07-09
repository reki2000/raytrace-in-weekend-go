package core

type Ray struct {
	Origin    *Vec3
	Direction *Vec3
	Time      double
}

func NewRay(origin, direction *Vec3, tm double) *Ray {
	return &Ray{origin, direction, tm}
}

func (r *Ray) At(t double) *Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}
