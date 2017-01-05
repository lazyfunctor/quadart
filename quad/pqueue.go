package quad

import "strconv"

type Rectangle struct {
	topLeftX     int
	topLeftY     int
	bottomRightX int
	bottomRightY int
}

func NewRectangle(topLeftX int, topLeftY int, bottomRightX int, bottomRightY int) Rectangle {
	return Rectangle{topLeftX: topLeftX, topLeftY: topLeftY, bottomRightX: bottomRightX, bottomRightY: bottomRightY}
}

func (r Rectangle) SplitTo4() (Rectangle, Rectangle, Rectangle, Rectangle) {
	xMid := r.topLeftX + (r.bottomRightX-r.topLeftX)/2
	yMid := r.topLeftY + (r.bottomRightY-r.topLeftY)/2
	return NewRectangle(r.topLeftX, r.topLeftY, xMid, yMid),
		NewRectangle(xMid+1, r.topLeftY, r.bottomRightX, yMid),
		NewRectangle(r.topLeftX, yMid+1, xMid, r.bottomRightY),
		NewRectangle(xMid+1, yMid+1, r.bottomRightX, r.bottomRightY)
}

func (r Rectangle) String() string {
	topLeftX := strconv.Itoa(r.topLeftX)
	topLeftY := strconv.Itoa(r.topLeftY)
	bottomRightX := strconv.Itoa(r.bottomRightX)
	bottomRightY := strconv.Itoa(r.bottomRightY)
	return "Rectangle(" + topLeftX + "," + topLeftY + ") -- (" + bottomRightX + ", " + bottomRightY + ")"
}

func (r Rectangle) Area() uint32 {
	width := r.bottomRightX - r.topLeftX + 1
	height := r.bottomRightY - r.topLeftY + 1
	return uint32(width) * uint32(height)
}

type RectangleDistance struct {
	Rect     Rectangle
	Distance float64
}

type DistanceHeap []*RectangleDistance

func (h DistanceHeap) Len() int {
	return len(h)
}

func (h DistanceHeap) Less(i, j int) bool {
	return h[i].Distance < h[j].Distance
}

func (h DistanceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *DistanceHeap) Push(x interface{}) {
	*h = append(*h, x.(*RectangleDistance))
}

func (h *DistanceHeap) Pop() interface{} {
	oldh := *h
	x := oldh[len(oldh)-1]
	newh := oldh[0 : len(oldh)-1]
	*h = newh
	return x
}
