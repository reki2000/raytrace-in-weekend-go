package main

import "github.com/reki2000/raytrace-in-weekend-go/internal/core"

func testScene() core.ObjectList {
	checkerTextture := core.NewCheckerTexture(core.NewSolidColor(c3(0.2, 0.3, 0.1)), core.NewSolidColor(c3(0.9, 0.9, 0.9)))
	//noiseTextture := core.NewNoiseTexture(5.0)
	marbleTextture := core.NewTurbulanceNoiseTexture(5.0, 20)
	//blueTexture := core.NewSolidColorRGB(0.1, 0.2, 0.5)

	earthTexture := core.NewImageTexture(loadImage("resource/earthmap.jpg"))

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

func lightScene() core.ObjectList {
	noiseTexture := core.NewTurbulanceNoiseTexture(4.0, 20)
	light := core.NewDiffuseLight(core.NewSolidColor(c3(4, 4, 4)))

	return core.ObjectList{
		core.NewSphere(p3(0, -1000, 0), 1000, core.NewLambertian(noiseTexture)),
		core.NewSphere(p3(0, 2, 0), 2, core.NewLambertian(noiseTexture)),
		core.NewSphere(p3(0, 7, 0), 2, light),
		core.NewXYRect(3, 5, 1, 3, -2, light),
	}
}

func cornellBoxScene(smoked bool) core.ObjectList {
	red := core.NewLambertian(core.NewSolidColor(c3(0.65, 0.05, 0.05)))
	white := core.NewLambertian(core.NewSolidColor(c3(0.73, 0.73, 0.73)))
	green := core.NewLambertian(core.NewSolidColor(c3(0.12, 0.45, 0.15)))
	light := core.NewDiffuseLight(core.NewSolidColor(c3(7, 7, 7)))

	var box1 core.Object = core.NewBox(p3(0, 0, 0), p3(165, 330, 165), white)
	box1 = core.NewRotateY(box1, 15)
	box1 = core.NewTranslate(box1, p3(265, 0, 295))

	var box2 core.Object = core.NewBox(p3(0, 0, 0), p3(165, 165, 165), white)
	box2 = core.NewRotateY(box2, -18)
	box2 = core.NewTranslate(box2, p3(130, 0, 65))

	objs := core.ObjectList{
		core.NewYZRect(0, 555, 0, 555, 555, green),
		core.NewYZRect(0, 555, 0, 555, 0, red),
		core.NewXZRect(113, 443, 127, 432, 554, light),
		core.NewXZRect(0, 555, 0, 555, 0, white),
		core.NewXZRect(0, 555, 0, 555, 555, white),
		core.NewXYRect(0, 555, 0, 555, 555, white),
	}

	if smoked {
		objs = append(objs,
			core.NewConstantMedium(box1, 0.01, core.NewSolidColor(c3(0, 0, 0))),
			core.NewConstantMedium(box2, 0.01, core.NewSolidColor(c3(1, 1, 1))))
	} else {
		objs = append(objs, box1, box2)
	}

	return objs
}

func finalScene() core.ObjectList {
	var boxes1 core.ObjectList
	ground := core.NewLambertian(core.NewSolidColor(c3(0.48, 0.83, 0.53)))

	const boxesPerSide = 20
	for i := 0; i < boxesPerSide; i++ {
		for j := 0; j < boxesPerSide; j++ {
			w := 100.0
			x0 := -1000.0 + Double(i)*w
			z0 := -1000.0 + Double(j)*w
			y0 := 0.0
			x1 := x0 + w
			y1 := randomDouble()*100 + 1
			z1 := z0 + w

			boxes1 = append(boxes1, core.NewBox(p3(x0, y0, z0), p3(x1, y1, z1), ground))
		}
	}

	light := core.NewDiffuseLight(core.NewSolidColor(c3(7, 7, 7)))
	lightRect := core.NewXZRect(123, 423, 147, 412, 554, light)

	center1 := p3(400, 400, 200)
	center2 := center1.Add(p3(30, 0, 0))
	movingSphereMaterial := core.NewLambertian(core.NewSolidColor(c3(0.7, 0.3, 0.1)))
	movingShere := core.NewMovingSphere(center1, center2, 50, movingSphereMaterial, 0, 1)

	boundary1 := core.NewSphere(p3(360, 150, 45), 50, core.NewDielectric(1.5))
	sphere1 := core.NewConstantMedium(boundary1, 0.2, core.NewSolidColor(c3(0.2, 0.4, 0.9)))

	boundary2 := core.NewSphere(p3(0, 0, 0), 5000, core.NewDielectric(1.5))
	sphere2 := core.NewConstantMedium(boundary1, 0.0001, core.NewSolidColor(c3(1, 1, 1)))

	boxes2 := core.ObjectList{}
	while := core.NewLambertian(core.NewSolidColor(c3(0.73, 0.73, 0.73)))
	for i := 0; i < 1000; i++ {
		boxes2 = append(boxes2, core.NewSphere(core.NewVec3Random(0, 165), 10, while))
	}

	objs := core.ObjectList{
		core.NewBvhNode(boxes1, 0.0, 1.0),
		lightRect,
		movingShere,
		core.NewSphere(p3(260, 150, 45), 50, core.NewDielectric(1.5)),
		core.NewSphere(p3(0, 150, 145), 50, core.NewMetal(c3(0.8, 0.8, 0.9), 10.0)),
		boundary1, sphere1,
		boundary2, sphere2,
		core.NewSphere(p3(400, 200, 400), 100,
			core.NewLambertian(core.NewImageTexture(loadImage("earthmap.jpg")))),
		core.NewSphere(p3(220, 280, 300), 80, core.NewLambertian(core.NewTurbulanceNoiseTexture(5, 20))),
		core.NewTranslate(core.NewRotateY(core.NewBvhNode(boxes2, 0, 1), 15), p3(-100, 270, 395)),
	}

	return objs

}
