# graphql_test [![CircleCI](https://circleci.com/gh/danesparza/graphql_test.svg?style=svg)](https://circleci.com/gh/danesparza/graphql_test)
Simple Golang based GraphQL endpoint test using [graphql-go](https://github.com/graphql-go/graphql), [graphiql](https://github.com/graphql/graphiql), and a MySQL database

### Quick start
Install the database script in your MySQL database:
```
mysql -u USERNAME starwars < starwars.sql
```

Set the following environment variables for database connectivity:

To test, navigate to `http://localhost:8080` in your browser to see the excellent [GraphiQL interface](https://github.com/graphql/graphiql) and submit queries that way.  (Sidenote: Need to customize where the GraphiQL interface points to?  Just adjust [the fetcher](https://github.com/danesparza/graphql_test/blob/master/static/index.html#L96)).

To test in code, POST to `http://localhost:8080/graphql` with `Content-Type: application/graphql` and the [GraphQL](http://graphql.org/learn/) query:
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

