package main

import (
	"bytes"
	"encoding/json"
	"os"
)

// Embedding is an embedding for an attribute.
type Embedding struct {
	AttributeName string
	Embedding     []float32
}

func convertEmbeddings(b []byte) (rez []float32) {
	err := json.Unmarshal(b, &rez)
	if err != nil {
		panic(err)
	}
	return rez
}

func load(filename string) ([]Embedding, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(buf, []byte("\n"))
	embeddings := make([]Embedding, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		values := bytes.Split(line, []byte("\t"))

		e := Embedding{
			AttributeName: string(values[0]),
			Embedding:     convertEmbeddings(values[1]),
		}
		embeddings = append(embeddings, e)
	}
	return embeddings, nil
}

func save(filename string, embeddings []Embedding) error {
	buf := bytes.NewBuffer(nil)
	for _, e := range embeddings {
		buf.WriteString(e.AttributeName)
		buf.WriteString("\t")
		b, err := json.Marshal(e.Embedding)
		if err != nil {
			return err
		}
		buf.Write(b)
		buf.WriteString("\n")
	}
	return os.WriteFile(filename, buf.Bytes(), 0600)
}
