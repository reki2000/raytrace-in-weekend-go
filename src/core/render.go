package core

import "math"

func RayColor(r *Ray, backGround Color, world ObjectList, depth int) Color {
	if depth <= 0 {
		return black
	}

	if hit, hr := world.hit(r, 0.001, math.Inf(0)); hit {
		emitted := hr.material.emitted(hr.u, hr.v, hr.p)
		if scattered, scatter, attenuation := hr.material.scatter(r, hr); scattered {
			return emitted.Add(RayColor(scatter, backGround, world, depth-1).MulVec(attenuation))
		} else {
			return emitted
		}
	}

	return backGround
}
