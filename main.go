package main

import (
	"net/http"

	"github.com/danesparza/graphql_test/datastores"
	gqlhandler "github.com/graphql-go/handler"
)

func main() {

	// create a graphql-go HTTP handler with our previously defined schema
	// and also set it to return pretty JSON output
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &datastores.StarWarsSchema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	//	serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)
	//	serve the GraphiQL UI at the root:
	http.Handle("/", fs)

	// and serve!
	http.ListenAndServe(":8080", nil)

}
