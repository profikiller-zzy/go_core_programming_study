package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
	"sync"
)

// 对3.3节中的Mandelbrot程序进行并行改造并查看并行版本和朴素版本之间的性能差异

const (
	xmin, ymin, xmax, ymax = -2, -2, +2, +2
	width, height          = 1024, 1024
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	parallelRender(img)
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

// naiveRender 将朴素版本的代码封装到一个函数中
func naiveRender(img *image.RGBA) {
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
}

// parallelRender 将并行版本的代码封装到一个函数中
func parallelRender(img *image.RGBA) {
	var wg sync.WaitGroup
	wg.Add(height)

	for py := 0; py < height; py++ {
		go func(py int) {
			defer wg.Done()
			y := float64(py)/height*(ymax-ymin) + ymin
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				img.Set(px, py, mandelbrot(z))
			}
		}(py)
	}

	wg.Wait()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
