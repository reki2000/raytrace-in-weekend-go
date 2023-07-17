package core

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	x, y, z double
}

func (v Vec3) e(axis int) double {
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

func (v Vec3) inv() Vec3          { return vec3(-v.x, -v.y, -v.z) }
func (v Vec3) add(x Vec3) Vec3    { return vec3(v.x+x.x, v.y+x.y, v.z+x.z) }
func (v Vec3) sub(x Vec3) Vec3    { return vec3(v.x-x.x, v.y-x.y, v.z-x.z) }
func (v Vec3) mul(n double) Vec3  { return vec3(v.x*n, v.y*n, v.z*n) }
func (v Vec3) mulVec(x Vec3) Vec3 { return vec3(v.x*x.x, v.y*x.y, v.z*x.z) }
func (v Vec3) div(n double) Vec3  { return vec3(v.x/n, v.y/n, v.z/n) }

func (v Vec3) lengthSquared() double { return v.x*v.x + v.y*v.y + v.z*v.z }
func (v Vec3) length() double        { return math.Sqrt(v.lengthSquared()) }

func (v Vec3) dot(x Vec3) double { return v.x*x.x + v.y*x.y + v.z*x.z }

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

func (v Vec3) refract(n Vec3, etaIOverEtaT double) Vec3 {
	cosTheta := v.inv().dot(n)
	rOutParallel := v.add(n.mul(cosTheta)).mul(etaIOverEtaT)
	rOutPerp := n.mul(-math.Sqrt(1.0 - rOutParallel.lengthSquared()))
	return rOutParallel.add(rOutPerp)
}

func NewVec3(x, y, z double) Vec3 {
	return Vec3{x, y, z}
}

func vec3(x, y, z double) Vec3 {
	return Vec3{x, y, z}
}

func NewVec3Random(min, max double) Vec3 {
	w := max - min
	return vec3(
		rand.Float64()*w+min,
		rand.Float64()*w+min,
		rand.Float64()*w+min,
	)
}

func (v Vec3) Sub(x Vec3) Vec3 { return vec3(v.x-x.x, v.y-x.y, v.z-x.z) }
func (v Vec3) Length() double  { return math.Sqrt(v.lengthSquared()) }
