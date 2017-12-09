package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/hijack", func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r, nil, 2<<10, 2<<10)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		go func() {
			defer conn.Close()
			for {
				msgType, data, err := conn.ReadMessage()
				if err != nil {
					log.Printf("error reading string: %v", err)
					return
				}
				conn.WriteMessage(1, []byte(fmt.Sprintf("You said: %q. With msgType %d\n", string(data), msgType)))
			}
			fmt.Println("Closed: ", conn)
		}()
	})
	http.ListenAndServe("localhost:8081", http.DefaultServeMux)
}
