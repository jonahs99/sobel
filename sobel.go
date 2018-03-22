package sobel

import (
	"image"
	"image/color"
	"math"
)

// ApplySobel returns a gray-scale image of the normalized sobel magnitude
func ApplySobel(src image.Image) image.Image {
	kernelX := []float64{
		1, 0, -1,
		2, 0, -2,
		1, 0, -1}
	kernelY := []float64{
		1, 2, 1,
		0, 0, 0,
		-1, -2, -1}

	sobelX, _, _ := applyKernel(src, kernelX, 3, 3)
	sobelY, w, h := applyKernel(src, kernelY, 3, 3)

	sobelXY := make([]float64, w*h)
	min := 0.0
	max := 0.0
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			i := x + y*w
			mag := math.Sqrt(math.Pow(sobelX[i], 2) + math.Pow(sobelY[i], 2))
			sobelXY[i] = mag
			if mag < min {
				min = mag
			}
			if mag > max {
				max = mag
			}
		}
	}

	out := image.NewGray(src.Bounds())
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			v := (sobelXY[x+y*w] - min) / (max - min)
			out.SetGray(x, y, color.Gray{uint8(v * 255)})
		}
	}

	return out
}

func applyKernel(src image.Image, kernel []float64, kw int, kh int) ([]float64, int, int) {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	out := make([]float64, w*h)

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {

			sum := 0.0

			for kx := 0; kx < kw; kx++ {
				for ky := 0; ky < kh; ky++ {
					offX, offY := kx-kw/2, ky-kh/2
					srcColor := src.At(x+offX, y+offY)
					r, g, b, _ := srcColor.RGBA()
					intensity := float64(r+g+b) / (3 * 256.0)
					sum += intensity * kernel[kx+kw*ky]
				}
			}

			out[x+y*w] = sum / float64(kw) / float64(kh)

		}
	}

	return out, w, h
}
