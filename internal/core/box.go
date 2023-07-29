package core

type Box struct {
	boxMin, boxMax Vec3
	sides          ObjectList
}

func NewBox(p0, p1 Vec3, material Material) *Box {
	sides := make(ObjectList, 6)
	sides[0] = NewXYRect(p0.x, p1.x, p0.y, p1.y, p1.z, material)
	sides[1] = NewXYRect(p0.x, p1.x, p0.y, p1.y, p0.z, material)
	sides[2] = NewXZRect(p0.x, p1.x, p0.z, p1.z, p1.y, material)
	sides[3] = NewXZRect(p0.x, p1.x, p0.z, p1.z, p0.y, material)
	sides[4] = NewYZRect(p0.y, p1.y, p0.z, p1.z, p1.x, material)
	sides[5] = NewYZRect(p0.y, p1.y, p0.z, p1.z, p0.x, material)
	return &Box{p0, p1, sides}
}

func (b *Box) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	return b.sides.hit(r, tMin, tMax)
}

func (b *Box) boundingBox(t0, t1 Double) (bool, *aabb) {
	return true, newAabb(b.boxMin, b.boxMax)
}
