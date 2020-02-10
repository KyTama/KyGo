package config

import (
	"log"
	"net/http"
)

func init() {

}

func Listen(addr string, handler http.Handler) {
	var err error
	err = http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatal(err)
	}
}