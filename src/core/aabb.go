package core

import "math"

type Aabb struct {
	Min, Max *Vec3
}

func NewAabb(a, b *Vec3) *Aabb {
	return &Aabb{a, b}
}

func (a *Aabb) Hit(r *Ray, tMin, tMax double) bool {
	return hit(a.Min.X-r.Origin.X, a.Max.X-r.Origin.X, r.invDirection.X, tMin, tMax) &&
		hit(a.Min.Y-r.Origin.Y, a.Max.Y-r.Origin.Y, r.invDirection.Y, tMin, tMax) &&
		hit(a.Min.Z-r.Origin.Z, a.Max.Z-r.Origin.Z, r.invDirection.Z, tMin, tMax)
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
	small := &Vec3{
		math.Min(box0.Min.X, box1.Min.X),
		math.Min(box0.Min.Y, box1.Min.Y),
		math.Min(box0.Min.Z, box1.Min.Z),
	}
	big := &Vec3{
		math.Max(box0.Max.X, box1.Max.X),
		math.Max(box0.Max.Y, box1.Max.Y),
		math.Max(box0.Max.Z, box1.Max.Z),
	}
	return NewAabb(small, big)
}
