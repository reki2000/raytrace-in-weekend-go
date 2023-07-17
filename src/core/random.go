package core

import (
	"math"
	"math/rand"
)

func randomInUnitSphere() Vec3 {
	for {
		p := NewVec3Random(-1, 1)
		if p.lengthSquared() >= 1 {
			continue
		}
		return p
	}
}

func randomOnUnitShpere() Vec3 {
	a := rand.Float64() * 2.0 * math.Pi
	z := rand.Float64()*2.0 - 1.0
	r := math.Sqrt(1 - z*z)
	return NewVec3(r*math.Cos(a), r*math.Sin(a), z)
}
