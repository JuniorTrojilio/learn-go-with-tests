package structmethodsandinterfaces

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	retangle := Rectangle{
		Widht:  10.0,
		Height: 10.0,
	}
	result := Perimeter(retangle)
	expect := 40.0

	if result != expect {
		t.Errorf("Result '%.2f, expect '%.2f'", result, expect)
	}
}

func TestArea(t *testing.T) {
	testsArea := []struct {
		shape  Shape
		expect float64
	}{
		{Rectangle{Widht: 12, Height: 6}, 12 * 6},
		{Circle{Radio: 10}, math.Pi * math.Pow(10, 2)},
		{Triangle{Base: 12, Height: 6}, (12 * 6) / 2},
	}

	for _, tt := range testsArea {
		result := tt.shape.Area()
		if result != tt.expect {
			t.Errorf("result '%.2f', expect '%.2f", result, tt.expect)
		}
	}
}
