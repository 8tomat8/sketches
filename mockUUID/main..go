package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
)

func main() {
	// V4
	r := rand.New(rand.NewSource(1))
	uuid.SetRand(r)
	for i := 0; i < 3; i++ {
		fmt.Println(uuid.Must(uuid.NewRandom()))
	}
}
