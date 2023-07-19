package core

import "math"

type aabb struct {
	min, max Vec3
}

func newAabb(a, b Vec3) *aabb {
	return &aabb{a, b}
}

func (a *aabb) hit(r *Ray, tMin, tMax Double) bool {
	return hitAabb(a.min.x-r.Origin.x, a.max.x-r.Origin.x, r.Direction.x, tMin, tMax) &&
		hitAabb(a.min.y-r.Origin.y, a.max.y-r.Origin.y, r.Direction.y, tMin, tMax) &&
		hitAabb(a.min.z-r.Origin.z, a.max.z-r.Origin.z, r.Direction.z, tMin, tMax)
}

func hitAabb(aabbMin, aabbMax, dir, tMin, tMax Double) bool {
	tMin, tMax = tMin*dir, tMax*dir

	var t0, t1 Double
	if dir < 0.0 {
		t1 = aabbMin
		t0 = aabbMax

		if t0 < tMin {
			tMin = t0
		}
		if t1 > tMax {
			tMax = t1
		}

		return tMax < tMin
	}

	t0 = aabbMin
	t1 = aabbMax

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
