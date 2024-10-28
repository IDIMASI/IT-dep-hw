package calculator

import (
	"errors"
	"log"

	"../shapes"
)

func TotalArea(logger *log.Logger, figures ...interface{}) (float64, error) {
	total := 0.0
	for _, figure := range figures {
		switch f := figure.(type) {
		case shapes.Circle:
			area, err := f.Area()
			if err != nil {
				logger.Println("Ошибка при вычислении площади круга:", err)
				return 0, err
			}
			total += area
		case shapes.Rectangle:
			area, err := f.Area()
			if err != nil {
				logger.Println("Ошибка при вычислении площади прямоугольника:", err)
				return 0, err
			}
			total += area
		default:
			err := errors.New("недопустимый тип фигуры")
			logger.Println(err)
			return 0, err
		}
	}
	return total, nil
}
