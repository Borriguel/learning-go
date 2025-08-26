package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{6.0, 5.0}
	got := Perimeter(rectangle)
	want := 22.0
	if got != want {
		t.Errorf("got '%.2f' want '%.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got '%.2f' want '%.2f", got, want)
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

	t.Run("triangle", func(t *testing.T) {
		triangle := Triangle{12.0, 6.0}
		verifyArea(t, triangle, 36.0)
	})
}
