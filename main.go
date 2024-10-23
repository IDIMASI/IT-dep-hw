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

func readFile(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lineCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineCount[line]++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lineCount, nil
}

func processLines(lineCount map[string]int) []string {
	uniqueLines := make([]string, 0)
	for line, count := range lineCount {
		if count == 1 {
			uniqueLines = append(uniqueLines, strings.ToUpper(line))
		}
	}
	sort.Strings(uniqueLines)
	return uniqueLines
}

func writeFile(fileName string, uniqueLines []string) error {
	outputFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	for _, line := range uniqueLines {
		byteCount := len([]byte(line))
		_, err := fmt.Fprintf(outputFile, "%s - %d байт\n", line, byteCount)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	fmt.Print("Введите имя входного файла: ")
	var inputFileName string
	fmt.Scanln(&inputFileName)

	lineCount, err := readFile(inputFileName)
	if err != nil {
		if os.IsNotExist(err) {
			logError(fmt.Errorf("Файл %s не существует.", inputFileName))
		} else {
			logError(fmt.Errorf("Ошибка при открытии файла %s: %v", inputFileName, err))
		}
		return
	}

	uniqueLines := processLines(lineCount)

	fmt.Print("Введите имя выходного файла: ")
	var outputFileName string
	fmt.Scanln(&outputFileName)

	if err := writeFile(outputFileName, uniqueLines); err != nil {
		logError(fmt.Errorf("Ошибка при записи в файл %s: %v", outputFileName, err))
		return
	}

	fmt.Println("Результат записан в файл:", outputFileName)
}
