package readfile

import (
	"testing"
)

func BenchmarkReadTxtFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readTxtFile(".fixtures/test.txt")
	}
}

func BenchmarkReadJsonLinesFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readJsonLinesFile(".fixtures/test.jsonl")
	}
}
