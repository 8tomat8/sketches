package main

import (
	"fmt"
	"log"

	"github.com/neurodrone/crdt"
	"github.com/pkg/errors"
)

func main() {
	obj := `{"a":1}`
	lwwset1, err := crdt.NewLWWSet()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init LWW set"))
	}
	lwwset2, err := crdt.NewLWWSet()
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to init LWW set"))
	}

	// Here, we remove the object first before we add it in. For a
	// 2P-set the object would be deemed absent from the set. But for
	// a LWW-set the object should be present because `.Add()` follows
	// `.Remove()`.
	lwwset1.Remove(obj)
	lwwset1.Add(obj)
	lwwset2.Add(`{"a":2}`)
	lwwset2.Remove(`{"a":1}`)

	// This should print 'true' because of the above.
	data1, err := lwwset1.MarshalJSON()
	fmt.Println(string(data1))

	lwwset2.Merge(lwwset1)
	data2, err := lwwset2.MarshalJSON()
	fmt.Println(string(data2))
}
