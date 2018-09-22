package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(_ http.ResponseWriter, r *http.Request) {
		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("########################")
		fmt.Println("From: ", r.RemoteAddr)
		fmt.Println("------------------------")
		fmt.Println(string(requestDump))
		fmt.Println("########################")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
