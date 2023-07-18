package core

import "math"

type aabb struct {
	min, max Vec3
}

func newAabb(a, b Vec3) *aabb {
	return &aabb{a, b}
}

func (a *aabb) hit(r *Ray, tMin, tMax Double) bool {
	return hit(a.min.x-r.Origin.x, a.max.x-r.Origin.x, r.invDirection.x, tMin, tMax) &&
		hit(a.min.y-r.Origin.y, a.max.y-r.Origin.y, r.invDirection.y, tMin, tMax) &&
		hit(a.min.z-r.Origin.z, a.max.z-r.Origin.z, r.invDirection.z, tMin, tMax)
}

func hit(aabbMin, aabbMax, inv, tMin, tMax Double) bool {
	var t0, t1 Double
	if inv < 0.0 {
		t1 = aabbMin * inv
		t0 = aabbMax * inv
	} else {
		t0 = aabbMin * inv
		t1 = aabbMax * inv
	}

	if t0 > tMin {
		tMin = t0
	}
	if t1 < tMax {
		tMax = t1
	}

	return tMax > tMin
}

func newSurroundingBox(box0, box1 *aabb) *aabb {
	small := vec3(
		math.Min(box0.min.x, box1.min.x),
		math.Min(box0.min.y, box1.min.y),
		math.Min(box0.min.z, box1.min.z),
	)
	big := vec3(
		math.Max(box0.max.x, box1.max.x),
		math.Max(box0.max.y, box1.max.y),
		math.Max(box0.max.z, box1.max.z),
	)
	return newAabb(small, big)
}
