package main

import (
	"container/heap"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"time"

	"github.com/MJKWoolnough/gopherjs/files"
	"github.com/gopherjs/gopherjs/js"
	"github.com/lazyfunctor/quadart/quad"
	"honnef.co/go/js/dom"
)

func splitShape(dh *quad.DistanceHeap, image image.Image, cnvs *js.Object) {
	rdMax := heap.Pop(dh).(*quad.RectangleDistance)
	fmt.Println(rdMax.Rect)
	rect1, rect2, rect3, rect4 := rdMax.Rect.SplitTo4()
	avg1, distance1 := quad.ComputeColorStats(rect1, image)
	avg2, distance2 := quad.ComputeColorStats(rect2, image)
	avg3, distance3 := quad.ComputeColorStats(rect3, image)
	avg4, distance4 := quad.ComputeColorStats(rect4, image)
	fmt.Println("avg color: distance", avg1, distance1, rect1)
	fmt.Println("avg color: distance", avg2, distance2, rect2)
	fmt.Println("avg color: distance", avg3, distance3, rect3)
	fmt.Println("avg color: distance", avg4, distance4, rect4)
	heap.Push(dh, &quad.RectangleDistance{
		Distance: distance1,
		Rect:     rect1,
	})
	heap.Push(dh, &quad.RectangleDistance{
		Distance: distance2,
		Rect:     rect2,
	})
	heap.Push(dh, &quad.RectangleDistance{
		Distance: distance3,
		Rect:     rect3,
	})
	heap.Push(dh, &quad.RectangleDistance{
		Distance: distance4,
		Rect:     rect4,
	})
	quad.RenderShape(rect1, avg1, cnvs)
	quad.RenderShape(rect2, avg2, cnvs)
	quad.RenderShape(rect3, avg3, cnvs)
	quad.RenderShape(rect4, avg4, cnvs)

	// split into four and render
}

func setup(cnvs *js.Object, image image.Image) {
	bounds := image.Bounds()
	xSize, ySize := bounds.Max.X, bounds.Max.Y
	// imgData := ctx.Call("createImageData", xSize, ySize)
	// data := imgData.Get("data") //"Get" for proerties
	cnvs.Call("setAttribute", "width", xSize)
	cnvs.Call("setAttribute", "height", ySize)
	fmt.Println(xSize, ySize)
}

func startQuad(cnvs *js.Object, image image.Image) {
	bounds := image.Bounds()
	xSize, ySize := bounds.Max.X, bounds.Max.Y
	var dh quad.DistanceHeap
	heap.Init(&dh)
	heap.Push(&dh, &quad.RectangleDistance{
		Distance: 0,
		Rect:     quad.NewRectangle(0, 0, xSize, ySize)})
	for i := 1; i <= 20000; i++ {
		splitShape(&dh, image, cnvs)
		time.Sleep(10 * time.Millisecond)
	}
}

func render() {
	document := js.Global.Get("document")
	cnvs := document.Call("getElementById", "cnvs")
	// ctx := cnvs.Call("getContext", "2d")

	// for pos := 0; pos < 256; pos++ {
	// }

	img := document.Call("getElementById", "imgInp")
	img.Call("addEventListener", "change", func() {
		go func() {
			f := img.Get("files").Index(0)
			imgFile := &dom.File{f}
			fr := files.NewFileReader(files.NewFile(imgFile))
			defer fr.Close()
			image, _, err := image.Decode(fr)
			if err != nil {
				return
			}
			setup(cnvs, image)
			startQuad(cnvs, image)

			// idx := 0
			// for j := 0; j < ySize; j++ {
			// 	for i := 0; i < xSize; i++ {
			// 		pixel := image.At(i, j)
			// 		red, green, blue, opacity := pixel.RGBA()
			// 		red8 := uint8(red / 0x101)
			// 		green8 := uint8(green / 0x101)
			// 		blue8 := uint8(blue / 0x101)
			// 		opacity8 := uint8(opacity / 0x101)
			// 		// fmt.Println(uint8(red/0x101), uint8(green/0x101), uint8(blue/0x101))
			// 		data.SetIndex(idx+0, red8)
			// 		data.SetIndex(idx+1, green8)
			// 		data.SetIndex(idx+2, blue8)
			// 		data.SetIndex(idx+3, opacity8)
			// 		idx += 4
			// 	}
			// }
			// ctx.Call("putImageData", imgData, 10, 10)

		}()
	})

}

func main() {
	render()

	/*js.Global.Set("miejs", map[string]interface{}{
		"Test": Test,
	})*/
}
