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
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello World!", nil
				}},
			"someRandomInt": &graphql.Field{
				Type: graphql.Int,
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

	// serve a GraphQL endpoint at `/graphql`
	http.Handle("/graphql", h)

	// and serve!
	http.ListenAndServe(":8080", nil)

}
