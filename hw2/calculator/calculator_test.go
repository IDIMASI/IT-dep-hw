package calculator

import (
	"log"
	"testing"

	"hw2/shapes"
)

func TestTotalArea(t *testing.T) {
	circle := shapes.Circle{Radius: 5}
	rectangle := shapes.Rectangle{Width: 4, Height: 5}

	totalArea, err := TotalArea(log.Default(), circle, rectangle)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedArea := 78.53981633974483 + 20.0 // π * 5^2 + 4 * 5
	if totalArea != expectedArea {
		t.Errorf("Expected total area %v, got %v", expectedArea, totalArea)
	}
}

func TestTotalAreaInvalidFigure(t *testing.T) {
	invalidFigure := struct{}{} // Неизвестный тип фигуры

	_, err := TotalArea(log.Default(), invalidFigure)
	if err == nil {
		t.Errorf("Expected an error for invalid figure, got none")
	}
}
