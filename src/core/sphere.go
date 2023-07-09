package core

import "math"

type sphere struct {
	center1, center2 *Vec3
	radius           double
	material         Material
	time0, time1     double
}

func (s *sphere) Center(t double) *Vec3 {
	return s.center1.Add(s.center2.Sub(s.center1).Mul((s.time1 - t) / (s.time1 - s.time0)))
}

func (s *sphere) Hit(r *Ray, tMin, tMax double) (bool, *HitRecord) {
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

func NewSphere(center *Vec3, radius double, mat Material) *sphere {
	return &sphere{center, center, radius, mat, -1.0, 1.0}
}

func NewMovingSphere(center1, center2 *Vec3, radius double, mat Material, time0, time1 double) *sphere {
	return &sphere{center1, center2, radius, mat, time0, time1}
}
