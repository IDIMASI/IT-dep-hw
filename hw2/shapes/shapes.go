package shapes

import (
	"errors"
	"math"
)

type Circle struct {
	Radius float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (c Circle) Area() (float64, error) {
	var ErrInvalidRadius = errors.New("некорректный радиус")
	if c.Radius <= 0 {
		return 0, ErrInvalidRadius
	}
	return math.Pi * c.Radius * c.Radius, nil
}

func (c Rectangle) Area() (float64, error) {
	var ErrInvalidDimensions = errors.New("Некорректный величины")
	if c.Width <= 0 || c.Height <= 0 {
		return 0, ErrInvalidDimensions
	}
	return math.Pi * c.Height * c.Width, nil
}
