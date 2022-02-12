package forms

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 12}
		waitedArea := float64(120)
		got := rectangle.Area()

		if waitedArea != got {
			t.Errorf("got %.2f, wanted %.2f", got, waitedArea)
		}
	})
	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		waitedArea := float64(math.Pi * 100)
		got := circle.Area()

		if waitedArea != got {
			t.Errorf("got %.2f, wanted %.2f", got, waitedArea)
		}
	})
}
