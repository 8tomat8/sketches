package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type Error int

func (e Error) Error() string {
	return strconv.Itoa(int(e))
}

const (
	ErrStopIt Error = iota + 100
)

type Book struct {
	ID     string   `json:"id, omitempty"`
	Title  string   `json:"title, omitempty"`
	Ganres []string `json:"ganres, omitempty"`
	Pages  int      `json:"pages, omitempty"`
	Price  float64  `json:"price, omitempty"`
}

func main() {
	data := bytes.NewReader([]byte(`{
        "id": "C97376B9-6C2E-41E5-9DBE-2E82C0EF114B",
        "title": "Book title2",
        "ganres": [
            "adventure"
        ],
        "pages": 234,
        "price": 25.43
    }`))

	bo := &Book{}

	err := json.NewDecoder(data).Decode(bo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", bo)

	inputData := bytes.NewReader([]byte(`{
        "title": "Changed TITLE!!!!"
    }`))

	err = json.NewDecoder(inputData).Decode(bo)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n", bo)
}
