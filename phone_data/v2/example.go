package v2

import (
	"context"
	"fmt"
	"github.com/taliesins/typedb-client-go/phone_data"
	"github.com/taliesins/typedb-client-go/v2/client/query"
	"github.com/taliesins/typedb-client-go/v2/client/transaction"
	grakn "github.com/taliesins/typedb-client-go/v2/grakn_protocol"
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

	graknClient := grakn.NewGraknCoreClient(conn)

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

func SetupDatabase(ctx context.Context, graknClient grakn.GraknCoreClient, databaseName string) (err error) {
	databaseAllResult, err := graknClient.DatabasesAll(ctx, &grakn.CoreDatabaseManager_All_Req{})

	if err != nil {
		return fmt.Errorf("could not list databases: %w", err)
	}
	log.Printf("databases: %v", databaseAllResult.Names)

	databaseContainsResult, err := graknClient.DatabasesContains(ctx, &grakn.CoreDatabaseManager_Contains_Req{
		Name: databaseName,
	})
	if err != nil {
		return fmt.Errorf("could not get database: %w", err)
	}

	if databaseContainsResult.Contains {
		log.Printf("database %v exists", databaseName)

		_, err = graknClient.DatabaseDelete(ctx, &grakn.CoreDatabase_Delete_Req{
			Name: databaseName,
		})

		if err != nil {
			return fmt.Errorf("could not delete database %s: %v", databaseName, err)
		}
		log.Printf("database %s deleted", databaseName)
	} else {
		log.Printf("database %v does not exist", databaseName)
	}

	_, err = graknClient.DatabasesCreate(ctx, &grakn.CoreDatabaseManager_Create_Req{
		Name: databaseName,
	})

	if err != nil {
		return fmt.Errorf("could not create database %s: %v", databaseName, err)
	}
	log.Printf("database %s created", databaseName)

	databaseAllResult, err = graknClient.DatabasesAll(ctx, &grakn.CoreDatabaseManager_All_Req{})

	if err != nil {
		return fmt.Errorf("could not list databases: %v", err)
	}

	log.Printf("databases: %v", databaseAllResult.Names)
	return err
}

func ExecuteSchema(graknClient grakn.GraknCoreClient, ctx context.Context, databaseName string, metadata map[string]string) (err error) {
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

func ExecuteData(graknClient grakn.GraknCoreClient, ctx context.Context, databaseName string, metadata map[string]string) (err error) {
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

	err = CreateTestDatabaseData(graknClient, ctx, dataSessionOpenResult.SessionId, metadata, latencyMillis)
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

func CreateTestDatabaseSchema(graknClient grakn.GraknCoreClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	gqlQuery, err := phone_data.GetPhoneCallsSchemaV2Gql()
	if err != nil {
		return fmt.Errorf("could not get phone calls gsql: %w", err)
	}

	transactionClient, err := graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	options := &grakn.Options{
		InferOpt: &grakn.Options_Infer{
			Infer: true,
		},
		ExplainOpt: &grakn.Options_Explain{
			Explain: true,
		},
		BatchSizeOpt: &grakn.Options_BatchSize{
			BatchSize: int32(0),
		},
	}
	transactionManager := transaction.NewTransactionManager(transactionClient, ctx)

	err = transactionManager.OpenTransaction(sessionId, grakn.Transaction_WRITE, options, latencyMillis, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	queryManager := query.NewQueryManager(transactionClient, ctx)

	err = queryManager.Define(gqlQuery, options, metadata)
	if err != nil {
		return fmt.Errorf("could not run define query: %w", err)
	}

	err = transactionManager.CommitTransaction(metadata)
	if err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	err = transactionManager.CloseTransaction()
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	return err
}

func GetTestDatabaseSchema(graknClient grakn.GraknCoreClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	gqlQuery := `
match 
	$x sub thing; 
get 
	$x;
`

	transactionClient, err := graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	options := &grakn.Options{
		InferOpt: &grakn.Options_Infer{
			Infer: true,
		},
		ExplainOpt: &grakn.Options_Explain{
			Explain: true,
		},
		BatchSizeOpt: &grakn.Options_BatchSize{
			BatchSize: int32(0),
		},
	}
	transactionManager := transaction.NewTransactionManager(transactionClient, ctx)

	err = transactionManager.OpenTransaction(sessionId, grakn.Transaction_READ, options, latencyMillis, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	queryManager := query.NewQueryManager(transactionClient, ctx)

	matchResponses, err := queryManager.Match(gqlQuery, options, metadata)
	if err != nil {
		return fmt.Errorf("could not run match query: %w", err)
	}

	for _, matchResponse := range matchResponses {
		log.Printf("database schema: %v", matchResponse)
	}

	err = transactionManager.CloseTransaction()
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	return err
}

func CreateTestDatabaseData(graknClient grakn.GraknCoreClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	gqlQueries, err := phone_data.GetPhoneCallsDataGql()
	if err != nil {
		return fmt.Errorf("could not get phone calls data gql: %w", err)
	}

	transactionClient, err := graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	options := &grakn.Options{
		InferOpt: &grakn.Options_Infer{
			Infer: true,
		},
		ExplainOpt: &grakn.Options_Explain{
			Explain: true,
		},
		BatchSizeOpt: &grakn.Options_BatchSize{
			BatchSize: int32(0),
		},
	}
	transactionManager := transaction.NewTransactionManager(transactionClient, ctx)

	err = transactionManager.OpenTransaction(sessionId, grakn.Transaction_WRITE, options, latencyMillis, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	queryManager := query.NewQueryManager(transactionClient, ctx)

	for _, gqlQuery := range gqlQueries {
		insertResponses, err := queryManager.Insert(gqlQuery, options, metadata)
		if err != nil {
			return fmt.Errorf("could not insert phone calls data: %w", err)
		}

		for _, insertResponse := range insertResponses {
			log.Printf("inserted: %v", insertResponse)
		}
	}

	err = transactionManager.CommitTransaction(metadata)
	if err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	err = transactionManager.CloseTransaction()
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	return err
}

func GetTestDatabaseData(graknClient grakn.GraknCoreClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	gqlQuery := `
match
	$person isa person;
get $person; 
count;
`

	transactionClient, err := graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	options := &grakn.Options{
		InferOpt: &grakn.Options_Infer{
			Infer: true,
		},
		ExplainOpt: &grakn.Options_Explain{
			Explain: true,
		},
		BatchSizeOpt: &grakn.Options_BatchSize{
			BatchSize: int32(0),
		},
	}
	transactionManager := transaction.NewTransactionManager(transactionClient, ctx)

	err = transactionManager.OpenTransaction(sessionId, grakn.Transaction_READ, options, latencyMillis, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	queryManager := query.NewQueryManager(transactionClient, ctx)

	matchAggregateResponses, err := queryManager.MatchAggregate(gqlQuery, options, metadata)
	if err != nil {
		return fmt.Errorf("could not get database data: %v \r\n %w", gqlQuery, err)
	}
	log.Printf("Number of people: %v", matchAggregateResponses.GetLongValue())


	gqlQuery = `
match
	$company isa company;
get $company; 
count;
`
	matchAggregateResponses, err = queryManager.MatchAggregate(gqlQuery, options, metadata)
	if err != nil {
		return fmt.Errorf("could not get database data: %v \r\n %w", gqlQuery, err)
	}
	log.Printf("Number of companies: %v", matchAggregateResponses.GetLongValue())

	gqlQuery = `
match
	$contract isa contract;
get $contract; 
count;
`

	matchAggregateResponses, err = queryManager.MatchAggregate(gqlQuery, options, metadata)
	if err != nil {
		return fmt.Errorf("could not get database data: %v \r\n %w", gqlQuery, err)
	}
	log.Printf("Number of contracts: %v", matchAggregateResponses.GetLongValue())

	gqlQuery = `
match
	$call isa call;
get $call; 
count;
`

	matchAggregateResponses, err = queryManager.MatchAggregate(gqlQuery, options, metadata)
	if err != nil {
		return fmt.Errorf("could not get database data: %v \r\n %w", gqlQuery, err)
	}
	log.Printf("Number of calls: %v", matchAggregateResponses.GetLongValue())

	err = transactionManager.CloseTransaction()
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	return err
}
