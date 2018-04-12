package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rodrigobotti/go-lessons/goroutines/model"
)

var (
	orchestrator sync.WaitGroup
)

func main() {
	measured("sequential", func() {
		csvToJSON("goroutines/saopaulo")
		csvToJSON("goroutines/riodejaneiro")
	})
	measured("concurrent", func() {
		orchestrator.Add(2)
		go orchestrated(func() { csvToJSON("goroutines/saopaulo") })
		go orchestrated(func() { csvToJSON("goroutines/riodejaneiro") })
		orchestrator.Wait()
	})
}

func orchestrated(f func()) {
	f()
	orchestrator.Done()
}

func measured(name string, f func()) {
	start := time.Now()
	f()
	fmt.Println(name, "took", time.Since(start))
}

func csvToJSON(fileName string) {
	fmt.Println(time.Now(), "-- Started Processing ", fileName, " --")
	file, err := os.Open(fileName + ".csv")
	if err != nil {
		fmt.Println("[main] Error when opening file", err)
		return
	}
	// printLines(file)
	// printCsv(file)
	jsonFile, err := os.Create(fileName + ".json")
	defer jsonFile.Close() // defers execution until after surrounding function returns; arguments are evaluated immediately
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
		fmt.Println(string(json))
		writer.WriteString("  " + string(json))
		if (ic + 1) < len(line) {
			writer.WriteString(",")
		}
		writer.WriteString("\n")
	})
	writer.WriteString("]")
	writer.Flush()
	fmt.Println(time.Now(), "-- Finished Processing ", fileName, " --")
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
