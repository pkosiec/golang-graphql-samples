package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkosiec/golang-graphql-samples/gqlgen/gqlschema"

	"github.com/pkosiec/golang-graphql-samples/gqlgen/rootresolver"
	"github.com/vektah/gqlgen/handler"
)

func main() {
	resolver := rootresolver.New()
	schema := gqlschema.MakeExecutableSchema(resolver)

	mux := http.NewServeMux()
	mux.Handle("/", handler.Playground("Dataloader", "/graphql"))
	mux.Handle("/graphql", handler.GraphQL(schema))

	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 3000)
	log.Printf("Listening on %s", addr)

	http.ListenAndServe(addr, mux)
}
