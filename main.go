package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func logError(err error) {
	logFile, logErr := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if logErr != nil {
		fmt.Println("Ошибка при открытии лог-файла:", logErr)
		return
	}
	defer logFile.Close()

	writer := bufio.NewWriter(logFile)
	writer.WriteString(fmt.Sprintf("%s\n", err.Error()))
	writer.Flush()
}

func readFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("can't os.Open on readFile: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("can't scan lines on readFile: %w", err)
	}

	return lines, nil
}

func processLines(lines []string) ([]string, error) {
	lineCount := make(map[string]int)
	for _, line := range lines {
		lineCount[line]++
	}

	uniqueLines := make([]string, 0)
	for line, count := range lineCount {
		if count == 1 {
			uniqueLines = append(uniqueLines, strings.ToUpper(line))
		}
	}
	sort.Strings(uniqueLines)
	return uniqueLines, nil
}

func writeFile(fileName string, uniqueLines []string) error {
	outputFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("can't os.Create on writeFile: %w", err)
	}
	defer outputFile.Close()

	var outputBuilder strings.Builder

	for _, line := range uniqueLines {
		byteCount := len([]byte(line))
		outputBuilder.WriteString(fmt.Sprintf("%s - %d байт\n", line, byteCount))
	}

	if _, err := outputFile.WriteString(outputBuilder.String()); err != nil {
		return fmt.Errorf("can't write to file on writeFile: %w", err)
	}

	return nil
}

func main() {
	fmt.Print("Введите имя входного файла: ")
	var inputFileName string
	fmt.Scanln(&inputFileName)

	lines, err := readFile(inputFileName)
	if err != nil {
		if os.IsNotExist(err) {
			logError(fmt.Errorf("Файл %s не существует: %w", inputFileName, err))
		} else {
			logError(fmt.Errorf("Ошибка при открытии файла %s: %w", inputFileName, err))
		}
		return
	}

	uniqueLines, err := processLines(lines)
	if err != nil {
		logError(fmt.Errorf("Ошибка при обработке строк: %w", err))
		return
	}

	fmt.Print("Введите имя выходного файла: ")
	var outputFileName string
	fmt.Scanln(&outputFileName)

	if err := writeFile(outputFileName, uniqueLines); err != nil {
		logError(fmt.Errorf("Ошибка при записи в файл %s: %w", outputFileName, err))
		return
	}

	fmt.Println("Результат записан в файл:", outputFileName)
}
