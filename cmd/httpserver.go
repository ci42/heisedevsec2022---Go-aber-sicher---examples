package main

import (
	"examples/bar"
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Hello Heise devSec() 2022.\n")
	bar.ServerTemplate(w)
	bar.MaybePanic()
}

func RecoveryHandler(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Fprintf(w, "recovered form %s", r)
			}
		}()
		handlerFunc(w, req)
	}
}

func main() {
	http.HandleFunc("/hello", RecoveryHandler(HelloServer))
	err := http.ListenAndServeTLS(":8443", "localhost+2.pem", "localhost+2-key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
