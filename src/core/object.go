package core

type Object interface {
	Hit(r *Ray, tMin, tMax double) (bool, *HitRecord)
	BoundingBox(t0, t1 double) (bool, *Aabb)
}

type ObjectList []Object

func (objs ObjectList) Hit(r *Ray, tMin, tMax double) (bool, *HitRecord) {
	hitAny := false
	closestSoFar := tMax
	var hitRecord *HitRecord

	for _, obj := range objs {
		if hit, hr := obj.Hit(r, tMin, closestSoFar); hit {
			hitAny = true
			closestSoFar = hr.T
			hitRecord = hr
		}
	}

	return hitAny, hitRecord
}

func (objs ObjectList) BoundingBox(t0, t1 double) (bool, *Aabb) {
	if len(objs) == 0 {
		return false, nil
	}

	var outputBox *Aabb
	firstBox := true

	for _, obj := range objs {
		ok, tempBox := obj.BoundingBox(t0, t1)
		if !ok {
			return false, nil
		}

		if firstBox {
			outputBox = tempBox
		} else {
			outputBox = NewSurroundingBox(outputBox, tempBox)
		}
		firstBox = false
	}

	return true, outputBox
}

type HitRecord struct {
	T         double
	P         Vec3
	Normal    Vec3
	FrontFace bool
	Mat       Material
	u, v      double
}

func NewHitRecord(t double, p Vec3, u, v double, r *Ray, outwardNormal Vec3, material Material) *HitRecord {
	hr := &HitRecord{T: t, P: p, Mat: material, u: u, v: v}
	hr.setFaceNormal(r, outwardNormal)
	return hr
}

func (hr *HitRecord) setFaceNormal(r *Ray, outwardNormal Vec3) {
	hr.FrontFace = r.Direction.dot(outwardNormal) < 0
	if hr.FrontFace {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal.inv()
	}
}
