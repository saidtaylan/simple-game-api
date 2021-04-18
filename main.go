package main

import "net/http"

func main() {
	connectToDB()
	server := &http.Server{
		Addr:    ":80",
		Handler: multiplexer(),
	}
	server.ListenAndServe()
}
