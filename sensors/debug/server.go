package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/felixge/go-ardrone-firmware/navdata"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", time.Now())
}

func readNavdata() {
	reader, err := navdata.NewReader()
	if err != nil {
		panic(err)
	}

	for{
		buf := make([]byte, 60)
		reader.Read(buf)
	}
}

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)

	go readNavdata()

	log.Println("Starting navdata debug server ...")
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}

	log.Println("Terminating")
}
