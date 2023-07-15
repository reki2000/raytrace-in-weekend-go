package core

import "math"

type movingSphere struct {
	center1, center2 *Vec3
	radius           double
	material         Material
	time0, time1     double
}

func NewMovingSphere(center1, center2 *Vec3, radius double, mat Material, time0, time1 double) *movingSphere {
	return &movingSphere{center1, center2, radius, mat, time0, time1}
}

func (s *movingSphere) Center(t double) *Vec3 {
	return s.center1.Add(s.center2.Sub(s.center1).Mul((s.time1 - t) / (s.time1 - s.time0)))
}

func (s *movingSphere) Hit(r *Ray, tMin, tMax double) (bool, *HitRecord) {
	// (P(t) - C) * (P(t) - C) = r^2
	// (A + tb - C) * (A + tb - C) = r^2
	// t^2 b^2 + 2tb(A-C) + (A-C)(A-C) - r^2 = 0
	center := s.Center(r.Time)
	oc := r.Origin.Sub(center) // A-C
	a := r.Direction.LengthSquared()
	halfB := oc.Dot(r.Direction)
	c := oc.LengthSquared() - s.radius*s.radius // (A-C)(A-C) - r^2
	discriminant := halfB*halfB - a*c

	if discriminant >= 0 {
		root := math.Sqrt(discriminant)

		temp := (-halfB - root) / a
		if temp < tMax && temp > tMin {
			p := r.At(temp)
			outwardNormal := p.Sub(center).Div_(s.radius)
			return true, NewHitRecord(temp, p, r, outwardNormal, s.material)
		}

		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			p := r.At(temp)
			outwardNormal := p.Sub(center).Div_(s.radius)
			return true, NewHitRecord(temp, p, r, outwardNormal, s.material)
		}
	}

	return false, nil
}

func (s *movingSphere) BoundingBox(t0, t1 double) (bool, *Aabb) {
	aabb0 := NewAabb(
		s.Center(t0).Sub(NewVec3(s.radius, s.radius, s.radius)),
		s.Center(t0).Add(NewVec3(s.radius, s.radius, s.radius)))
	aabb1 := NewAabb(
		s.Center(t1).Sub(NewVec3(s.radius, s.radius, s.radius)),
		s.Center(t1).Add(NewVec3(s.radius, s.radius, s.radius)))
	return true, NewSurroundingBox(aabb0, aabb1)
}
