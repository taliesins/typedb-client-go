package main

import (
	"context"
	grakn "github.com/taliesins/client-go/grakn_protocol_v1"
	grpc "google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	address := ":48555"
	username := ""
	password := ""

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Set up context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	client := grakn.NewKeyspaceServiceClient(conn)

	databaseAllResult, err := client.Retrieve(ctx, &grakn.Keyspace_Retrieve_Req{
		Username: username,
		Password: password,
	})
	if err != nil {
		log.Fatalf("could not list databases: %v", err)
	}
	log.Printf("databases: %v", databaseAllResult.Names)
}