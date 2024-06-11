package readfile

import (
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

	dataMap := make(map[string]Data)
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&dataMap); err != nil {
		fmt.Errorf("failed to decode JSON: %s", err)
		return
	}

	// Verifica se o item a ser substituído existe no mapa
	if _, ok := dataMap["KNwqaWGXHd"]; !ok {
		return
	}

	// Substitui o item no mapa
	dataMap["KNwqaWGXHd"] = Data{Name: "novo nome", Hash: "Novo hash"}

	// Move o cursor para o início do arquivo para reescrever o mapa no arquivo
	if _, err := file.Seek(0, 0); err != nil {
		return
	}

	// Trunca o arquivo para limpar o conteúdo antigo
	if err := file.Truncate(0); err != nil {
		return
	}

	// Codifica o mapa atualizado de volta no arquivo JSONL
	updatedData, err := json.Marshal(dataMap)
	if err != nil {
		return
	}

	if err := os.WriteFile(filePath, updatedData, 0644); err != nil {
		return
	}
}
