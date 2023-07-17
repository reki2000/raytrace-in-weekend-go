package core

import "math/rand"

type Perlin struct {
	pointCount int
	random     []double
	permX      []int
	permY      []int
	permZ      []int
}

func NewPerlin() *Perlin {
	pointCount := 256
	random := make([]double, pointCount)
	for i := 0; i < pointCount; i++ {
		random[i] = rand.Float64()
	}

	series := generateSeries(pointCount)
	permX := permutate(series)
	permY := permutate(series)
	permZ := permutate(series)

	return &Perlin{pointCount, random, permX, permY, permZ}
}

func generateSeries(count int) []int {
	series := make([]int, count)
	for i := 0; i < count; i++ {
		series[i] = i
	}
	return series
}

func permutate(slice []int) []int {
	count := len(slice)
	for i := count - 1; i > 0; i-- {
		from := rand.Intn(i + 1)
		slice[i], slice[from] = slice[from], slice[i]
	}
	return slice
}

func (perlin *Perlin) Noise(p *Vec3) double {
	i := int(4*p.X) & 255
	j := int(4*p.Y) & 255
	k := int(4*p.Z) & 255
	return perlin.random[perlin.permX[i]^perlin.permY[j]^perlin.permZ[k]]
}
