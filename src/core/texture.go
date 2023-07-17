package core

import (
	"image"
	"math"
)

type Texture interface {
	Value(u, v float64, p *Vec3) *Vec3
}

type SolidColor struct {
	Color *Vec3
}

func NewSolidColor(rgb *Vec3) *SolidColor {
	return &SolidColor{rgb}
}

func NewSolidColorRGB(r, g, b double) *SolidColor {
	return &SolidColor{NewVec3(r, g, b)}
}

func (sc *SolidColor) Value(u, v float64, p *Vec3) *Vec3 {
	return sc.Color
}

type CheckerTexture struct {
	odd, even Texture
}

func NewCheckerTexture(odd, even Texture) *CheckerTexture {
	return &CheckerTexture{odd, even}
}

func (ct *CheckerTexture) Value(u, v float64, p *Vec3) *Vec3 {
	sines := math.Sin(10*p.X) * math.Sin(10*p.Y) * math.Sin(10*p.Z)
	if sines < 0 {
		return ct.odd.Value(u, v, p)
	} else {
		return ct.even.Value(u, v, p)
	}
}

type NoiseTexture struct {
	scale  float64
	perlin Perlin
}

func NewNoiseTexture(scale float64) *NoiseTexture {
	return &NoiseTexture{scale, *NewPerlin()}
}

func (nt *NoiseTexture) Value(u, v float64, p *Vec3) *Vec3 {
	return NewVec3(0.5, 0.5, 0.5).Mul_(1 + nt.perlin.Noise(p.Mul(nt.scale)))
}

type TurbulanceNoiseTexture struct {
	scale  float64
	period float64
	perlin Perlin
}

func NewTurbulanceNoiseTexture(scale float64, period float64) *TurbulanceNoiseTexture {
	return &TurbulanceNoiseTexture{scale, period, *NewPerlin()}
}

func (nt *TurbulanceNoiseTexture) Value(u, v float64, p *Vec3) *Vec3 {
	return NewVec3(0.5, 0.5, 0.5).Mul_(1 + math.Sin(nt.scale*p.Z+nt.period*nt.perlin.Turbulance(p)))
}

type ImageTexture struct {
	width, height int
	data          []Vec3
}

func NewImageTexture(src image.Image) *ImageTexture {
	width := src.Bounds().Dx()
	height := src.Bounds().Dy()
	data := make([]Vec3, height*width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			data[y*width+x] = Vec3{double(r) / 65535, double(g) / 65535, double(b) / 65535}
		}
	}
	return &ImageTexture{width, height, data}
}

func (it *ImageTexture) Value(u, v float64, p *Vec3) *Vec3 {
	if it.data == nil {
		return NewVec3(0, 1, 1)
	}

	u = clamp(u, 0, 1)
	v = 1 - clamp(v, 0, 1)

	i := int(u * double(it.width))
	j := int(v * double(it.height))

	if i >= it.width {
		i = it.width - 1
	}
	if j >= it.height {
		j = it.height - 1
	}

	return &it.data[j*it.width+i]
}

func clamp(x, min, max double) double {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
