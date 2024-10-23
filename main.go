package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Print("Введите имя входного файла: ")
	var inputFile string
	fmt.Scanln(&inputFile)
	file, err := os.Open(inputFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл не существует.")
		} else {
			fmt.Println("Ошибка при открытии файла:", err)
		}
		return
	}
	defer file.Close()
	lineCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineCount[line]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}
	uniqueLines := make([]string, 0)
	for line, count := range lineCount {
		if count == 1 {
			uniqueLines = append(uniqueLines, strings.ToUpper(line))
		}
	}
	sort.Strings(uniqueLines)
	fmt.Print("Введите имя выходного файла: ")
	var outputFileName string
	fmt.Scanln(&outputFileName)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Ошибка при создании выходного файла:", err)
		return
	}
	defer outputFile.Close()
	for _, line := range uniqueLines {
		byteCount := len([]byte(line))
		_, err := fmt.Fprintf(outputFile, "%s - %d байт\n", line, byteCount)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}
	fmt.Println("Результат записан в файл:", outputFileName)
}
