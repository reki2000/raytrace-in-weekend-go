package core

import "math"

type sphere struct {
	center   *Vec3
	radius   double
	material Material
}

func NewSphere(center *Vec3, radius double, mat Material) *sphere {
	return &sphere{center, radius, mat}
}

func (s *sphere) Hit(r *Ray, tMin, tMax double) (bool, *HitRecord) {
	// (P(t) - C) * (P(t) - C) = r^2
	// (A + tb - C) * (A + tb - C) = r^2
	// t^2 b^2 + 2tb(A-C) + (A-C)(A-C) - r^2 = 0
	center := s.center
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

func (s *sphere) BoundingBox(t0, t1 double) (bool, *Aabb) {
	return true, NewAabb(
		s.center.Sub(NewVec3(s.radius, s.radius, s.radius)),
		s.center.Add(NewVec3(s.radius, s.radius, s.radius)))
}
