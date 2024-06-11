package readfile

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Data struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

func readJsonLinesFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count += 1
		line := scanner.Text()
		var data Data
		if err := json.Unmarshal([]byte(line), &data); err != nil {
			log.Printf("Failed to unmarshal JSON: %s", err)
			return
		}
		fmt.Printf("\rJSONL ### Progress: %d - Name: %s, Hash: %s%%", count, data.Name, data.Hash)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
}
