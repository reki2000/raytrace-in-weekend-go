package core

import (
	"math"
	"math/rand"
)

type Material interface {
	Scatter(r *Ray, hr *HitRecord) (bool, *Ray, *Vec3)
}

type lambertian struct {
	Albedo *Vec3
}

func NewLambertian(albedo *Vec3) *lambertian {
	return &lambertian{albedo}
}

func (l *lambertian) Scatter(r *Ray, hr *HitRecord) (bool, *Ray, *Vec3) {
	scatterDirection := hr.Normal.Add(randomOnUnitShpere())
	return true, &Ray{hr.P, scatterDirection}, l.Albedo
}

type metal struct {
	Albedo *Vec3
	Fuzz   double
}

func NewMetal(albedo *Vec3, fuzz double) *metal {
	if fuzz < 1 {
		return &metal{albedo, fuzz}
	} else {
		return &metal{albedo, 1.0}
	}
}

func (m *metal) Scatter(r *Ray, hr *HitRecord) (bool, *Ray, *Vec3) {
	reflected := r.Direction.Norm().Reflect(hr.Normal)
	scattered := &Ray{hr.P, reflected.Add(randomOnUnitShpere().Mul_(m.Fuzz))}
	if (scattered.Direction.Dot(hr.Normal)) > 0 {
		return true, scattered, m.Albedo
	} else {
		return false, nil, nil
	}
}

type dielectric struct {
	RefIdx double
}

func NewDielectric(refIdx double) *dielectric {
	return &dielectric{refIdx}
}

func color3(r, g, b double) *Vec3 {
	return NewVec3(r, g, b)
}

func (d *dielectric) Scatter(r *Ray, hr *HitRecord) (bool, *Ray, *Vec3) {
	var etaIOverEtaT double
	if hr.FrontFace {
		etaIOverEtaT = 1.0 / d.RefIdx
	} else {
		etaIOverEtaT = d.RefIdx
	}

	unitDirection := r.Direction.Norm()
	cosTheta := math.Min(unitDirection.Inv().Dot(hr.Normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	if etaIOverEtaT*sinTheta > 1.0 {
		reflected := unitDirection.Reflect(hr.Normal)
		return true, NewRay(hr.P, reflected), color3(1.0, 1.0, 1.0)
	}

	reflectProb := schlick(cosTheta, etaIOverEtaT)
	if rand.Float64() < reflectProb {
		reflected := unitDirection.Reflect(hr.Normal)
		return true, NewRay(hr.P, reflected), color3(1.0, 1.0, 1.0)
	}

	refracted := unitDirection.Refract(hr.Normal, etaIOverEtaT)
	return true, NewRay(hr.P, refracted), color3(1.0, 1.0, 1.0)
}

func schlick(cosine double, refIdx double) double {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

func randomOnUnitShpere() *Vec3 {
	a := rand.Float64() * 2.0 * math.Pi
	z := rand.Float64()*2.0 - 1.0
	r := math.Sqrt(1 - z*z)
	return NewVec3(r*math.Cos(a), r*math.Sin(a), z)
}
