package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"./calculator"
	"./shapes"
)

func main() {
	logFile, err := os.OpenFile("errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Ошибка при открытии лог-файла:", err)
		return
	}
	defer logFile.Close()

	logger := log.New(logFile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

	if len(os.Args) < 2 {
		logger.Println("Укажите файл с фигурами")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		logger.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	var figures []interface{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			logger.Println("Некорректная строка:", line)
			continue
		}

		shapeType := parts[0]
		switch shapeType {
		case "Circle":
			if len(parts) != 2 {
				logger.Println("Некорректные параметры для круга:", line)
				continue
			}
			radius, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				logger.Println("Ошибка при парсинге радиуса:", err)
				continue
			}
			figures = append(figures, shapes.Circle{Radius: radius})

		case "Rectangle":
			if len(parts) != 3 {
				logger.Println("Некорректные параметры для прямоугольника:", line)
				continue
			}
			width, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				logger.Println("Ошибка при парсинге ширины:", err)
				continue
			}
			height, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				logger.Println("Ошибка при парсинге высоты:", err)
				continue
			}
			figures = append(figures, shapes.Rectangle{Width: width, Height: height})

		default:
			logger.Println("Неизвестная фигура:", shapeType)
		}
	}

	totalArea, err := calculator.TotalArea(logger, figures...)
	if err != nil {
		logger.Println("Ошибка при вычислении площади:", err)
		return
	}

	fmt.Printf("Суммарная площадь: %.2f\n", totalArea)
}
