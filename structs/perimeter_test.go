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
	rectangle := Rectangle{8.0, 7.5}
	result := Area(rectangle)
	expected := 60.0
	if result != expected {
		t.Errorf("result '%.2f' expected '%.2f", result, expected)
	}
}
