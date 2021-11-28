package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goombaio/namegenerator"
)

func main() {
	http.HandleFunc("/", help)
	http.ListenAndServe("localhost:8080", nil)
}

func help(writer http.ResponseWriter, request *http.Request) {
	seed := time.Now().UTC().Unix()
	rName := namegenerator.NewNameGenerator(seed)
	fmt.Fprintf(writer, "Hello %s", rName.Generate())
}
