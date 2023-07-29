package core

import "math"

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

func (r *Ray) Color(backGround Color, world ObjectList, depth int) Color {
	if depth <= 0 {
		return black
	}

	if hit, hr := world.hit(r, 0.001, math.Inf(0)); hit {
		emitted := hr.material.emitted(hr.u, hr.v, hr.p)
		if scattered, scatter, attenuation := hr.material.scatter(r, hr); scattered {
			return emitted.Add(scatter.Color(backGround, world, depth-1).MulVec(attenuation))
		} else {
			return emitted
		}
	}

	return backGround
}
