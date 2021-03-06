package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Rect{10.0, 10.0}.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("want %.2f but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {

	cases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rect", shape: Rect{Width: 10, Height: 10}, want: 100},
		{name: "Circle", shape: Circle{Radius: 10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, want: 36},
	}

	for _, tt := range cases {
		tc := tt
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			if got != tc.want {
				t.Errorf("want %g but got %g", tc.want, got)
			}
		})
	}
}
