package core

type Ray struct {
	Origin       *Vec3
	Direction    *Vec3
	Time         double
	invDirection *Vec3
}

func NewRay(origin, direction *Vec3, tm double) *Ray {
	inv := NewVec3(1.0/direction.X, 1.0/direction.Y, 1.0/direction.Z)
	return &Ray{origin, direction, tm, inv}
}

func (r *Ray) At(t double) *Vec3 {
	return r.Origin.Add(r.Direction.Mul(t))
}
