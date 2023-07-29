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

type xzRect struct {
	x0, x1, z0, z1, k Double
	material          Material
}

func NewXZRect(x0, x1, z0, z1, k Double, material Material) *xzRect {
	return &xzRect{x0, x1, z0, z1, k, material}
}

func (rect *xzRect) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	t := (rect.k - r.Origin.y) / r.Direction.y
	if t < tMin || t > tMax {
		return false, nil
	}

	x := r.Origin.x + t*r.Direction.x
	z := r.Origin.z + t*r.Direction.z
	if x < rect.x0 || x > rect.x1 || z < rect.z0 || z > rect.z1 {
		return false, nil
	}

	u := (x - rect.x0) / (rect.x1 - rect.x0)
	v := (z - rect.z0) / (rect.z1 - rect.z0)
	rec := newHitRecord(t, r.At(t), u, v, r, vec3(0, 1, 0), rect.material)
	return true, rec
}

func (rect *xzRect) boundingBox(t0, t1 Double) (bool, *aabb) {
	return true, newAabb(vec3(rect.x0, rect.k-0.0001, rect.z0), vec3(rect.x1, rect.k+0.0001, rect.z1))
}

type yzRect struct {
	y0, y1, z0, z1, k Double
	material          Material
}

func NewYZRect(y0, y1, z0, z1, k Double, material Material) *yzRect {
	return &yzRect{y0, y1, z0, z1, k, material}
}

func (rect *yzRect) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	t := (rect.k - r.Origin.x) / r.Direction.x
	if t < tMin || t > tMax {
		return false, nil
	}

	y := r.Origin.y + t*r.Direction.y
	z := r.Origin.z + t*r.Direction.z
	if y < rect.y0 || y > rect.y1 || z < rect.z0 || z > rect.z1 {
		return false, nil
	}

	u := (y - rect.y0) / (rect.y1 - rect.y0)
	v := (z - rect.z0) / (rect.z1 - rect.z0)
	rec := newHitRecord(t, r.At(t), u, v, r, vec3(1, 0, 0), rect.material)
	return true, rec
}

func (rect *yzRect) boundingBox(t0, t1 Double) (bool, *aabb) {
	return true, newAabb(vec3(rect.k-0.0001, rect.y0, rect.z0), vec3(rect.k+0.0001, rect.y1, rect.z1))
}
