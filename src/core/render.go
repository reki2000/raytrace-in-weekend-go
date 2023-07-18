package core

import "math"

var black = NewColor(0, 0, 0)

func RayColor(r *Ray, world ObjectList, depth int) Color {
	if depth <= 0 {
		return black
	}

	if hit, hr := world.hit(r, 0.001, math.Inf(0)); hit {
		if scattered, scatter, attenuation := hr.material.scatter(r, hr); scattered {
			return RayColor(scatter, world, depth-1).MulVec(attenuation)
		} else {
			return black
		}
	}

	unitDirection := r.Direction.norm()
	t := 0.5 * (unitDirection.y + 1.0)
	v1 := NewColor(1.0, 1.0, 1.0)
	v2 := NewColor(0.5, 0.7, 1.0)
	return v1.Mul(1.0 - t).Add(v2.Mul(t))
}
