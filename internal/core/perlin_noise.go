package core

import (
	"math"
	"math/rand"
)

type perlinNoise struct {
	pointCount      int
	permX           []int
	permY           []int
	permZ           []int
	randomVec       []Vec3
	turbulanceDepth int
}

func newPerlinNoise() *perlinNoise {
	pointCount := 256
	series := generateSeries(pointCount)
	permX := permutate(series)
	permY := permutate(series)
	permZ := permutate(series)

	randomVec := make([]Vec3, pointCount)

	for i := 0; i < pointCount; i++ {
		randomVec[i] = NewVec3Random(-1, 1).norm()
	}

	turbulanceDepth := 7

	return &perlinNoise{pointCount, permX, permY, permZ, randomVec, turbulanceDepth}
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

func perlinInterporlate(c [2][2][2]Vec3, u, v, w Double) Double {
	// 3d hermite cubic
	uu := u * u * (3 - 2*u)
	vv := v * v * (3 - 2*v)
	ww := w * w * (3 - 2*w)

	accum := 0.0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				ii := Double(i)
				jj := Double(j)
				kk := Double(k)
				weight := vec3(u-ii, v-jj, w-kk)
				accum += (ii*uu + (1-ii)*(1-uu)) *
					(jj*vv + (1-jj)*(1-vv)) *
					(kk*ww + (1-kk)*(1-ww)) * weight.dot(c[i][j][k])
			}
		}
	}
	return accum
}

func (perlin perlinNoise) noise(p Vec3) Double {
	u := p.x - math.Floor(p.x)
	v := p.y - math.Floor(p.y)
	w := p.z - math.Floor(p.z)

	// 3d hermite cubic
	u = u * u * (3 - 2*u)
	v = v * v * (3 - 2*v)
	w = w * w * (3 - 2*w)

	c := [2][2][2]Vec3{}

	i := int(math.Floor(p.x))
	j := int(math.Floor(p.y))
	k := int(math.Floor(p.z))

	for di := 0; di < 2; di++ {
		for dj := 0; dj < 2; dj++ {
			for dk := 0; dk < 2; dk++ {
				c[di][dj][dk] = perlin.randomVec[perlin.permX[(i+di)&255]^
					perlin.permY[(j+dj)&255]^
					perlin.permZ[(k+dk)&255]]
			}
		}
	}

	return perlinInterporlate(c, u, v, w)

}

func (perlin *perlinNoise) turbulance(p Vec3) Double {
	accum := 0.0
	tempP := p
	weight := 1.0
	for i := 0; i < perlin.turbulanceDepth; i++ {
		accum += weight * perlin.noise(tempP)
		weight *= 0.5
		tempP = tempP.mul(2)
	}
	return math.Abs(accum)
}
