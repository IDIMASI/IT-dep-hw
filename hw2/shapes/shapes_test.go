package shapes

import (
	"testing"
)

func TestCircleArea(t *testing.T) {
	circle := Circle{Radius: 5}
	area, err := circle.Area()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedArea := 78.53981633974483 // π * 5^2
	if area != expectedArea {
		t.Errorf("Expected area %v, got %v", expectedArea, area)
	}
}

func TestCircleAreaInvalidRadius(t *testing.T) {
	circle := Circle{Radius: -1}
	area, err := circle.Area()
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	if area != 0 {
		t.Errorf("Expected area 0, got %v", area)
	}
}

func TestRectangleArea(t *testing.T) {
	rectangle := Rectangle{Width: 4, Height: 5}
	area, err := rectangle.Area()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedArea := 20.0 // 4 * 5
	if area != expectedArea {
		t.Errorf("Expected area %v, got %v", expectedArea, area)
	}
}

func TestRectangleAreaInvalidDimensions(t *testing.T) {
	rectangle := Rectangle{Width: 0, Height: 5}
	area, err := rectangle.Area()
	if err == nil {
		t.Errorf("Expected an error, got none")
	}
	if area != 0 {
		t.Errorf("Expected area 0, got %v", area)
	}
}
