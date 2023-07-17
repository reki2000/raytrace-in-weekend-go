package core

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(r *Ray, hr *HitRecord) (bool, *Ray, Color)
}

type lambertian struct {
	Albedo Texture
}

func NewLambertian(albedo Texture) *lambertian {
	return &lambertian{albedo}
}

func (l *lambertian) Scatter(r *Ray, hr *HitRecord) (bool, *Ray, Color) {
	scatterDirection := hr.Normal.add(randomOnUnitShpere())
	return true, NewRay(hr.P, scatterDirection, r.Time), l.Albedo.Value(hr.u, hr.v, hr.P)
}

type metal struct {
	Albedo Color
	Fuzz   double
}

func NewMetal(albedo Color, fuzz double) *metal {
	if fuzz < 1 {
		return &metal{albedo, fuzz}
	} else {
		return &metal{albedo, 1.0}
	}
}

func (m *metal) Scatter(r *Ray, hr *HitRecord) (bool, *Ray, Color) {
	reflected := r.Direction.norm().reflect(hr.Normal)
	scattered := NewRay(hr.P, reflected.add(randomOnUnitShpere().mul(m.Fuzz)), r.Time)
	if (scattered.Direction.dot(hr.Normal)) > 0 {
		return true, scattered, m.Albedo
	} else {
		return false, nil, Color{}
	}
}

type dielectric struct {
	RefIdx double
}

func NewDielectric(refIdx double) *dielectric {
	return &dielectric{refIdx}
}

var white = NewColor(1.0, 1.0, 1.0)

func (d *dielectric) Scatter(r *Ray, hr *HitRecord) (bool, *Ray, Color) {
	var etaIOverEtaT double
	if hr.FrontFace {
		etaIOverEtaT = 1.0 / d.RefIdx
	} else {
		etaIOverEtaT = d.RefIdx
	}

	unitDirection := r.Direction.norm()
	cosTheta := math.Min(unitDirection.inv().dot(hr.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	if etaIOverEtaT*sinTheta > 1.0 {
		reflected := unitDirection.reflect(hr.Normal)
		return true, NewRay(hr.P, reflected, r.Time), white
	}

	reflectProb := schlick(cosTheta, etaIOverEtaT)
	if rand.Float64() < reflectProb {
		reflected := unitDirection.reflect(hr.Normal)
		return true, NewRay(hr.P, reflected, r.Time), white
	}

	refracted := unitDirection.refract(hr.Normal, etaIOverEtaT)
	return true, NewRay(hr.P, refracted, r.Time), white
}

func schlick(cosine double, refIdx double) double {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}
