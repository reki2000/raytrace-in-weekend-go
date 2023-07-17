package core

import "math"

type Aabb struct {
	Min, Max Vec3
}

func NewAabb(a, b Vec3) *Aabb {
	return &Aabb{a, b}
}

func (a *Aabb) Hit(r *Ray, tMin, tMax double) bool {
	return hit(a.Min.x-r.Origin.x, a.Max.x-r.Origin.x, r.invDirection.x, tMin, tMax) &&
		hit(a.Min.y-r.Origin.y, a.Max.y-r.Origin.y, r.invDirection.y, tMin, tMax) &&
		hit(a.Min.z-r.Origin.z, a.Max.z-r.Origin.z, r.invDirection.z, tMin, tMax)
}

func hit(aabbMin, aabbMax, inv, tMin, tMax double) bool {
	var t0, t1 double
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

func NewSurroundingBox(box0, box1 *Aabb) *Aabb {
	small := NewVec3(
		math.Min(box0.Min.x, box1.Min.x),
		math.Min(box0.Min.y, box1.Min.y),
		math.Min(box0.Min.z, box1.Min.z),
	)
	big := NewVec3(
		math.Max(box0.Max.x, box1.Max.x),
		math.Max(box0.Max.y, box1.Max.y),
		math.Max(box0.Max.z, box1.Max.z),
	)
	return NewAabb(small, big)
}
