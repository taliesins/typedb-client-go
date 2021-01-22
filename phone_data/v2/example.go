package v2

import (
	"context"
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/taliesins/client-go/phone_data"
	"github.com/taliesins/client-go/v2/client"
	grakn "github.com/taliesins/client-go/v2/grakn_protocol"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Run() {
	address := ":1729"

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		panic(err)
	}
	defer conn.Close()

	// Set up context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	graknClient := grakn.NewGraknClient(conn)

	databaseName := "test"
	err = SetupDatabase(ctx, graknClient, databaseName)
	if err != nil {
		log.Fatalf("Can't setup database: %v", err)
		panic(err)
	}

	metadata := map[string]string{}

	err = ExecuteSchema(graknClient, ctx, databaseName, metadata)
	if err != nil {
		log.Fatalf("Can't setup schema: %v", err)
		panic(err)
	} else {
		log.Printf("Schema Complete")
	}

	err = ExecuteData(graknClient, ctx, databaseName, metadata)
	if err != nil {
		log.Fatalf("Can't setup data: %v", err)
		panic(err)
	} else {
		log.Printf("Data complete")
	}
}

func SetupDatabase(ctx context.Context, graknClient grakn.GraknClient, databaseName string) (err error) {
	databaseAllResult, err := graknClient.DatabaseAll(ctx, &grakn.Database_All_Req{
	})

	if err != nil {
		return fmt.Errorf("could not list databases: %w", err)
	}
	log.Printf("databases: %v", databaseAllResult.Names)

	databaseContainsResult, err := graknClient.DatabaseContains(ctx, &grakn.Database_Contains_Req{
		Name: databaseName,
	})
	if err != nil {
		return fmt.Errorf("could not get database: %w", err)
	}

	if databaseContainsResult.Contains {
		log.Printf("database %v exists", databaseName)

		_, err = graknClient.DatabaseDelete(ctx, &grakn.Database_Delete_Req{
			Name: databaseName,
		})

		if err != nil {
			return fmt.Errorf("could not delete database %s: %v", databaseName, err)
		}
		log.Printf("database %s deleted", databaseName)
	} else {
		log.Printf("database %v does not exist", databaseName)
	}

	_, err = graknClient.DatabaseCreate(ctx, &grakn.Database_Create_Req{
		Name: databaseName,
	})

	if err != nil {
		return fmt.Errorf("could not create database %s: %v", databaseName, err)
	}
	log.Printf("database %s created", databaseName)

	databaseAllResult, err = graknClient.DatabaseAll(ctx, &grakn.Database_All_Req{
	})

	if err != nil {
		return fmt.Errorf("could not list databases: %v", err)
	}

	log.Printf("databases: %v", databaseAllResult.Names)
	return err
}

func ExecuteSchema(graknClient grakn.GraknClient, ctx context.Context, databaseName string, metadata map[string]string) (err error) {
	latencyMillis := int32(100)

	schemaSessionOpenResult, err := graknClient.SessionOpen(ctx, &grakn.Session_Open_Req{
		Database: databaseName,
		Type:     grakn.Session_SCHEMA,
		Options:  &grakn.Options{},
	})
	if err != nil {
		return fmt.Errorf("could not open session: %w", err)
	} else {
		log.Printf("Opened session: %v", schemaSessionOpenResult.SessionId)
	}

	err = CreateTestDatabaseSchema(graknClient, ctx, schemaSessionOpenResult.SessionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not create schema: %w", err)
	}

	err = GetTestDatabaseSchema(graknClient, ctx, schemaSessionOpenResult.SessionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get schema: %w", err)
	}

	_, err = graknClient.SessionClose(ctx, &grakn.Session_Close_Req{
		SessionId: schemaSessionOpenResult.SessionId,
	})
	if err != nil {
		return fmt.Errorf("could not close session: %w", err)
	}

	return err
}

func ExecuteData(graknClient grakn.GraknClient, ctx context.Context, databaseName string, metadata map[string]string) (err error){
	latencyMillis := int32(100)

	dataSessionOpenResult, err := graknClient.SessionOpen(ctx, &grakn.Session_Open_Req{
		Database: databaseName,
		Type:     grakn.Session_DATA,
		Options:  &grakn.Options{},
	})
	if err != nil {
		return fmt.Errorf("could not open session: %w", err)
	} else {
		log.Printf("Opened session: %v", dataSessionOpenResult.SessionId)
	}

	err = CreateTestDatabaseData(graknClient, ctx, 	dataSessionOpenResult.SessionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not create data: %w", err)
	}

	err = GetTestDatabaseData(graknClient, ctx, dataSessionOpenResult.SessionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get data: %w", err)
	}

	_, err = graknClient.SessionClose(ctx, &grakn.Session_Close_Req{
		SessionId: dataSessionOpenResult.SessionId,
	})
	if err != nil {
		return fmt.Errorf("could not close session: %w", err)
	}

	return err
}

func CreateTestDatabaseSchema(graknClient grakn.GraknClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	query, err := phone_data.GetPhoneCallsSchemaV2Gql()
	if err != nil {
		return fmt.Errorf("could not get phone calls gsql: %w", err)
	}

	infer := true
	explain := true
	batchSize := int32(0)

	transactionId := ksuid.New().String()

	schemaTransactionClient, err := graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(schemaTransactionClient, sessionId, transactionId, grakn.Transaction_WRITE, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	err = client.RunDefineQuery(schemaTransactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not run define query: %w", err)
	}

	err = client.CommitTransaction(schemaTransactionClient, transactionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	err = client.CloseTransaction(schemaTransactionClient)
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	return err
}

func GetTestDatabaseSchema(graknClient grakn.GraknClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	query := `
match 
	$x sub thing; 
get 
	$x;
`

	infer := true
	explain := true
	batchSize := int32(0)

	transactionId := ksuid.New().String()

	schemaTransactionClient, err := graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(schemaTransactionClient, sessionId, transactionId, grakn.Transaction_READ, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	matchResponses, err := client.RunMatchQuery(schemaTransactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get database schema: %w", err)
	}

	for _, matchResponse := range matchResponses {
		answers := matchResponse.GetAnswers()
		log.Printf("database schema: %v", answers)
	}

	err = client.CloseTransaction(schemaTransactionClient)
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	return err
}

func CreateTestDatabaseData(graknClient grakn.GraknClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	infer := true
	explain := true
	batchSize := int32(0)

	gql, err := phone_data.GetPhoneCallsDataGql()
	if err != nil {
		return fmt.Errorf("could not get phone calls data gql: %w", err)
	}

	for _, query := range gql {
		transactionId := ksuid.New().String()

		transactionClient, err := graknClient.Transaction(ctx)
		if err != nil {
			return fmt.Errorf("could not create transaction client: %w", err)
		}

		err = client.OpenTransaction(transactionClient, sessionId, transactionId, grakn.Transaction_WRITE, metadata, latencyMillis)
		if err != nil {
			return fmt.Errorf("could not open transaction: %w", err)
		}

		insertResponses, err := client.RunInsertQuery(transactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
		if err != nil {
			return fmt.Errorf("could not insert phone calls data: %w", err)
		}

		err = client.CommitTransaction(transactionClient, transactionId, metadata, latencyMillis)
		if err != nil {
			return fmt.Errorf("could not commit transaction: %w", err)
		}

		err = client.CloseTransaction(transactionClient)
		if err != nil {
			return fmt.Errorf("could not close transaction: %w", err)
		}

		for _, insertResponse := range insertResponses {
			answers := insertResponse.GetAnswers()
			log.Printf("inserted: %v", answers)
		}
	}

	return err
}

func GetTestDatabaseData(graknClient grakn.GraknClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	infer := false
	explain := false
	batchSize := int32(0)

	transactionId := ksuid.New().String()

	dataTransactionClient, err := graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(dataTransactionClient, sessionId, transactionId, grakn.Transaction_READ, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	query := `
match
	$person isa person;
get $person; 
count;
`

	matchAggregateResponses, err := client.RunMatchAggregateQuery(dataTransactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get database data: %v \r\n %w", query, err)
	}

	for _, matchResponse := range matchAggregateResponses {
		answer := matchResponse.GetAnswer()
		log.Printf("Number of people: %v", answer.GetLongValue())
	}

	query = `
match
	$company isa company;
get $company; 
count;
`

	matchAggregateResponses, err = client.RunMatchAggregateQuery(dataTransactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	for _, matchResponse := range matchAggregateResponses {
		answer := matchResponse.GetAnswer()
		log.Printf("Number of companies: %v", answer.GetLongValue())
	}

	query = `
match
	$contract isa contract;
get $contract; 
count;
`

	matchAggregateResponses, err = client.RunMatchAggregateQuery(dataTransactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	for _, matchResponse := range matchAggregateResponses {
		answer := matchResponse.GetAnswer()
		log.Printf("Number of contracts: %v", answer.GetLongValue())
	}

	query = `
match
	$call isa call;
get $call; 
count;
`

	matchAggregateResponses, err = client.RunMatchAggregateQuery(dataTransactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	for _, matchResponse := range matchAggregateResponses {
		answer := matchResponse.GetAnswer()
		log.Printf("Number of calls: %v", answer.GetLongValue())
	}

	err = client.CloseTransaction(dataTransactionClient)
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	return err
}
