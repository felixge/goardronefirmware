package main

import (
	"log"
	"time"
	"github.com/felixge/go-ardrone-firmware/navdata"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("start")

	navdataReader, err := navdata.NewReader()
	if err != nil {
		panic(err)
	}

	start := time.Now()
	for{
		data, err := navdataReader.ReadNavdata()
		if err != nil {
			panic(err)
		}

		log.Printf("navdata: %+v\n", data)

		if time.Since(start) > 30 * time.Second {
			break
		}
	}

	log.Println("quit")
}
