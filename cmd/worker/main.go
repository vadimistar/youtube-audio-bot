package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}

		w.Write(body)
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), mux))
}
