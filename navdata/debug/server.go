package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", time.Now())
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	log.Println("Starting navdata debug server ...")

	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}

	log.Println("Terminating")
}
