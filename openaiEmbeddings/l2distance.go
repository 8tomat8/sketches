package main

import (
	"math"
	"sort"
)

func l2distance(embeddings []Embedding, emb []float32) ([]SearchResult, error) {
	results := make([]SearchResult, len(embeddings))
	for i, row := range embeddings {
		distance := float64(0)
		for j, elem := range row.Embedding {
			difference := float64(elem - emb[j])
			distance += difference * difference
		}
		results[i].Score = math.Sqrt(distance)
		results[i].AttributeName = row.AttributeName
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Score < results[j].Score
	})
	return results, nil
}
