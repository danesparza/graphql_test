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
