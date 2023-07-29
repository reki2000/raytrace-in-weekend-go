package core

import (
	"math"
)

type Material interface {
	scatter(r *Ray, hr *hitRecord) (bool, *Ray, Color)
	emitted(u, v Double, p Vec3) Color
}

type lambertian struct {
	Albedo Texture
}

func NewLambertian(albedo Texture) *lambertian {
	return &lambertian{albedo}
}

func (l *lambertian) scatter(r *Ray, hr *hitRecord) (bool, *Ray, Color) {
	scatterDirection := hr.normal.add(randomOnUnitShpere())
	return true, NewRay(hr.p, scatterDirection, r.Time), l.Albedo.value(hr.u, hr.v, hr.p)
}

func (l *lambertian) emitted(u, v Double, p Vec3) Color { return black }

type metal struct {
	Albedo Color
	Fuzz   Double
}

func NewMetal(albedo Color, fuzz Double) *metal {
	if fuzz < 1 {
		return &metal{albedo, fuzz}
	} else {
		return &metal{albedo, 1.0}
	}
}

func randomOnUnitShpere() Vec3 {
	a := randomDouble() * 2.0 * math.Pi
	z := randomDouble()*2.0 - 1.0
	r := math.Sqrt(1 - z*z)
	return vec3(r*math.Cos(a), r*math.Sin(a), z)
}

func (m *metal) scatter(r *Ray, hr *hitRecord) (bool, *Ray, Color) {
	reflected := r.Direction.norm().reflect(hr.normal)
	scattered := NewRay(hr.p, reflected.add(randomOnUnitShpere().mul(m.Fuzz)), r.Time)
	if (scattered.Direction.dot(hr.normal)) > 0 {
		return true, scattered, m.Albedo
	} else {
		return false, nil, Color{}
	}
}

func (m *metal) emitted(u, v Double, p Vec3) Color { return black }

type dielectric struct {
	RefIdx Double
}

func NewDielectric(refIdx Double) *dielectric {
	return &dielectric{refIdx}
}

func (d *dielectric) scatter(r *Ray, hr *hitRecord) (bool, *Ray, Color) {
	var etaIOverEtaT Double
	if hr.frontFace {
		etaIOverEtaT = 1.0 / d.RefIdx
	} else {
		etaIOverEtaT = d.RefIdx
	}

	unitDirection := r.Direction.norm()
	cosTheta := math.Min(unitDirection.inv().dot(hr.normal), 1.0)
	sinTheta := math.Sqrt(1.0 - cosTheta*cosTheta)
	if etaIOverEtaT*sinTheta > 1.0 {
		reflected := unitDirection.reflect(hr.normal)
		return true, NewRay(hr.p, reflected, r.Time), white
	}

	reflectProb := schlick(cosTheta, etaIOverEtaT)
	if randomDouble() < reflectProb {
		reflected := unitDirection.reflect(hr.normal)
		return true, NewRay(hr.p, reflected, r.Time), white
	}

	refracted := unitDirection.refract(hr.normal, etaIOverEtaT)
	return true, NewRay(hr.p, refracted, r.Time), white
}

func schlick(cosine Double, refIdx Double) Double {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1-cosine), 5)
}

func (d *dielectric) emitted(u, v Double, p Vec3) Color { return black }

type diffuseLight struct {
	emit Texture
}

func NewDiffuseLight(emit Texture) *diffuseLight {
	return &diffuseLight{emit}
}

func (d *diffuseLight) scatter(r *Ray, hr *hitRecord) (bool, *Ray, Color) {
	return false, nil, Color{}
}

func (d *diffuseLight) emitted(u, v Double, p Vec3) Color {
	return d.emit.value(u, v, p)
}

type isotropic struct {
	albedo Texture
}

func NewIsotropic(albedo Texture) *isotropic {
	return &isotropic{albedo}
}

func (i *isotropic) scatter(r *Ray, hr *hitRecord) (bool, *Ray, Color) {
	return true, NewRay(hr.p, randomOnUnitShpere(), r.Time), i.albedo.value(hr.u, hr.v, hr.p)
}

func (i *isotropic) emitted(u, v Double, p Vec3) Color { return black }
