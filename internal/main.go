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
	"runtime"
	"sync"
	"time"

	"github.com/reki2000/raytrace-in-weekend-go/internal/core"
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

func loadImage(filename string) image.Image {
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

	return image
}

func main() {
	scene := flag.String("scene", "random", "scene type")
	maxDepth := flag.Int("depth", 10, "depth for ray tracing")
	imageWidth := flag.Int("width", 384, "image width")
	samples := flag.Int("samples", 16, "samples per pixel")
	aspectRatio := flag.Float64("aspect", 16.0/9.0, "aspect ratio")
	threads := flag.Int("threads", runtime.NumCPU(), "number of parallel threads")
	flag.Parse()

	seed := time.Now().UnixNano()
	rand.Seed(seed)

	// camera settings
	vfovDegree := 90.0
	lookFrom := p3(13, 2, 3)
	lookAt := p3(0, 0, 0)
	vup := p3(0, 1, 0)
	distToFocus := 10.0
	aperture := 0.1
	time0, time1 := 0.0, 1.0
	camera := core.NewCamera(*aspectRatio, vfovDegree, aperture, distToFocus, lookFrom, lookAt, vup, time0, time1)

	// locate objetcs
	world := testScene()
	backGround := c3(0.7, 0.8, 1.0)
	if *scene == "random" {
		world = randomScene()
	} else if *scene == "light" {
		world = lightScene()
		backGround = c3(0.0, 0.0, 0.0)
		camera = core.NewCamera(*aspectRatio, vfovDegree, 0.01, 10, p3(40, 5, 4), p3(0, 3, 0), vup, time0, time1)
	} else if *scene == "cornell" {
		world = cornellBoxScene(false)
		backGround = c3(0.0, 0.0, 0.0)
		camera = core.NewCamera(*aspectRatio, 40.0, 0.0, 1, p3(278, 278, -800), p3(278, 278, 0), vup, time0, time1)
	} else if *scene == "smoke" {
		world = cornellBoxScene(true)
		backGround = c3(0.0, 0.0, 0.0)
		camera = core.NewCamera(*aspectRatio, 40.0, 0.0, 1, p3(278, 278, -800), p3(278, 278, 0), vup, time0, time1)
	} else if *scene == "final" {
		world = finalScene()
		backGround = c3(0.0, 0.0, 0.0)
		*aspectRatio = 1.0
		camera = core.NewCamera(*aspectRatio, 28.0, 0.0, 1, p3(500, 230, -800), p3(300, 260, 0), vup, time0, time1)
	}
	//fmt.Fprintf(os.Stderr, "world: %s\n", world)

	// ray tracing settings
	imageHeight := int(Double(*imageWidth) / *aspectRatio)

	// rendering
	buffer := image.NewRGBA(image.Rect(0, 0, *imageWidth, imageHeight))

	wg := sync.WaitGroup{}
	limit := make(chan int, *threads)

	for j := 0; j < imageHeight; j++ {
		limit <- 1
		wg.Add(1)
		go func(j int) {
			defer func() { <-limit; wg.Done() }()

			for i := 0; i < *imageWidth; i++ {
				pixelColor := c3(0, 0, 0)

				for s := 0; s < *samples; s++ {
					u := (Double(i) + randomDouble()) / Double(*imageWidth-1)
					v := (Double(j) + randomDouble()) / Double(imageHeight-1)
					r := camera.GetRay(u, v)

					pixelColor = pixelColor.Add(r.Color(backGround, world, *maxDepth))
				}

				r, g, b := antiAlias(pixelColor, *samples)
				buffer.Set(i, imageHeight-j-1, color.RGBA{r, g, b, 255})
			}
		}(j)
	}
	wg.Wait()

	png.Encode(os.Stdout, buffer)
}

func antiAlias(c core.Color, samples int) (uint8, uint8, uint8) {
	r := uint8(math.Min(math.Sqrt(c.R/Double(samples)), 0.999) * 256.0)
	g := uint8(math.Min(math.Sqrt(c.G/Double(samples)), 0.999) * 256.0)
	b := uint8(math.Min(math.Sqrt(c.B/Double(samples)), 0.999) * 256.0)
	return r, g, b
}
