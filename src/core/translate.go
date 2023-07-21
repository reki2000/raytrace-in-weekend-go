package core

import "math"

type Translate struct {
	obj    Object
	offset Vec3
	bbox   *aabb
}

func NewTranslate(obj Object, offset Vec3) *Translate {
	return &Translate{obj, offset, nil}
}

func (t *Translate) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	movedR := NewRay(r.Origin.Sub(t.offset), r.Direction, r.Time)
	if hit, rec := t.obj.hit(movedR, tMin, tMax); hit {
		rec.p = rec.p.add(t.offset)
		return true, rec
	}
	return false, nil
}

func (t *Translate) boundingBox(t0, t1 Double) (bool, *aabb) {
	if t.bbox != nil {
		return true, t.bbox
	}
	if ok, bbox := t.obj.boundingBox(t0, t1); ok {
		t.bbox = newAabb(bbox.min.add(t.offset), bbox.max.add(t.offset))
		return true, t.bbox
	}
	return false, nil
}

type rotateY struct {
	obj      Object
	sinTheta Double
	cosTheta Double
	bbox     *aabb
}

func NewRotateY(obj Object, angle Double) *rotateY {
	radians := (math.Pi / 180.0) * angle
	sinTheta := math.Sin(radians)
	cosTheta := math.Cos(radians)

	ok, bbox := obj.boundingBox(0, 1)
	if !ok {
		panic("no bounding box in rotateY constructor")
	}

	min := vec3(math.Inf(1), math.Inf(1), math.Inf(1))
	max := vec3(math.Inf(-1), math.Inf(-1), math.Inf(-1))
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				x := Double(i)*bbox.max.x + (1-Double(i))*bbox.min.x
				y := Double(j)*bbox.max.y + (1-Double(j))*bbox.min.y
				z := Double(k)*bbox.max.z + (1-Double(k))*bbox.min.z
				newX := cosTheta*x + sinTheta*z
				newZ := -sinTheta*x + cosTheta*z
				tester := vec3(newX, y, newZ)
				for c := 0; c < 3; c++ {
					e := tester.e(c)
					if e > max.e(c) {
						max = max.newE(c, e)
					} else if e < min.e(c) {
						min = min.newE(c, e)
					}
				}
			}
		}
	}
	return &rotateY{obj, sinTheta, cosTheta, newAabb(min, max)}
}

func (ry *rotateY) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	origin := vec3(ry.cosTheta*r.Origin.x-ry.sinTheta*r.Origin.z, r.Origin.y, ry.sinTheta*r.Origin.x+ry.cosTheta*r.Origin.z)
	direction := vec3(ry.cosTheta*r.Direction.x-ry.sinTheta*r.Direction.z, r.Direction.y, ry.sinTheta*r.Direction.x+ry.cosTheta*r.Direction.z)
	rotatedR := NewRay(origin, direction, r.Time)
	if hit, rec := ry.obj.hit(rotatedR, tMin, tMax); hit {
		p := vec3(ry.cosTheta*rec.p.x+ry.sinTheta*rec.p.z, rec.p.y, -ry.sinTheta*rec.p.x+ry.cosTheta*rec.p.z)
		normal := vec3(ry.cosTheta*rec.normal.x+ry.sinTheta*rec.normal.z, rec.normal.y, -ry.sinTheta*rec.normal.x+ry.cosTheta*rec.normal.z)
		return true, newHitRecord(rec.t, p, rec.u, rec.v, r, normal, rec.material)
	}
	return false, nil
}

func (ry *rotateY) boundingBox(t0, t1 Double) (bool, *aabb) {
	return true, ry.bbox
}
