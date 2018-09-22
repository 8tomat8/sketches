package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"gopkg.in/h2non/filetype.v1"
)

func main() {
	links := [][2]string{
		{"doc", "http://techslides.com/demos/samples/sample.doc"},
		{"docx", "http://techslides.com/demos/samples/sample.docx"},
		{"pdf", "http://techslides.com/demos/samples/sample.pdf"},

		{"zip", "http://techslides.com/demos/samples/sample.zip"},
		{"rar", "http://techslides.com/demos/samples/sample.rar"},
		{"tar", "http://techslides.com/demos/samples/sample.tar"},

		{"mp3", "http://techslides.com/demos/samples/sample.mp3"},
		{"wav", "http://techslides.com/demos/samples/sample.wav"},

		//{"txt", "http://techslides.com/demos/samples/sample.txt"},
		{"xml", "http://techslides.com/demos/samples/sample.xml"},
		{"json", "http://techslides.com/demos/samples/sample.json"},

		{"avi", "http://techslides.com/demos/samples/sample.avi"},
		{"mov", "http://techslides.com/demos/samples/sample.mov"},
		{"mp4", "http://techslides.com/demos/samples/sample.mp4"},
		{"mpg", "http://techslides.com/demos/samples/sample.mpg"},
		{"wmv", "http://techslides.com/demos/samples/sample.wmv"},
		{"flv", "http://techslides.com/demos/samples/sample.flv"},
		{"swf", "http://techslides.com/demos/samples/sample.swf"},
		{"webm", "http://techslides.com/demos/samples/sample.webm"},
		{"mkv", "http://techslides.com/demos/samples/sample.mkv"},
		{"real-mp4", "https://dmbqekwh0sti7.cloudfront.net/qwtv%5E8528226876203780818.mp4"},
		//{"real-mkv", "http://127.0.0.1:8100/Ready.Player.One.2018.WEB-DL.1080p.DUB.Line.IVA%28RUS.UKR.ENG%29.ExKinoRay.mkv"},

		{"png-noext", "https://dummyimage.com/600x400/000/00ffd5"},
		{"png", "https://dummyimage.com/600x400/000/00ffd5.png"},
		{"jpg", "https://dummyimage.com/600x400/000/00ffd5.jpg"},
	}

	for _, link := range links {
		resp, err := http.Get(link[1])
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to execute http request"))
		}

		f := resp.Body
		buf := make([]byte, 256)
		n, err := f.Read(buf)
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to read data"))
		}

		t, err := filetype.Match(buf[:n])
		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to determine file type"))
		}
		fmt.Println(link[0], t.Extension, t.MIME.Value, resp.ContentLength)

		if err = f.Close(); err != nil {
			log.Fatal(errors.Wrap(err, "failed to close data reader"))
		}
	}
}
