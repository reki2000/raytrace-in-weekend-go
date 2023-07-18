package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"math"
	"math/rand"
	"os"
	"time"

	"github.com/reki2000/raytrace-in-weekend-go/src/core"
)

type Double = core.Double

func randomDouble() Double {
	return rand.Float64()
}

func c3(r, g, b Double) core.Color {
	return core.NewColor(r, g, b)
}

func p3(x, y, z Double) core.Vec3 {
	return core.NewVec3(x, y, z)
}

func randomScene() core.ObjectList {
	world := core.ObjectList{}

	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMat := randomDouble()
			center := p3(Double(a)+0.9*randomDouble(), 0.2, Double(b)+0.9*randomDouble())

			if center.Sub(p3(4, 0.2, 0)).Length() > 0.9 {
				var material core.Material
				if chooseMat < 0.8 {
					// diffuse
					albedo := core.NewColorRandom(0.0, 1.0).MulVec(core.NewColorRandom(0.0, 1.0))
					material = core.NewLambertian(core.NewSolidColor(albedo))
				} else if chooseMat < 0.95 {
					// metal
					albedo := core.NewColorRandom(0.5, 1)
					fuzz := randomDouble() * 0.5
					material = core.NewMetal(albedo, fuzz)
				} else {
					// glass
					material = core.NewDielectric(1.5)
				}
				world = append(world, core.NewSphere(center, 0.2, material))
			}
		}
	}
	bvh := core.NewBvhNode(world, 0.0, 1.0)
	world = core.ObjectList{bvh}

	groundMaterial := core.NewLambertian(core.NewSolidColor(c3(0.5, 0.5, 0.5)))
	world = append(world, core.NewSphere(p3(0, -1000, 0), 1000, groundMaterial))

	material1 := core.NewDielectric(1.5)
	material2 := core.NewLambertian(core.NewSolidColor(c3(0.4, 0.2, 0.1)))
	material3 := core.NewMetal(c3(0.7, 0.6, 0.5), 0.0)
	world = append(world,
		core.NewSphere(p3(0, 1, 0), 1.0, material1),
		core.NewSphere(p3(-4, 1, 0), 1.0, material2),
		core.NewSphere(p3(4, 1, 0), 1.0, material3),
	)

	return world
}

func testScene() core.ObjectList {
	checkerTextture := core.NewCheckerTexture(core.NewSolidColor(c3(0.2, 0.3, 0.1)), core.NewSolidColor(c3(0.9, 0.9, 0.9)))
	//noiseTextture := core.NewNoiseTexture(5.0)
	marbleTextture := core.NewTurbulanceNoiseTexture(5.0, 20)
	//blueTexture := core.NewSolidColorRGB(0.1, 0.2, 0.5)

	file, err := os.Open("resource/earthmap.jpg")
	if err != nil {
		pwd, _ := os.Getwd()
		fmt.Fprintf(os.Stderr, "pwd: %s", pwd)
		panic(err)
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	earthTexture := core.NewImageTexture(image)

	world := core.ObjectList{
		core.NewSphere(p3(0, -100.5, -1), 100, core.NewLambertian(checkerTextture)),

		core.NewSphere(p3(0, 0.15, 0), 0.7, core.NewLambertian(earthTexture)),

		core.NewSphere(p3(0, 0, -1.2), 0.5, core.NewMetal(c3(0.8, 0.6, 0.2), 0.2)),

		core.NewSphere(p3(-1, 0.5, 1), 1.0, core.NewDielectric(1.5)),
		core.NewSphere(p3(-1, 0.5, 1), -0.9, core.NewDielectric(1.5)),

		core.NewMovingSphere(p3(2, -0.3, 1), p3(2, -0.3, -1), 0.2, core.NewLambertian(core.NewSolidColor(c3(0.0, 0.8, 0.8))), -4.0, 5.0),

		core.NewSphere(p3(-4, 0, 1.4), 0.5, core.NewLambertian(marbleTextture)),
	}

	bvh := core.NewBvhNode(world, 0.0, 1.0)
	return core.ObjectList{bvh}
}

func main() {
	scene := flag.String("scene", "random", "scene type")
	flag.Parse()

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	// canvas settings
	aspectRatio := 16.0 / 9.0
	imageWidth := 384
	imageHeight := int(Double(imageWidth) / aspectRatio)

	// camera settings
	vfovDegree := 90.0
	lookFrom := p3(13, 2, 3)
	lookAt := p3(0, 0, 0)
	vup := p3(0, 1, 0)
	distToFocus := 10.0
	aperture := 0.1
	time0, time1 := 0.0, 1.0
	camera := core.NewCamera(aspectRatio, vfovDegree, aperture, distToFocus, lookFrom, lookAt, vup, time0, time1)

	// locate objetcs
	world := testScene()
	if *scene == "random" {
		world = randomScene()
	}
	//fmt.Fprintf(os.Stderr, "world: %s\n", world)

	// ray tracing settings
	samples := 32
	maxDepth := 10

	// rendering
	buffer := image.NewRGBA(image.Rect(0, 0, imageWidth, imageHeight))

	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			pixelColor := c3(0, 0, 0)

			for s := 0; s < samples; s++ {
				u := (Double(i) + randomDouble()) / Double(imageWidth-1)
				v := (Double(j) + randomDouble()) / Double(imageHeight-1)
				r := camera.GetRay(u, v)

				pixelColor = pixelColor.Add(core.RayColor(r, world, maxDepth))
			}

			r, g, b := antiAlias(pixelColor, samples)
			buffer.Set(i, imageHeight-j-1, color.RGBA{r, g, b, 255})
		}
	}

	png.Encode(os.Stdout, buffer)
}

func antiAlias(c core.Color, samples int) (uint8, uint8, uint8) {
	r := uint8(math.Min(math.Sqrt(c.R/Double(samples)), 0.999) * 256.0)
	g := uint8(math.Min(math.Sqrt(c.G/Double(samples)), 0.999) * 256.0)
	b := uint8(math.Min(math.Sqrt(c.B/Double(samples)), 0.999) * 256.0)
	return r, g, b
}
