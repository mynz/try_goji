package main

import (
	"fmt"
	"goji.io"
	"goji.io/pat"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	fmt.Println("try-goji")

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		panic(err)
	}
}
