package core

type Ray struct {
	Origin    *Vec3
	Direction *Vec3
}

func NewRay(origin, direction *Vec3) *Ray {
	return &Ray{origin, direction}
}

func (r *Ray) At(t double) *Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}
