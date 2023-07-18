package core

type Object interface {
	hit(r *Ray, tMin, tMax Double) (bool, *hitRecord)
	boundingBox(t0, t1 Double) (bool, *aabb)
}

type ObjectList []Object

func (objs ObjectList) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	hitAny := false
	closestSoFar := tMax
	var hitRecord *hitRecord

	for _, obj := range objs {
		if hit, hr := obj.hit(r, tMin, closestSoFar); hit {
			hitAny = true
			closestSoFar = hr.t
			hitRecord = hr
		}
	}

	return hitAny, hitRecord
}

func (objs ObjectList) boundingBox(t0, t1 Double) (bool, *aabb) {
	if len(objs) == 0 {
		return false, nil
	}

	var outputBox *aabb
	firstBox := true

	for _, obj := range objs {
		ok, tempBox := obj.boundingBox(t0, t1)
		if !ok {
			return false, nil
		}

		if firstBox {
			outputBox = tempBox
		} else {
			outputBox = newSurroundingBox(outputBox, tempBox)
		}
		firstBox = false
	}

	return true, outputBox
}

type hitRecord struct {
	t         Double
	p         Vec3
	normal    Vec3
	frontFace bool
	material  Material
	u, v      Double
}

func newHitRecord(t Double, p Vec3, u, v Double, r *Ray, outwardNormal Vec3, material Material) *hitRecord {
	hr := &hitRecord{t: t, p: p, material: material, u: u, v: v}
	hr.setFaceNormal(r, outwardNormal)
	return hr
}

func (hr *hitRecord) setFaceNormal(r *Ray, outwardNormal Vec3) {
	hr.frontFace = r.Direction.dot(outwardNormal) < 0
	if hr.frontFace {
		hr.normal = outwardNormal
	} else {
		hr.normal = outwardNormal.inv()
	}
}
