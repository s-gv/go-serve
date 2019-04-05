package main

import (
	"net/http"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
    port := "8080"
    if len(os.Args) == 2 {
        port = os.Args[1]
    }
	println("Serving at http://0.0.0.0:" + port)
	panic(http.ListenAndServe("0.0.0.0:" + port, http.FileServer(http.Dir(wd))))
}
