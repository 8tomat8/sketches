package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/graymeta/stow"
	"github.com/graymeta/stow/s3"
)

func main() {
	kind := "s3"
	config := stow.ConfigMap{
		s3.ConfigAccessKeyID: os.Getenv("aws_access_key_id"),
		s3.ConfigSecretKey:   os.Getenv("aws_secret_key"),
		s3.ConfigRegion:      os.Getenv("aws_region"),
	}
	location, err := stow.Dial(kind, config)
	if err != nil {
		log.Fatal(err)
	}
	defer location.Close()

	ctr, err := location.Container(os.Getenv("aws_s3_bucket_name"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = ctr.Put(os.Getenv("aws_s3_item_name")+strconv.Itoa(int(time.Now().UTC().Unix())), bytes.NewReader([]byte(os.Getenv("aws_s3_item_content"))), 42, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success upload")
}
