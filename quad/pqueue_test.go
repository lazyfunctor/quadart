package quad

import "testing"
import _ "reflect"

import "container/heap"

func TestMarshalBool1(t *testing.T) {
	var dh DistanceHeap
	heap.Init(&dh)
	heap.Push(&dh, &RectangleDistance{
		Distance: 1.7,
		Rect:     Rectangle{topLeftX: 100, topLeftY: 100, bottomRightX: 200, bottomRightY: 200}})
	heap.Push(&dh, &RectangleDistance{
		Distance: 1.9,
		Rect:     Rectangle{topLeftX: 0, topLeftY: 0, bottomRightX: 100, bottomRightY: 100}})

	heap.Push(&dh, &RectangleDistance{
		Distance: 1.3,
		Rect:     Rectangle{topLeftX: 200, topLeftY: 200, bottomRightX: 300, bottomRightY: 300}})
	rd1 := heap.Pop(&dh).(*RectangleDistance)
	rd2 := heap.Pop(&dh).(*RectangleDistance)
	rd3 := heap.Pop(&dh).(*RectangleDistance)

	// fmt.Println(rd1.rect)

	if rd1.Rect.topLeftX != 0 && rd1.Rect.topLeftY != 0 {
		t.Error("Problem with priority queue")
	}
	if rd2.Rect.topLeftX != 100 && rd2.Rect.topLeftY != 100 {
		t.Error("Problem with priority queue")
	}
	if rd3.Rect.topLeftX != 200 && rd3.Rect.topLeftY != 200 {
		t.Error("Problem with priority queue")
	}

}
