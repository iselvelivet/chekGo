package main

import (
	"context"
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	log.Println("Server is running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))

	ctx := context.Background()

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "yourpassword"), // Your credentials
	})
	if err != nil {
		log.Fatalf("Failed to create a client: %v", err)
	}

	// Test the connection by listing databases
	dbs, err := c.Databases(ctx)
	if err != nil {
		log.Fatalf("Failed to list databases: %v", err)
	}
	fmt.Println("Databases:", dbs)

}

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"hello": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "world", nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)
