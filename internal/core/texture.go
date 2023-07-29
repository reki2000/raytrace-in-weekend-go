package core

import (
	"image"
	"math"
)

type Texture interface {
	value(u, v Double, p Vec3) Color
}

// SolidColor
type SolidColor struct {
	Color Color
}

func NewSolidColor(rgb Color) *SolidColor {
	return &SolidColor{rgb}
}

func (sc *SolidColor) value(u, v Double, p Vec3) Color {
	return sc.Color
}

// Checker
type CheckerTexture struct {
	odd, even Texture
}

func NewCheckerTexture(odd, even Texture) *CheckerTexture {
	return &CheckerTexture{odd, even}
}

func (ct *CheckerTexture) value(u, v Double, p Vec3) Color {
	sines := math.Sin(10*p.x) * math.Sin(10*p.y) * math.Sin(10*p.z)
	if sines < 0 {
		return ct.odd.value(u, v, p)
	} else {
		return ct.even.value(u, v, p)
	}
}

// Noise
type NoiseTexture struct {
	scale  Double
	perlin perlinNoise
}

func NewNoiseTexture(scale Double) *NoiseTexture {
	return &NoiseTexture{scale, *newPerlinNoise()}
}

func (nt *NoiseTexture) value(u, v Double, p Vec3) Color {
	return NewColor(0.5, 0.5, 0.5).Mul(1 + nt.perlin.noise(p.mul(nt.scale)))
}

// TurbulanceNoise (marble)
type TurbulanceNoiseTexture struct {
	scale  Double
	period Double
	perlin perlinNoise
}

func NewTurbulanceNoiseTexture(scale Double, period Double) *TurbulanceNoiseTexture {
	return &TurbulanceNoiseTexture{scale, period, *newPerlinNoise()}
}

func (nt *TurbulanceNoiseTexture) value(u, v Double, p Vec3) Color {
	return NewColor(0.5, 0.5, 0.5).Mul(1 + math.Sin(nt.scale*p.z+nt.period*nt.perlin.turbulance(p)))
}

// Image
type ImageTexture struct {
	width, height int
	data          []Color
}

func NewImageTexture(src image.Image) *ImageTexture {
	width := src.Bounds().Dx()
	height := src.Bounds().Dy()
	data := make([]Color, height*width)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			data[y*width+x] = NewColor(Double(r)/65535, Double(g)/65535, Double(b)/65535)
		}
	}
	return &ImageTexture{width, height, data}
}

func (it *ImageTexture) value(u, v Double, p Vec3) Color {
	if it.data == nil {
		return NewColor(0, 1, 1)
	}

	u = clamp(u, 0, 1)
	v = 1 - clamp(v, 0, 1)

	i := int(u * Double(it.width))
	j := int(v * Double(it.height))

	if i >= it.width {
		i = it.width - 1
	}
	if j >= it.height {
		j = it.height - 1
	}

	return it.data[j*it.width+i]
}

func clamp(x, min, max Double) Double {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}
