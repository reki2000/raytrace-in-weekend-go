package core

import "math"

type movingSphere struct {
	center1, center2 Vec3
	radius           double
	material         Material
	time0, time1     double
}

func NewMovingSphere(center1, center2 Vec3, radius double, mat Material, time0, time1 double) *movingSphere {
	return &movingSphere{center1, center2, radius, mat, time0, time1}
}

func (s *movingSphere) Center(t double) Vec3 {
	return s.center1.add(s.center2.sub(s.center1).mul((s.time1 - t) / (s.time1 - s.time0)))
}

func (s *movingSphere) Hit(r *Ray, tMin, tMax double) (bool, *HitRecord) {
	// (P(t) - C) * (P(t) - C) = r^2
	// (A + tb - C) * (A + tb - C) = r^2
	// t^2 b^2 + 2tb(A-C) + (A-C)(A-C) - r^2 = 0
	center := s.Center(r.Time)
	oc := r.Origin.sub(center) // A-C
	a := r.Direction.lengthSquared()
	halfB := oc.dot(r.Direction)
	c := oc.lengthSquared() - s.radius*s.radius // (A-C)(A-C) - r^2
	discriminant := halfB*halfB - a*c

	if discriminant >= 0 {
		root := math.Sqrt(discriminant)

		temp := (-halfB - root) / a
		if temp < tMax && temp > tMin {
			p := r.At(temp)
			outwardNormal := p.sub(center).div(s.radius)
			u, v := getSphereUv(outwardNormal)
			return true, NewHitRecord(temp, p, u, v, r, outwardNormal, s.material)
		}

		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			p := r.At(temp)
			outwardNormal := p.sub(center).div(s.radius)
			u, v := getSphereUv(outwardNormal)
			return true, NewHitRecord(temp, p, u, v, r, outwardNormal, s.material)
		}
	}

	return false, nil
}

func (s *movingSphere) BoundingBox(t0, t1 double) (bool, *Aabb) {
	aabb0 := NewAabb(
		s.Center(t0).sub(NewVec3(s.radius, s.radius, s.radius)),
		s.Center(t0).add(NewVec3(s.radius, s.radius, s.radius)))
	aabb1 := NewAabb(
		s.Center(t1).sub(NewVec3(s.radius, s.radius, s.radius)),
		s.Center(t1).add(NewVec3(s.radius, s.radius, s.radius)))
	return true, NewSurroundingBox(aabb0, aabb1)
}
