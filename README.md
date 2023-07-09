# RayTraceInWeekend Go

This repository provides a Go language implementation of Peter Shirley's [Ray Tracing in One Weekend](https://raytracing.github.io/). A [Japanese version](https://inzkyk.xyz/ray_tracing_in_one_weekend/) of the original book is also available.

## How to Run

This is designed to render one image per execution. To simplify the process of repeatedly compiling and rendering images, we recommend using the `make loop` command. This command is designed to work best with a .png viewer capable of detecting changes in the displayed image and reloading it automatically.

```shell
make loop
```

Please note that you need to have `make` installed in your system for the above command to work.
