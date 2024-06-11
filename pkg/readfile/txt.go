package readfile

import (
	"bufio"
	"encoding/json"
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
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for i, line := range lines {
		lineSplit := strings.SplitN(line, ",", 2)
		if lineSplit[0] == "KNwqaWGXHd" {
			lines[i] = "asdfadsadfs"
			break
		}
	}

	if _, err := file.Seek(0, 0); err != nil {
		return
	}

	if err := file.Truncate(0); err != nil {
		return
	}

	for _, line := range lines {
		if _, err := fmt.Fprintln(file, line); err != nil {
			return
		}
	}

	_, err = json.Marshal(lines)
	if err != nil {
		return
	}
}
