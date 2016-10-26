# graphql_test
Simple Golang based GraphQL endpoint test

To test, POST to `http://localhost:8080/graphql` with `Content-Type: application/graphql` and the body:
```
{
	"data": {
		"latestPost": "Hello World!"
	}
}
```
