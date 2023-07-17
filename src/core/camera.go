package core

import (
	"math"
	"math/rand"
)

type Camera struct {
	Origin          Vec3
	LowerLeftCorner Vec3
	horizontal      Vec3
	vertical        Vec3
	lensRadius      double
	u, v, w         Vec3
	time0, time1    double
}

func NewCamera(
	aspectRatio, vfov, aperture, focusDist double,
	lookFrom, lookAt, vup Vec3,
	time0, time1 double) *Camera {
	theta := vfov * math.Pi / 180
	h := math.Tan(theta / 2)

	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.sub(lookAt).norm()
	u := vup.cross(w).norm()
	v := w.cross(u)

	origin := lookFrom
	horizontal := u.mul(viewportWidth)
	vertical := v.mul(viewportHeight)
	lowerLeftCorner :=
		origin.sub(horizontal.div(2)).sub(vertical.div(2)).sub(w.mul(focusDist))

	return &Camera{origin, lowerLeftCorner, horizontal, vertical, aperture / 2, u, v, w, time0, time1}
}

func (c *Camera) GetRay(u, v double) *Ray {
	rd := randomInUnitDisk().mul(c.lensRadius)
	offset := c.u.mul(rd.x).add(c.v.mul(rd.y))

	return NewRay(
		c.Origin.add(offset),
		c.LowerLeftCorner.add(c.horizontal.mul(u)).add(c.vertical.mul(v)).sub(c.Origin).sub(offset),
		rand.Float64()*(c.time1-c.time0)+c.time0)
}

func randomInUnitDisk() Vec3 {
	a := rand.Float64() * 2.0 * math.Pi
	r := rand.Float64()
	return NewVec3(r*math.Cos(a), r*math.Sin(a), 0)
}
