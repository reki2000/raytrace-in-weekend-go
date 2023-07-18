package core

type Ray struct {
	Origin       Vec3
	Direction    Vec3
	Time         Double
	invDirection Vec3
}

func NewRay(origin, direction Vec3, tm Double) *Ray {
	inv := vec3(1.0/direction.x, 1.0/direction.y, 1.0/direction.z)
	return &Ray{origin, direction, tm, inv}
}

func (r *Ray) At(t Double) Vec3 {
	return r.Origin.add(r.Direction.mul(t))
}
