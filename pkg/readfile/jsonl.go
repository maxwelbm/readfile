package readfile

import (
	"encoding/json"
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

	dataMap := make(map[string]Data)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&dataMap); err != nil {
		return
	}

	if _, ok := dataMap["KNwqaWGXHd"]; !ok {
		return
	}

	dataMap["KNwqaWGXHd"] = Data{Name: "asdfadsfdsf", Hash: "asdfadsfdsfadsf"}

	if _, err := file.Seek(0, 0); err != nil {
		return
	}

	if err := file.Truncate(0); err != nil {
		return
	}

	_, err = json.Marshal(dataMap)
	if err != nil {
		return
	}
}
