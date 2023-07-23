# RayTraceInWeekend Go

This repository provides a Go language implementation of Peter Shirley's [Ray Tracing in One Weekend](https://raytracing.github.io/). A [Japanese version](https://inzkyk.xyz/ray_tracing_in_one_weekend/) of the original book is also available.

![image](https://github.com/reki2000/raytrace-in-weekend-go/assets/2533597/a879e6e4-973b-4297-85ac-bccd83c0da10)
![image](https://github.com/reki2000/raytrace-in-weekend-go/assets/2533597/cf82a52e-8b43-4c1c-bc8d-009ddadc941f)

_rendered in 40-60sec with i5-8250_

## How to Run

```
go run ./src > output.png
```

This takes some options:

```
 -aspect float
        aspect ratio (default 1.7777777777777777)
  -depth int
        depth for ray tracing (default 10)
  -samples int
        samples per pixel (default 16)
  -scene string
        scene type (default "random")
  -threads int
        number of parallel threads (default 8)
  -width int
        image width (default 384)
```

## Preset Renderers

```
make scene_final
```
![image](https://github.com/reki2000/raytrace-in-weekend-go/assets/2533597/5975dd16-e84c-491c-b5f7-40493a4e48e0)

SOme other preset renderers are available as make target:
- `scene_test`, `scene_random`, `scene_light`, `scene_cornell`, `scene_somke` 

## Development

This is designed to render one image per execution. To simplify the process of repeatedly compiling and rendering images, we recommend using the `make loop` command. This command is designed to work best with a .png viewer capable of detecting changes in the displayed image and reloading it automatically.

```shell
make loop
```

Please note that you need to have `make` installed in your system for the above command to work.

