package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Data struct {
	Name string `json:"name"`
	Hash string `json:"hash"`
}

func generateHash(name string) string {
	hash := sha256.Sum256([]byte(name))
	return hex.EncodeToString(hash[:])
}

func generateRandomName() string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func generateData() Data {
	name := generateRandomName()
	return Data{
		Name: name,
		Hash: generateHash(name),
	}
}

func main() {
	rand.NewSource(time.Now().UnixNano())

	pathDir := "./pkg/readfile/.fixtures/"
	txtFile, err := os.Create(filepath.Join(pathDir, "test.txt"))
	if err != nil {
		log.Fatalf("Failed to create TXT file: %s", err)
	}
	defer txtFile.Close()

	jsonlFile, err := os.Create(filepath.Join(pathDir, "test.jsonl"))
	if err != nil {
		log.Fatalf("Failed to create JSONL file: %s", err)
	}
	defer jsonlFile.Close()

	for i := 0; i < 140; i++ {
		data := generateData()

		txtLine := fmt.Sprintf("%s,%s\n", data.Name, data.Hash)
		if _, err := txtFile.WriteString(txtLine); err != nil {
			log.Fatalf("Failed to write to TXT file: %s", err)
		}

		jsonLine, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("Failed to marshal JSON: %s", err)
		}
		if _, err := jsonlFile.WriteString(string(jsonLine) + "\n"); err != nil {
			log.Fatalf("Failed to write to JSONL file: %s", err)
		}
	}

	fmt.Println("Arquivos gerados com sucesso!")
}
