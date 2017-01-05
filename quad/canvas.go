package quad

import (
	"fmt"
	"image"
	"math"
	"strconv"

	"github.com/gopherjs/gopherjs/js"
)

func ComputeColorStats(rect Rectangle, image image.Image) ([4]uint32, float64) {
	var reds, greens, blues, alphas, pixels uint32
	for i := rect.topLeftX; i <= rect.bottomRightX; i++ {
		for j := rect.topLeftY; j <= rect.bottomRightY; j++ {
			pixel := image.At(i, j)
			red, green, blue, alpha := pixel.RGBA()
			reds += red
			blues += blue
			greens += green
			alphas += alpha
			pixels++
		}
	}
	avgColor := [4]float64{float64(reds) / float64(pixels), float64(greens) / float64(pixels),
		float64(blues) / float64(pixels), float64(alphas) / float64(pixels)}
	roundedAvgColor := [4]uint32{uint32(avgColor[0]), uint32(avgColor[1]), uint32(avgColor[2]), uint32(avgColor[3])}
	fmt.Println(avgColor, roundedAvgColor)
	var mss float64
	for x := rect.topLeftX; x <= rect.bottomRightX; x++ {
		for y := rect.topLeftY; y <= rect.bottomRightY; y++ {
			pixel := image.At(x, y)
			r, g, b, a := pixel.RGBA()
			mss += math.Pow(float64(r)-avgColor[0], 2.0) +
				math.Pow(float64(g)-avgColor[1], 2.0) +
				math.Pow(float64(b)-avgColor[2], 2.0) +
				math.Pow(float64(a)-avgColor[3], 2.0)
		}
	}
	rms := math.Sqrt(mss)
	score := -rms * math.Pow(float64(rect.Area()), 0.15)
	return roundedAvgColor, score
}

func getColorString(color [4]uint32) string {
	r := strconv.Itoa(int(uint8(color[0] / 0x101)))
	g := strconv.Itoa(int(uint8(color[1] / 0x101)))
	b := strconv.Itoa(int(uint8(color[2] / 0x101)))
	a := strconv.Itoa(int(uint8(color[3] / 0x101)))
	return "RGBA(" + r + ", " + g + ", " + b + ", " + a + ")"
}

func RenderShape(rect Rectangle, avgColor [4]uint32, cnvs *js.Object) {
	colorString := getColorString(avgColor)
	fmt.Println(colorString)
	ctx := cnvs.Call("getContext", "2d")
	ctx.Set("fillStyle", colorString)
	width := rect.bottomRightX - rect.topLeftX + 1
	height := rect.bottomRightY - rect.topLeftY + 1
	fmt.Println(rect.topLeftX, rect.topLeftY, width, height)
	ctx.Call("fillRect", rect.topLeftX, rect.topLeftY, width, height)
	ctx.Set("strokeStyle", "RGBA(0, 0, 0, 100)")
	if width*height > 10 {
		ctx.Set("lineWidth", 0.15)
		ctx.Call("strokeRect", rect.topLeftX, rect.topLeftY, width, height)
	}
}
