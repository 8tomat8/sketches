package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/couchbase/go-couchbase"
)

var (
	pool   = flag.String("p", "default", "pool")
	bucket = flag.String("b", "bucket", "bucket")
	addr   = flag.String("addr", "http://couchbase-couchbase-cluster:8091/", "addr")
	user   = flag.String("usr", "user", "user")
	pass   = flag.String("pwd", "password", "pass")
)

func main() {
	flag.Parse()
	fmt.Println(*pool)
	fmt.Println(*bucket)
	fmt.Println(*addr)
	fmt.Println(*user)
	fmt.Println(*pass)

	cli, err := couchbase.ConnectWithAuthCreds(*addr, *user, *pass)
	if err != nil {
		log.Fatalf("Error connecting:  %v", err)
	}

	p, err := cli.GetPool(*pool)
	if err != nil {
		log.Fatal("get pool", err)
	}

	bucket, err := p.GetBucket(*bucket)
	if err != nil {
		log.Fatal("get bucket", err)
	}

	err = bucket.Set("someKey", 0, []string{"an", "example", "list"})
	if err != nil {
		log.Fatalf("failed to set: %v", err)
	}

	res := []string{}
	err = bucket.Get("someKey", &res)
	if err != nil {
		log.Fatalf("failed to get: %v", err)
	}

	fmt.Println(res)
}
