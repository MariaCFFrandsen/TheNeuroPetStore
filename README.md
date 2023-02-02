# TheNeuroPetStore

This is a simple project consisting of a Swagger API and postgres database. 
The current functionality allows the user to get all pets in the database through the Swagger API. The Swagger API is generated from a yaml spec found in api/swagger.yaml

In order to implement further endpoints, expand this yaml file and re-generate operators, models and server.
This can be done by running the command found in api/generator.go

```
//go:generate swagger generate server -f swagger.yaml --default-scheme http
```

You can start the server with the command with a specific port with:

```
go run cmd/main.go --scheme http --port=8080
```

You can start a postgres database with Docker with:

```
docker run --name postgres-db -e POSTGRES_PASSWORD=docker -p 5432:5432 -d postgres
```

Specify name and password as you please.

You can connect to the database through DataGrip with the following information:

```
Host: localhost
Port: 5432
User: postgres
Password: your-password
```

In order to query the database, this project uses github.com/kyleconroy/sqlc .
This requires that you expand the internal/sql/queries.sql with your queries in postgres dialect.
You regenerate models with the below command also found in internal/sql/generator.go

```
//go:generate go run github.com/kyleconroy/sqlc/cmd/sqlc@latest generate
```