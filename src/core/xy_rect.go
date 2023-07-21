package core

type xyRect struct {
	x0, x1, y0, y1, k Double
	material          Material
}

func NewXYRect(x0, x1, y0, y1, k Double, material Material) *xyRect {
	return &xyRect{x0, x1, y0, y1, k, material}
}

func (rect *xyRect) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	t := (rect.k - r.Origin.z) / r.Direction.z
	if t < tMin || t > tMax {
		return false, nil
	}

	x := r.Origin.x + t*r.Direction.x
	y := r.Origin.y + t*r.Direction.y
	if x < rect.x0 || x > rect.x1 || y < rect.y0 || y > rect.y1 {
		return false, nil
	}

	u := (x - rect.x0) / (rect.x1 - rect.x0)
	v := (y - rect.y0) / (rect.y1 - rect.y0)
	rec := newHitRecord(t, r.At(t), u, v, r, vec3(0, 0, 1), rect.material)
	return true, rec
}

func (rect *xyRect) boundingBox(t0, t1 Double) (bool, *aabb) {
	return true, newAabb(vec3(rect.x0, rect.y0, rect.k-0.0001), vec3(rect.x1, rect.y1, rect.k+0.0001))
}
