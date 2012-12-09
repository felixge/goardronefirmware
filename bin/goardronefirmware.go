package main

import "github.com/felixge/goardronefirmware"
import "github.com/felixge/goardronefirmware/http"
import "log"

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds)
	firmware, err := goardronefirmware.Start()
	if err != nil {
		panic(err)
	}
	http.Serve(firmware)
}
