package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"embed"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/pkg/errors"
)

//go:embed urls.txt urls_nf.txt
var urlsFile embed.FS

const workers = 150

func main() {
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	cli := &http.Client{Transport: customTransport}

	// urls, err := urlsFile.Open("urls.txt")
	urls, err := urlsFile.Open("urls_nf.txt")
	if err != nil {
		log.Fatal(err)
	}

	tasks := make(chan string, workers)
	wg := &sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(wg *sync.WaitGroup) {
			x := 0
			for task := range tasks {
				resp, err := cli.Get(task)
				if err != nil {
					log.Println(errors.Wrapf(err, "get %s", task))
					continue
				}

				if resp.StatusCode >= http.StatusBadRequest || resp.StatusCode < http.StatusOK {
					log.Printf("status code %d for %s\n", resp.StatusCode, task)
					continue
				}

				var buf bytes.Buffer
				_, err = io.Copy(&buf, resp.Body)
				if err != nil {
					log.Println(errors.Wrapf(err, "copy reader for %s", task))
					continue
				}
				// Just a hack to avoid compiler optimization
				x += buf.Len()
			}
			fmt.Println(x)
		}(wg)
	}

	scanner := bufio.NewScanner(urls)
	for i := 0; scanner.Scan(); i++ {
		if i%10 == 0 {
			log.Printf("Passed %d\n", i)
		}
		tasks <- scanner.Text()
	}

	close(tasks)
	wg.Wait()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
