package core

import (
	"math"
	"math/rand"
)

type double = float64

type Vec3 struct {
	X, Y, Z double
}

func (v *Vec3) Inv() *Vec3           { return &Vec3{-v.X, -v.Y, -v.Z} }
func (v *Vec3) Add(x *Vec3) *Vec3    { return &Vec3{v.X + x.X, v.Y + x.Y, v.Z + x.Z} }
func (v *Vec3) Sub(x *Vec3) *Vec3    { return &Vec3{v.X - x.X, v.Y - x.Y, v.Z - x.Z} }
func (v *Vec3) Mul(n double) *Vec3   { return &Vec3{v.X * n, v.Y * n, v.Z * n} }
func (v *Vec3) MulVec(x *Vec3) *Vec3 { return &Vec3{v.X * x.X, v.Y * x.Y, v.Z * x.Z} }
func (v *Vec3) Div(n double) *Vec3   { return &Vec3{v.X / n, v.Y / n, v.Z / n} }

func (v *Vec3) Inv_() *Vec3           { v.X = -v.X; v.Y = -v.Y; v.Z = -v.Z; return v }
func (v *Vec3) Add_(x *Vec3) *Vec3    { v.X += x.X; v.Y += x.Y; v.Z += x.Z; return v }
func (v *Vec3) Sub_(x *Vec3) *Vec3    { v.X -= x.X; v.Y -= x.Y; v.Z -= x.Z; return v }
func (v *Vec3) Mul_(n double) *Vec3   { v.X *= n; v.Y *= n; v.Z *= n; return v }
func (v *Vec3) MulVec_(x *Vec3) *Vec3 { v.X *= x.X; v.Y *= x.Y; v.Z *= x.Z; return v }
func (v *Vec3) Div_(n double) *Vec3   { v.X /= n; v.Y /= n; v.Z /= n; return v }

func (v *Vec3) LengthSquared() double { return v.X*v.X + v.Y*v.Y + v.Z*v.Z }
func (v *Vec3) Length() double        { return math.Sqrt(v.LengthSquared()) }

func (v *Vec3) Dot(x *Vec3) double {
	return v.X*x.X + v.Y*x.Y + v.Z*x.Z
}

func (v *Vec3) Cross(x *Vec3) *Vec3 {
	return &Vec3{
		v.Y*x.Z - v.Z*x.Y,
		v.Z*x.X - v.X*x.Z,
		v.X*x.Y - v.Y*x.X,
	}
}

func (v *Vec3) Norm() *Vec3 {
	return v.Div(v.Length())
}

func (v *Vec3) Reflect(x *Vec3) *Vec3 {
	return v.Sub(x.Mul(v.Dot(x) * 2.0))
}

func (v *Vec3) Refract(n *Vec3, etaIOverEtaT double) *Vec3 {
	cosTheta := v.Inv().Dot(n)
	rOutParallel := v.Add(n.Mul(cosTheta)).Mul(etaIOverEtaT)
	rOutPerp := n.Mul(-math.Sqrt(1.0 - rOutParallel.LengthSquared()))
	return rOutParallel.Add(rOutPerp)
}

func NewVec3(x, y, z double) *Vec3 {
	return &Vec3{x, y, z}
}

func NewVec3Random(min, max double) *Vec3 {
	w := max - min
	return &Vec3{
		rand.Float64()*w + min,
		rand.Float64()*w + min,
		rand.Float64()*w + min}

}
