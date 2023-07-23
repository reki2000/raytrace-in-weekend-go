package core

import (
	"math"
)

type constantMedium struct {
	boundary      Object
	negInvDensity Double
	phaseFunc     Material
}

func NewConstantMedium(boundary Object, density Double, texture Texture) *constantMedium {
	return &constantMedium{boundary, -1 / density, NewIsotropic(texture)}
}

func (m *constantMedium) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	hit1, rec1 := m.boundary.hit(r, math.Inf(-1), math.Inf(1))
	if !hit1 {
		return false, nil
	}

	hit2, rec2 := m.boundary.hit(r, rec1.t+0.0001, math.Inf(1))
	if !hit2 {
		return false, nil
	}

	if rec1.t < tMin {
		rec1.t = tMin
	}
	if rec2.t > tMax {
		rec2.t = tMax
	}

	if rec1.t >= rec2.t {
		return false, nil
	}
	if rec1.t < 0 {
		rec1.t = 0
	}

	rayLength := r.Direction.Length()
	distanceInsideBoundary := (rec2.t - rec1.t) * rayLength
	hitDistance := m.negInvDensity * math.Log(randomDouble())
	if hitDistance > distanceInsideBoundary {
		return false, nil
	}

	t := rec1.t + hitDistance/rayLength
	return true, newHitRecord(t, r.At(t), rec1.u, rec1.v, r, vec3(1, 0, 0), m.phaseFunc)
}

func (m *constantMedium) boundingBox(t0, t1 Double) (bool, *aabb) {
	return m.boundary.boundingBox(t0, t1)
}
