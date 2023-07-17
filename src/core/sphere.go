package core

import "math"

type sphere struct {
	center   Vec3
	radius   double
	material Material
}

func NewSphere(center Vec3, radius double, mat Material) *sphere {
	return &sphere{center, radius, mat}
}

func (s *sphere) Hit(r *Ray, tMin, tMax double) (bool, *HitRecord) {
	// (P(t) - C) * (P(t) - C) = r^2
	// (A + tb - C) * (A + tb - C) = r^2
	// t^2 b^2 + 2tb(A-C) + (A-C)(A-C) - r^2 = 0
	center := s.center
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

func (s *sphere) BoundingBox(t0, t1 double) (bool, *Aabb) {
	radius := math.Abs(s.radius)
	return true, NewAabb(
		s.center.sub(vec3(radius, radius, radius)),
		s.center.add(vec3(radius, radius, radius)))
}

func getSphereUv(p Vec3) (double, double) {
	phi := math.Atan2(p.z, p.x)
	theta := math.Asin(p.y)
	u := 1 - (phi+math.Pi)/(2*math.Pi)
	v := (theta + math.Pi/2) / math.Pi
	return u, v
}
