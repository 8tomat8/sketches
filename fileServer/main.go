package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func main() {
	curPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to get current file path"))
	}

	// Base
	port := flag.String("port", "8100", "port to serve on")
	directory := flag.String("dir", curPath, "the directory of static file to host")
	// TLS
	certPath := flag.String("cert", "", "path to the tls certificate")
	keyPath := flag.String("key", "", "path to the tls key")

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	if *certPath != "" && *keyPath != "" {
		log.Printf("Serving %s on HTTPS port: %s\n", *directory, *port)
		log.Fatal(http.ListenAndServeTLS(":"+*port, *certPath, *keyPath, nil))
	} else {
		log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
		log.Fatal(http.ListenAndServe(":"+*port, nil))
	}
}
