package core

import (
	"math"
	"math/rand"
)

type Camera struct {
	Origin          *Vec3
	LowerLeftCorner *Vec3
	horizontal      *Vec3
	vertical        *Vec3
	lensRadius      double
	u, v, w         *Vec3
}

func NewCamera(
	aspectRatio, vfov, aperture, focusDist double,
	lookFrom, lookAt, vup *Vec3) *Camera {
	theta := vfov * math.Pi / 180
	h := math.Tan(theta / 2)

	viewportHeight := 2.0 * h
	viewportWidth := aspectRatio * viewportHeight

	w := lookFrom.Sub(lookAt).Norm()
	u := vup.Cross(w).Norm()
	v := w.Cross(u)

	origin := lookFrom
	horizontal := u.Mul(viewportWidth)
	vertical := v.Mul(viewportHeight)
	lowerLeftCorner :=
		origin.Sub(horizontal.Div(2)).Sub(vertical.Div(2)).Sub(w.Mul(focusDist))

	return &Camera{origin, lowerLeftCorner, horizontal, vertical, aperture / 2, u, v, w}
}

func (c *Camera) GetRay(u, v double) *Ray {
	rd := randomInUnitDisk().Mul_(c.lensRadius)
	offset := c.u.Mul(rd.X).Add_(c.v.Mul(rd.Y))

	return NewRay(
		c.Origin.Add(offset),
		c.LowerLeftCorner.Add(c.horizontal.Mul(u)).Add_(c.vertical.Mul(v)).Sub_(c.Origin).Sub_(offset))
}

func randomInUnitDisk() *Vec3 {
	a := rand.Float64() * 2.0 * math.Pi
	r := rand.Float64()
	return NewVec3(r*math.Cos(a), r*math.Sin(a), 0)
}
