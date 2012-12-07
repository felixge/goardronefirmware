package main

import (
	"log"
	//"time"
	"github.com/felixge/go-ardrone-firmware/navdata"
	"io"
	"os"
)

func main() {
	log.SetFlags(log.Lmicroseconds)
	log.Println("start")

	navdataReader, err := navdata.NewReader()
	if err != nil {
		panic(err)
	}

	fixture, err := os.OpenFile(
		"/data/video/navdata.bin",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0,
	)
	if err != nil {
		panic(err)
	}

	n, err := io.Copy(fixture, navdataReader)
	if err != nil {
		panic(err)
	}

	log.Printf("Wrote %d bytes\n", n)

	//start := time.Now()
	//for{
	//data, err := navdataReader.ReadNavdata()
	//if err != nil {
	//panic(err)
	//}

	//log.Printf("navdata: %+v\n", data)

	//if time.Since(start) > 30 * time.Second {
	//break
	//}
	//}

	log.Println("quit")
}
