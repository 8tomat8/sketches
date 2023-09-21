package main

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

// OpenAI is a wrapper around the OpenAI API.
type OpenAI struct {
	client *openai.Client
}

// NewOpenAI returns a new OpenAI.
func NewOpenAI() *OpenAI {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	return &OpenAI{client: client}
}

// QueryEmbedding returns the embedding for the text.
func (o *OpenAI) QueryEmbedding(ctx context.Context, text string) ([]float32, error) {
	emb, err := o.client.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Input: text,
		Model: openai.AdaEmbeddingV2,
		User:  "emb-test",
	})
	if err != nil {
		return nil, err
	}
	return emb.Data[0].Embedding, nil
}
