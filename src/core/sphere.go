package core

import "math"

type sphere struct {
	center   Vec3
	radius   Double
	material Material
}

func NewSphere(center Vec3, radius Double, mat Material) *sphere {
	return &sphere{center, radius, mat}
}

func (s *sphere) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
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
			return true, newHitRecord(temp, p, u, v, r, outwardNormal, s.material)
		}

		temp = (-halfB + root) / a
		if temp < tMax && temp > tMin {
			p := r.At(temp)
			outwardNormal := p.sub(center).div(s.radius)
			u, v := getSphereUv(outwardNormal)
			return true, newHitRecord(temp, p, u, v, r, outwardNormal, s.material)
		}
	}

	return false, nil
}

func (s *sphere) boundingBox(t0, t1 Double) (bool, *aabb) {
	radius := math.Abs(s.radius)
	return true, newAabb(
		s.center.sub(vec3(radius, radius, radius)),
		s.center.add(vec3(radius, radius, radius)))
}

func getSphereUv(p Vec3) (Double, Double) {
	phi := math.Atan2(p.z, p.x)
	theta := math.Asin(p.y)
	u := 1 - (phi+math.Pi)/(2*math.Pi)
	v := (theta + math.Pi/2) / math.Pi
	return u, v
}

// moving sphere
type movingSphere struct {
	sphere
	center2      Vec3
	time0, time1 Double
}

func NewMovingSphere(center1, center2 Vec3, radius Double, mat Material, time0, time1 Double) *movingSphere {
	return &movingSphere{sphere{center1, radius, mat}, center2, time0, time1}
}

func (s *movingSphere) centerAt(t Double) Vec3 {
	return s.center.add(s.center2.sub(s.center).mul((s.time1 - t) / (s.time1 - s.time0)))
}

func (s *movingSphere) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	sphere := sphere{s.centerAt(r.Time), s.radius, s.material}
	return sphere.hit(r, tMin, tMax)
}

func (s *movingSphere) boundingBox(t0, t1 Double) (bool, *aabb) {
	aabb0 := newAabb(
		s.centerAt(t0).sub(vec3(s.radius, s.radius, s.radius)),
		s.centerAt(t0).add(vec3(s.radius, s.radius, s.radius)))
	aabb1 := newAabb(
		s.centerAt(t1).sub(vec3(s.radius, s.radius, s.radius)),
		s.centerAt(t1).add(vec3(s.radius, s.radius, s.radius)))
	return true, newSurroundingBox(aabb0, aabb1)
}
