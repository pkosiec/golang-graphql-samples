package main

import (
	"log"
	"net/http"

	"github.com/vektah/gqlgen/example/selection"
	"github.com/vektah/gqlgen/handler"
)

func main() {
	http.Handle("/", handler.Playground("Selection Demo", "/query"))
	http.Handle("/query", handler.GraphQL(selection.MakeExecutableSchema(&selection.SelectionResolver{})))
	log.Fatal(http.ListenAndServe(":8086", nil))
}
