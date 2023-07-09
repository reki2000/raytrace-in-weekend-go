package core

type Object interface {
	Hit(r *Ray, tMin, tMax double) (bool, *HitRecord)
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

type HitRecord struct {
	T         double
	P         *Vec3
	Normal    *Vec3
	FrontFace bool
	Mat       Material
}

func NewHitRecord(t double, p *Vec3, r *Ray, outwardNormal *Vec3, material Material) *HitRecord {
	hr := &HitRecord{T: t, P: p, Mat: material}
	hr.setFaceNormal(r, outwardNormal)
	return hr
}

func (hr *HitRecord) setFaceNormal(r *Ray, outwardNormal *Vec3) {
	hr.FrontFace = r.Direction.Dot(outwardNormal) < 0
	if hr.FrontFace {
		hr.Normal = outwardNormal
	} else {
		hr.Normal = outwardNormal.Inv()
	}
}
