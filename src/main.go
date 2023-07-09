package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/reki2000/raytrace-in-weekend-go/src/core"
)

type double = float64

func color3(r, g, b double) *core.Vec3 {
	return core.NewVec3(r, g, b)
}

func point3(x, y, z double) *core.Vec3 {
	return core.NewVec3(x, y, z)
}

func randomInUnitSphere() *core.Vec3 {
	for {
		p := core.NewVec3Random(-1, 1)
		if p.LengthSquared() >= 1 {
			continue
		}
		return p
	}
}

func rayColor(r *core.Ray, world core.ObjectList, depth int) *core.Vec3 {
	if depth <= 0 {
		return color3(0, 0, 0)
	}

	if hit, hr := world.Hit(r, 0.001, math.Inf(0)); hit {
		if scattered, scatter, attenuation := hr.Mat.Scatter(r, hr); scattered {
			return rayColor(scatter, world, depth-1).MulVec(attenuation)
		} else {
			return color3(0, 0, 0)
		}
	}

	unitDirection := r.Direction.Norm()
	t := 0.5 * (unitDirection.Y + 1.0)
	v1 := color3(1.0, 1.0, 1.0)
	v2 := color3(0.5, 0.7, 1.0)
	return v1.Mul_(1.0 - t).Add_(v2.Mul_(t))
}

func randomScene() core.ObjectList {
	world := core.ObjectList{}

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := rand.Float64()
			center := point3(double(a)+0.9*rand.Float64(), 0.2, double(b)+0.9*rand.Float64())

			if center.Sub(point3(4, 0.2, 0)).Length() > 0.9 {
				var material core.Material
				if chooseMat < 0.8 {
					// diffuse
					albedo := core.NewVec3Random(0.0, 1.0).MulVec(core.NewVec3Random(0.0, 1.0))
					material = core.NewLambertian(albedo)
				} else if chooseMat < 0.95 {
					// metal
					albedo := core.NewVec3Random(0.5, 1)
					fuzz := rand.Float64() * 0.5
					material = core.NewMetal(albedo, fuzz)
				} else {
					// glass
					material = core.NewDielectric(1.5)
				}
				world = append(world, core.NewSphere(center, 0.2, material))
			}
		}
	}

	groundMaterial := core.NewLambertian(color3(0.5, 0.5, 0.5))
	world = append(world, core.NewSphere(point3(0, -1000, 0), 1000, groundMaterial))

	material1 := core.NewDielectric(1.5)
	material2 := core.NewLambertian(color3(0.4, 0.2, 0.1))
	material3 := core.NewMetal(color3(0.7, 0.6, 0.5), 0.0)
	world = append(world,
		core.NewSphere(point3(0, 1, 0), 1.0, material1),
		core.NewSphere(point3(-4, 1, 0), 1.0, material2),
		core.NewSphere(point3(4, 1, 0), 1.0, material3),
	)

	return world
}

func testScene() core.ObjectList {
	world := core.ObjectList{
		core.NewSphere(point3(0, 0, -1), 0.5, core.NewLambertian(color3(0.1, 0.2, 0.5))),
		core.NewSphere(point3(0, -100.5, -1), 100, core.NewLambertian(color3(0.8, 0.8, 0.0))),
		core.NewSphere(point3(1, 0, -1), 0.5, core.NewMetal(color3(0.8, 0.6, 0.2), 0.3)),
		core.NewSphere(point3(-1, 0, -1), 0.5, core.NewDielectric(1.5)),
		core.NewSphere(point3(-1, 0, -1), -0.45, core.NewDielectric(1.5)),
	}
	return world
}

func main() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	// canvas settings
	aspectRatio := 16.0 / 9.0
	imageWidth := 384
	imageHeight := int(double(imageWidth) / aspectRatio)

	// camera settings
	vfovDegree := 90.0
	lookFrom := point3(13, 2, 3)
	lookAt := point3(0, 0, 0)
	vup := point3(0, 1, 0)
	distToFocus := 10.0
	aperture := 0.1
	camera := core.NewCamera(aspectRatio, vfovDegree, aperture, distToFocus, lookFrom, lookAt, vup)

	// locate objetcs
	world := randomScene()

	// ray tracing settings
	samples := 32
	maxDepth := 10

	// rendering
	buffer := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			pixelColor := color3(0, 0, 0)

			for s := 0; s < samples; s++ {
				u := (double(i) + rand.Float64()) / double(imageWidth-1)
				v := (double(j) + rand.Float64()) / double(imageHeight-1)
				r := camera.GetRay(u, v)

				pixelColor.Add_(rayColor(r, world, maxDepth))
			}

			r, g, b := antiAlias(pixelColor, samples)
			buffer.Set(i, imageHeight-j-1, color.RGBA{r, g, b, 255})
		}
	}

	png.Encode(os.Stdout, buffer)
}

func antiAlias(c *core.Vec3, samples int) (uint8, uint8, uint8) {
	r := uint8(math.Min(math.Sqrt(c.X/double(samples)), 0.999) * 256.0)
	g := uint8(math.Min(math.Sqrt(c.Y/double(samples)), 0.999) * 256.0)
	b := uint8(math.Min(math.Sqrt(c.Z/double(samples)), 0.999) * 256.0)
	return r, g, b
}
