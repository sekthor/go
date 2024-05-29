package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func main() {
	ctx := context.Background()
	// URI examples: "neo4j://localhost", "neo4j+s://xxx.databases.neo4j.io"
	dbUri := os.Getenv("DB_URI")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	driver, err := neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))

	if err != nil {
		log.Fatal("could not connect to database")
	}

	defer driver.Close(ctx)

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}

	result, err := neo4j.ExecuteQuery(ctx, driver,
		"MATCH (r:Recipe {name: $name}) RETURN r",
		map[string]any{
			"name": "Lentil Bolognese",
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"),
	)

	for _, record := range result.Records {
		fmt.Println(record.AsMap())
	}
}
