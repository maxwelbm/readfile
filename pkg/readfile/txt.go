package readfile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func processLine(line string) {
}

func readTxtFile(filePath string) {
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

		lineSplit := strings.SplitN(line, ",", 2)
		fmt.Printf("\rTXT   ### Progress: %d - Name: %s, Hash: %s%%", count, lineSplit[0], lineSplit[1])
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
}
