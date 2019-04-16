package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	oauth_user_scope_validator "github.com/graphql-services/oauth-user-scope-validator"
)

const defaultPort = "80"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(oauth_user_scope_validator.NewExecutableSchema(oauth_user_scope_validator.Config{Resolvers: &oauth_user_scope_validator.Resolver{}})))

	http.HandleFunc("/healthcheck", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("OK"))
		res.WriteHeader(200)
	})

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
