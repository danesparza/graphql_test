# graphql_test
Simple Golang based GraphQL endpoint test

To test, POST to `http://localhost:8080/graphql` with `Content-Type: application/graphql` and the [GraphQL](http://graphql.org/learn/) query:
```graphql
query Root{ latestPost }
```

If the test is successful, you should get the following response back:
```json
{
	"data": {
		"latestPost": "Hello World!"
	}
}
```

You can also navigate to http://localhost:8080 in your browser to see the excellent [GraphiQL interface](https://github.com/graphql/graphiql) and submit queries that way.  Need to customize where the GraphiQL interface points to?  Just adjust [the fetcher](https://github.com/danesparza/graphql_test/blob/master/static/index.html#L96).
