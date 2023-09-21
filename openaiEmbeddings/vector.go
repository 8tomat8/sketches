package main

import (
	"context"
)

// AttributeSearch is a search for an attribute in a vector space.
type AttributeSearch struct {
	openAIcli  *OpenAI
	embeddings []Embedding
}

// SearchResult is a result of a search.
type SearchResult struct {
	AttributeName string
	Score         float64
}

// NewAttributeSearch returns a new AttributeSearch.
func NewAttributeSearch(embeddings []Embedding) *AttributeSearch {
	return &AttributeSearch{openAIcli: NewOpenAI(), embeddings: embeddings}
}

// L2Distance returns the L2 distance between the text and the embeddings.
func (a *AttributeSearch) L2Distance(ctx context.Context, text string) ([]SearchResult, error) {
	emb, err := a.openAIcli.QueryEmbedding(ctx, text)
	if err != nil {
		return nil, err
	}
	return l2distance(a.embeddings, emb)
}
