package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/felixge/goardronefirmware"
	"log"
	"net/http"
)

type server struct {
	firmware *goardronefirmware.Firmware
}

func Serve(f *goardronefirmware.Firmware) {
	server := &server{
		firmware: f,
	}
	server.start()
}

func (s *server) start() {
	http.Handle("/", s)
	http.ListenAndServe(":80", nil)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("http: %s %s\n", r.Method, r.URL)

	var data interface{}
	var err error
	switch r.URL.Path {
	//case "/sensors":
		//data, err = s.getSensors()
	//case "/burnout":
		//data, err = s.firmware.Burnout()
	//case "/disco":
		//data, err = s.firmware.Disco()
	case "/favicon.ico":
		return
	default:
		err = errors.New("404")
	}

	if err != nil {
		log.Printf("http: error %+v\n", err)
		data = struct{Error string}{err.Error()}
	}

	response, err := json.Marshal(data)
	if err != nil {
		fmt.Fprintf(w, "Error: %+v\n", err)
		return
	}

	w.Write(response)
}
