package main

import (
	"flag"
	"net/http"
)

func main() {
	flag.Parse()

	s := &http.Server{
		Addr: ":8080",
	}
	s.ListenAndServe()

}
