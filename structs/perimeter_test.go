package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{6.0, 5.0}
	result := Perimeter(rectangle)
	expected := 22.0
	if result != expected {
		t.Errorf("result '%.2f' expected '%.2f", result, expected)
	}
}

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, shape Shape, expected float64) {
		t.Helper()
		result := shape.Area()
		if result != expected {
			t.Errorf("result '%.2f' expected '%.2f", result, expected)
		}
	}
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{8.0, 7.5}
		verifyArea(t, rectangle, 60.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{6.0}
		verifyArea(t, circle, 108.00)
	})
}
