// package comment
package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"os"
	"sync"
)

var (
	cmd    = flag.String("cmd", "find", "find vector similarities")
	entity = flag.String("e", "product", "")
	val    = flag.String("val", "", "value to search for")
)

func main() {
	flag.Parse()

	ctx := context.Background()
	switch *cmd {
	case "find":
		data, err := load("attrs_" + *entity + "_emb.csv")
		if err != nil {
			panic(err)
		}

		search := NewAttributeSearch(data)
		fmt.Printf("Searching for %s\n", *val)
		searchResult, err := search.L2Distance(ctx, *val)
		if err != nil {
			panic(err)
		}
		for _, e := range searchResult[:10] {
			fmt.Printf("Attribute name: %s, score: %.9f\n", e.AttributeName, e.Score)
		}
	case "generate":
		oaiocli := NewOpenAI()
		fbb, err := os.ReadFile("attrs_" + *entity + ".txt")
		if err != nil {
			panic(err)
		}
		fb := bytes.Split(fbb, []byte("\n"))

		sem := make(chan struct{}, 10)
		mu := &sync.Mutex{}
		embedings := make([]Embedding, 0, len(fb))
		for i, f := range fb {
			name := string(f)
			if name == "" {
				continue
			}
			fmt.Printf("%d/%d: %s\n", i+1, len(fb), name)
			sem <- struct{}{}
			go func(name string) {
				defer func() { <-sem }()
				embeding, err := oaiocli.QueryEmbedding(ctx, name)
				if err != nil {
					panic(err)
				}
				mu.Lock()
				defer mu.Unlock()
				embedings = append(embedings, Embedding{
					AttributeName: name,
					Embedding:     embeding,
				})
			}(name)
		}
		err = save("attrs_"+*entity+"_emb.csv", embedings)
		if err != nil {
			panic(err)
		}
	}
}
