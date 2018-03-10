package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/rodrigobotti/go-lessons/files/model"
)

func main() {
	file, err := os.Open("files/cities.csv")
	if err != nil {
		fmt.Println("[main] Error when opening file", err)
		return
	}
	// printLines(file)
	// printCsv(file)
	jsonFile, err := os.Create("files/cities.json")
	if err != nil {
		fmt.Println("[main] failed to create file", err)
		return
	}
	writer := bufio.NewWriter(jsonFile)
	writer.WriteString("[\n")
	iterateCsv(file, func(il int, ic int, line []string, item string) {
		data := strings.Split(item, "/")
		city := model.City{Name: data[0], State: data[1]}
		json, err := json.Marshal(city)
		if err != nil {
			fmt.Println("[main] error serializing city json for item", item, err)
		}
		writer.WriteString("  " + string(json))
		if (ic + 1) < len(line) {
			writer.WriteString(",")
		}
		writer.WriteString("\n")
	})
	writer.WriteString("]")
	writer.Flush()
	jsonFile.Close()
}

func printLines(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("line is:", line)
	}
	file.Close()
}

func iterateCsv(file *os.File, iterator func(int, int, []string, string)) {
	csvReader := csv.NewReader(file)
	content, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("[main] error reading csv", err)
		return
	}
	for indexLine, line := range content {
		for indexCol, item := range line {
			iterator(indexLine, indexCol, line, item)
		}
	}
	file.Close()
}

func printCsv(file *os.File) {
	iterateCsv(file, func(indexLine int, indexCol int, line []string, item string) {
		fmt.Printf("item [%d, %d] is: %s\n", indexLine, indexCol, item)
	})
}
