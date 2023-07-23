package core

import (
	"math"
)

type Vec3 struct {
	x, y, z Double
}

func (v Vec3) e(axis int) Double {
	switch axis {
	case 0:
		return v.x
	case 1:
		return v.y
	case 2:
		return v.z
	}
	panic("invalid axis")
}

func (v Vec3) newE(axis int, value Double) Vec3 {
	switch axis {
	case 0:
		return vec3(value, v.y, v.z)
	case 1:
		return vec3(v.x, value, v.z)
	case 2:
		return vec3(v.x, v.y, value)
	}
	panic("invalid axis")
}

func (v Vec3) inv() Vec3          { return vec3(-v.x, -v.y, -v.z) }
func (v Vec3) add(x Vec3) Vec3    { return vec3(v.x+x.x, v.y+x.y, v.z+x.z) }
func (v Vec3) sub(x Vec3) Vec3    { return vec3(v.x-x.x, v.y-x.y, v.z-x.z) }
func (v Vec3) mul(n Double) Vec3  { return vec3(v.x*n, v.y*n, v.z*n) }
func (v Vec3) mulVec(x Vec3) Vec3 { return vec3(v.x*x.x, v.y*x.y, v.z*x.z) }
func (v Vec3) div(n Double) Vec3  { return vec3(v.x/n, v.y/n, v.z/n) }

func (v Vec3) lengthSquared() Double { return v.x*v.x + v.y*v.y + v.z*v.z }
func (v Vec3) length() Double        { return math.Sqrt(v.lengthSquared()) }

func (v Vec3) dot(x Vec3) Double { return v.x*x.x + v.y*x.y + v.z*x.z }

func (v Vec3) cross(x Vec3) Vec3 {
	return vec3(
		v.y*x.z-v.z*x.y,
		v.z*x.x-v.x*x.z,
		v.x*x.y-v.y*x.x,
	)
}

func (v Vec3) norm() Vec3 {
	return v.div(v.length())
}

func (v Vec3) reflect(x Vec3) Vec3 {
	return v.sub(x.mul(v.dot(x) * 2.0))
}

func (v Vec3) refract(n Vec3, etaIOverEtaT Double) Vec3 {
	cosTheta := v.inv().dot(n)
	rOutParallel := v.add(n.mul(cosTheta)).mul(etaIOverEtaT)
	rOutPerp := n.mul(-math.Sqrt(1.0 - rOutParallel.lengthSquared()))
	return rOutParallel.add(rOutPerp)
}

func NewVec3(x, y, z Double) Vec3 {
	return Vec3{x, y, z}
}

func vec3(x, y, z Double) Vec3 {
	return Vec3{x, y, z}
}

func NewVec3Random(min, max Double) Vec3 {
	return vec3(
		randomDoubleW(min, max),
		randomDoubleW(min, max),
		randomDoubleW(min, max),
	)
}

func (v Vec3) Add(x Vec3) Vec3 { return vec3(v.x+x.x, v.y+x.y, v.z+x.z) }
func (v Vec3) Sub(x Vec3) Vec3 { return vec3(v.x-x.x, v.y-x.y, v.z-x.z) }
func (v Vec3) Length() Double  { return math.Sqrt(v.lengthSquared()) }
