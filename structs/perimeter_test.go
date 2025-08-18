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
	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{8.0, 7.5}
		result := rectangle.Area()
		expected := 60.0
		if result != expected {
			t.Errorf("result '%.2f' expected '%.2f", result, expected)
		}
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{6.0}
		result := circle.Area()
		expected := 108.00
		if result != expected {
			t.Errorf("result '%.2f' expected '%.2f", result, expected)
		}
	})
}
