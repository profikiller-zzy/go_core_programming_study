package mandelbrot

import (
	"image"
	"testing"
)

func BenchmarkNaiveRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		naiveRender(img)
	}
}

func BenchmarkParallelRender(b *testing.B) {
	for i := 0; i < b.N; i++ {
		img := image.NewRGBA(image.Rect(0, 0, width, height))
		parallelRender(img)
	}
}
