package core

type Ray struct {
	Origin    Vec3
	Direction Vec3
	Time      Double
}

func NewRay(origin, direction Vec3, tm Double) *Ray {
	return &Ray{origin, direction, tm}
}

func (r *Ray) At(t Double) Vec3 {
	return r.Origin.add(r.Direction.mul(t))
}
