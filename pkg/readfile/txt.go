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
			lines[i] = "valorescollido"
			break
		}
	}

	if _, err := file.Seek(0, 0); err != nil {
		fmt.Errorf("failed to seek to beginning of file: %s", err.Error())
		return
	}

	if err := file.Truncate(0); err != nil {
		fmt.Errorf("failed to truncate file: %s", err)
		return
	}

	for _, line := range lines {
		if _, err := fmt.Fprintln(file, line); err != nil {
			fmt.Errorf("failed to write line to file: %s", err)
			return
		}
	}

	updatedData, err := json.Marshal(lines)
	if err != nil {
		return
	}

	if err := os.WriteFile(filePath, updatedData, 0644); err != nil {
		return
	}

}
