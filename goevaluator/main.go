package main

import (
	"log"

	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/pkg/errors"
)

func main() {
	exp, err := govaluate.NewEvaluableExpression("14 > 1")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create expr"))
	}

	result, err := exp.Eval(nil)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to evaluate"))
	}
	fmt.Println(result)
}
