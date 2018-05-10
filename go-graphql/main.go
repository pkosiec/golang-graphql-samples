package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/pkosiec/golang-graphql-samples/go-graphql/schema"
)

func main() {
	schema, err := schema.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	handler := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 3000)
	log.Printf("Listening on %s", addr)
	http.ListenAndServe(addr, handler)
}
