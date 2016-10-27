package main

import (
	"math/rand"
	"net/http"

	"github.com/graphql-go/graphql"
	gqlhandler "github.com/graphql-go/handler"
)

func main() {

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"latestPost": &graphql.Field{
				Type:        graphql.String,
				Description: "The latest post",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello World!", nil
				}},
			"someRandomInt": &graphql.Field{
				Type:        graphql.Int,
				Description: "A random int.  No, really",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return rand.Intn(100), nil
				}},
		},
	})

	schemaConfig := graphql.SchemaConfig{Query: rootQuery}
	schema, _ := graphql.NewSchema(schemaConfig)

	// create a graphl-go HTTP handler with our previously defined schema
	// and we also set it to return pretty JSON output
	h := gqlhandler.New(&gqlhandler.Config{
		Schema: &schema,
		Pretty: true,
	})

	// static file server to serve Graphiql in-browser editor
	fs := http.FileServer(http.Dir("static"))

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)
	http.Handle("/", fs)

	// and serve!
	http.ListenAndServe(":8080", nil)

}
