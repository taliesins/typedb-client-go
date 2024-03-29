package v2

import (
	"context"
	"fmt"
	"github.com/taliesins/typedb-client-go/phone_data"
	"github.com/taliesins/typedb-client-go/v2/client/query"
	"github.com/taliesins/typedb-client-go/v2/client/transaction"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	typeDBClient := typedb_protocol.NewTypeDBClient(conn)

	databaseName := "test"

	err = SetupDatabase(ctx, typeDBClient, databaseName)
	if err != nil {
		log.Fatalf("Can't setup database: %v", err)
		panic(err)
	}

	metadata := map[string]string{}

	err = ExecuteSchema(typeDBClient, ctx, databaseName, metadata)
	if err != nil {
		log.Fatalf("Can't setup schema: %v", err)
		panic(err)
	} else {
		log.Printf("Schema Complete")
	}

	err = ExecuteData(typeDBClient, ctx, databaseName, metadata)
	if err != nil {
		log.Fatalf("Can't setup data: %v", err)
		panic(err)
	} else {
		log.Printf("Data complete")
	}
}

func SetupDatabase(ctx context.Context, typeDBClient typedb_protocol.TypeDBClient, databaseName string) (err error) {
	databaseAllResult, err := typeDBClient.DatabasesAll(ctx, &typedb_protocol.CoreDatabaseManager_All_Req{})

	if err != nil {
		return fmt.Errorf("could not list databases: %w", err)
	}
	log.Printf("databases: %v", databaseAllResult.Names)

	databaseContainsResult, err := typeDBClient.DatabasesContains(ctx, &typedb_protocol.CoreDatabaseManager_Contains_Req{
		Name: databaseName,
	})
	if err != nil {
		return fmt.Errorf("could not get database: %w", err)
	}

	if databaseContainsResult.Contains {
		log.Printf("database %v exists", databaseName)

		_, err = typeDBClient.DatabaseDelete(ctx, &typedb_protocol.CoreDatabase_Delete_Req{
			Name: databaseName,
		})

		if err != nil {
			return fmt.Errorf("could not delete database %s: %v", databaseName, err)
		}
		log.Printf("database %s deleted", databaseName)
	} else {
		log.Printf("database %v does not exist", databaseName)
	}

	_, err = typeDBClient.DatabasesCreate(ctx, &typedb_protocol.CoreDatabaseManager_Create_Req{
		Name: databaseName,
	})

	if err != nil {
		return fmt.Errorf("could not create database %s: %v", databaseName, err)
	}
	log.Printf("database %s created", databaseName)

	databaseAllResult, err = typeDBClient.DatabasesAll(ctx, &typedb_protocol.CoreDatabaseManager_All_Req{})

	if err != nil {
		return fmt.Errorf("could not list databases: %v", err)
	}

	log.Printf("databases: %v", databaseAllResult.Names)
	return err
}

func ExecuteSchema(typeDBClient typedb_protocol.TypeDBClient, ctx context.Context, databaseName string, metadata map[string]string) (err error) {
	latencyMillis := int32(100)

	schemaSessionOpenResult, err := typeDBClient.SessionOpen(ctx, &typedb_protocol.Session_Open_Req{
		Database: databaseName,
		Type:     typedb_protocol.Session_SCHEMA,
		Options:  &typedb_protocol.Options{},
	})
	if err != nil {
		return fmt.Errorf("could not open session: %w", err)
	} else {
		log.Printf("Opened session: %v", schemaSessionOpenResult.SessionId)
	}

	err = CreateTestDatabaseSchema(typeDBClient, ctx, schemaSessionOpenResult.SessionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not create schema: %w", err)
	}

	err = GetTestDatabaseSchema(typeDBClient, ctx, schemaSessionOpenResult.SessionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get schema: %w", err)
	}

	_, err = typeDBClient.SessionClose(ctx, &typedb_protocol.Session_Close_Req{
		SessionId: schemaSessionOpenResult.SessionId,
	})
	if err != nil {
		return fmt.Errorf("could not close session: %w", err)
	}

	return err
}

func ExecuteData(typeDBClient typedb_protocol.TypeDBClient, ctx context.Context, databaseName string, metadata map[string]string) (err error) {
	latencyMillis := int32(100)

	dataSessionOpenResult, err := typeDBClient.SessionOpen(ctx, &typedb_protocol.Session_Open_Req{
		Database: databaseName,
		Type:     typedb_protocol.Session_DATA,
		Options:  &typedb_protocol.Options{
			TransactionTimeoutOpt: &typedb_protocol.Options_TransactionTimeoutMillis{
				TransactionTimeoutMillis: 1000*120,
			},
			SessionIdleTimeoutOpt: &typedb_protocol.Options_SessionIdleTimeoutMillis{
				SessionIdleTimeoutMillis: 1000*240,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("could not open session: %w", err)
	} else {
		log.Printf("Opened session: %v", dataSessionOpenResult.SessionId)
	}

	err = CreateTestDatabaseData(typeDBClient, ctx, dataSessionOpenResult.SessionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not create data: %w", err)
	}

	err = GetTestDatabaseData(typeDBClient, ctx, dataSessionOpenResult.SessionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get data: %w", err)
	}

	_, err = typeDBClient.SessionClose(ctx, &typedb_protocol.Session_Close_Req{
		SessionId: dataSessionOpenResult.SessionId,
	})
	if err != nil {
		return fmt.Errorf("could not close session: %w", err)
	}

	return err
}

func CreateTestDatabaseSchema(typeDBClient typedb_protocol.TypeDBClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	gqlQuery, err := phone_data.GetPhoneCallsSchemaV2Gql()
	if err != nil {
		return fmt.Errorf("could not get phone calls gsql: %w", err)
	}

	transactionClient, err := typeDBClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	transactionOptions := &typedb_protocol.Options{
		InferOpt: &typedb_protocol.Options_Infer{
			Infer: true,
		},
		ExplainOpt: &typedb_protocol.Options_Explain{
			Explain: false,
		},
	}
	transactionManager := transaction.NewTransactionManager(transactionClient, ctx)

	err = transactionManager.OpenTransaction(sessionId, typedb_protocol.Transaction_WRITE, transactionOptions, latencyMillis, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	queryOptions := &typedb_protocol.Options{
		ExplainOpt: &typedb_protocol.Options_Explain{
			Explain: false,
		},
	}
	queryManager := query.NewQueryManager(transactionClient, ctx)

	err = queryManager.Define(gqlQuery, queryOptions, metadata)
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

func GetTestDatabaseSchema(typeDBClient typedb_protocol.TypeDBClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	gqlQuery := `
match 
	$x sub thing; 
get 
	$x;
`

	transactionClient, err := typeDBClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	transactionOptions := &typedb_protocol.Options{
		ExplainOpt: &typedb_protocol.Options_Explain{
			Explain: true,
		},
	}
	transactionManager := transaction.NewTransactionManager(transactionClient, ctx)

	err = transactionManager.OpenTransaction(sessionId, typedb_protocol.Transaction_READ, transactionOptions, latencyMillis, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	queryManager := query.NewQueryManager(transactionClient, ctx)

	queryOptions := &typedb_protocol.Options{
		ExplainOpt: &typedb_protocol.Options_Explain{
			Explain: true,
		},
	}
	matchResponses, err := queryManager.Match(gqlQuery, queryOptions, metadata)
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

func CreateTestDatabaseData(typeDBClient typedb_protocol.TypeDBClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	gqlQueries, err := phone_data.GetPhoneCallsDataGql()
	if err != nil {
		return fmt.Errorf("could not get phone calls data gql: %w", err)
	}

	transactionClient, err := typeDBClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	transactionOptions := &typedb_protocol.Options{
		InferOpt: &typedb_protocol.Options_Infer{
			Infer: true,
		},
		ExplainOpt: &typedb_protocol.Options_Explain{
			Explain: false,
		},
	}
	transactionManager := transaction.NewTransactionManager(transactionClient, ctx)

	err = transactionManager.OpenTransaction(sessionId, typedb_protocol.Transaction_WRITE, transactionOptions, latencyMillis, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	queryManager := query.NewQueryManager(transactionClient, ctx)

	queryOptions := &typedb_protocol.Options{
		ExplainOpt: &typedb_protocol.Options_Explain{
			Explain: false,
		},
	}

	for _, gqlQuery := range gqlQueries {
		insertResponses, err := queryManager.Insert(gqlQuery, queryOptions, metadata)
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

func GetTestDatabaseData(typeDBClient typedb_protocol.TypeDBClient, ctx context.Context, sessionId []byte, metadata map[string]string, latencyMillis int32) (err error) {
	gqlQuery := `
match
	$person isa person;
get $person; 
count;
`

	transactionClient, err := typeDBClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	transactionOptions := &typedb_protocol.Options{
		InferOpt: &typedb_protocol.Options_Infer{
			Infer: true,
		},
		ExplainOpt: &typedb_protocol.Options_Explain{
			Explain: true,
		},
	}
	transactionManager := transaction.NewTransactionManager(transactionClient, ctx)

	err = transactionManager.OpenTransaction(sessionId, typedb_protocol.Transaction_READ, transactionOptions, latencyMillis, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	queryManager := query.NewQueryManager(transactionClient, ctx)

	queryOptions := &typedb_protocol.Options{
		ExplainOpt: &typedb_protocol.Options_Explain{
			Explain: false,
		},
	}

	matchAggregateResponses, err := queryManager.MatchAggregate(gqlQuery, queryOptions, metadata)
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
	matchAggregateResponses, err = queryManager.MatchAggregate(gqlQuery, queryOptions, metadata)
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

	matchAggregateResponses, err = queryManager.MatchAggregate(gqlQuery, queryOptions, metadata)
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

	matchAggregateResponses, err = queryManager.MatchAggregate(gqlQuery, queryOptions, metadata)
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
