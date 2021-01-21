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
	}
	defer conn.Close()

	// Set up context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	graknClient := grakn.NewGraknClient(conn)

	databaseName := "Test"
	err = SetupDatabase(ctx, graknClient, databaseName)
	if err != nil {
		log.Fatalf("Can't setup database: %v", err)
	}

	metadata := map[string]string{}

	err = ExecuteSchema(graknClient, ctx, databaseName, metadata)
	if err != nil {
		log.Fatalf("Can't setup schema: %v", err)
	}

	err = ExecuteData(graknClient, ctx, databaseName, metadata)
	if err != nil {
		log.Fatalf("Can't setup data: %v", err)
	}
}

func SetupDatabase(ctx context.Context, graknClient grakn.GraknClient, databaseName string) (err error) {
	databaseAllResult, err := graknClient.DatabaseAll(ctx, &grakn.Database_All_Req{
	})

	if err != nil {
		log.Fatalf("could not list databases: %v", err)
		return err
	}
	log.Printf("databases: %v", databaseAllResult.Names)

	databaseContainsResult, err := graknClient.DatabaseContains(ctx, &grakn.Database_Contains_Req{
		Name: databaseName,
	})

	if err != nil {
		log.Fatalf("could not get database: %v", err)
		return err
	}

	if databaseContainsResult.Contains {
		log.Printf("database %v exists", databaseName)

		_, err = graknClient.DatabaseDelete(ctx, &grakn.Database_Delete_Req{
			Name: databaseName,
		})

		if err != nil {
			log.Fatalf("could not delete database %s: %v", databaseName, err)
			return err
		}
		log.Printf("database %s deleted", databaseName)
	} else {
		log.Printf("database %v does not exist", databaseName)
	}

	_, err = graknClient.DatabaseCreate(ctx, &grakn.Database_Create_Req{
		Name: databaseName,
	})

	if err != nil {
		log.Fatalf("could not create database %s: %v", databaseName, err)
		return err
	}
	log.Printf("database %s created", databaseName)

	databaseAllResult, err = graknClient.DatabaseAll(ctx, &grakn.Database_All_Req{
	})

	if err != nil {
		log.Fatalf("could not list databases: %v", err)
		return err
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

	transactionId := ksuid.New().String()

	schemaTransactionClient, err := graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(schemaTransactionClient, schemaSessionOpenResult.SessionId, transactionId, grakn.Transaction_WRITE, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	err = CreateTestDatabaseSchema(schemaTransactionClient, transactionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not create database schema: %w", err)
	}

	err = client.CommitTransaction(schemaTransactionClient, transactionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	err = client.CloseTransaction(schemaTransactionClient)
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	transactionId = ksuid.New().String()

	schemaTransactionClient, err = graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(schemaTransactionClient, schemaSessionOpenResult.SessionId, transactionId, grakn.Transaction_READ, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	err = GetTestDatabaseSchema(schemaTransactionClient, transactionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get database schema: %w", err)
	}

	err = client.CloseTransaction(schemaTransactionClient)
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
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

	transactionId := ksuid.New().String()

	dataTransactionClient, err := graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(dataTransactionClient, dataSessionOpenResult.SessionId, transactionId, grakn.Transaction_WRITE, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	err = CreateTestDatabaseData(dataTransactionClient, transactionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not create database data: %w", err)
	}

	err = client.CommitTransaction(dataTransactionClient, transactionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	err = client.CloseTransaction(dataTransactionClient)
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	transactionId = ksuid.New().String()

	dataTransactionClient, err = graknClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(dataTransactionClient, dataSessionOpenResult.SessionId, transactionId, grakn.Transaction_READ, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	err = GetTestDatabaseData(dataTransactionClient, transactionId, metadata, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	err = client.CloseTransaction(dataTransactionClient)
	if err != nil {
		return fmt.Errorf("could not close transaction: %w", err)
	}

	_, err = graknClient.SessionClose(ctx, &grakn.Session_Close_Req{
		SessionId: dataSessionOpenResult.SessionId,
	})
	if err != nil {
		return fmt.Errorf("could not close session: %w", err)
	}

	return err
}

func CreateTestDatabaseSchema(transactionClient grakn.Grakn_TransactionClient, transactionId string, metadata map[string]string, latencyMillis int32) (err error) {
	query, err := phone_data.GetPhoneCallsSchemaV2Gsql()
	if err != nil {
		return fmt.Errorf("could not get phone calls gsql: %w", err)
	}

	infer := true
	explain := true
	batchSize := int32(0)

	return client.RunDefineQuery(transactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
}


func GetTestDatabaseSchema(transactionClient grakn.Grakn_TransactionClient, transactionId string, metadata map[string]string, latencyMillis int32) (err error) {
	query := `
match 
	$x sub thing; 
get 
	$x;
`

	infer := true
	explain := true
	batchSize := int32(0)

	matchResponses, err := client.RunMatchQuery(transactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get database schema: %w", err)
	}

	for _, matchResponse := range matchResponses {
		answer := matchResponse.GetAnswer()
		log.Printf("database schema: %v", answer.String())
	}

	return err
}

func CreateTestDatabaseData(transactionClient grakn.Grakn_TransactionClient, transactionId string, metadata map[string]string, latencyMillis int32) (err error) {
	query, err := phone_data.GetPhoneCallsDataGsql()
	if err != nil {
		return fmt.Errorf("could not get phone calls gsql: %w", err)
	}

	infer := true
	explain := true
	batchSize := int32(0)

	insertResponses, err := client.RunInsertQuery(transactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not insert phone calls data: %w", err)
	}

	for _, insertResponse := range insertResponses {
		answer := insertResponse.GetAnswer()
		log.Printf("inserted: %v", answer.String())
	}

	return err
}

func GetTestDatabaseData(transactionClient grakn.Grakn_TransactionClient, transactionId string, metadata map[string]string, latencyMillis int32) (err error) {
	query := `
match
	$company isa company;
	$people isa people;
	$contract isa contract;
	$call isa call;
get $company;
`

	infer := true
	explain := true
	batchSize := int32(0)

	matchResponses, err := client.RunMatchQuery(transactionClient, transactionId, metadata, query, infer, explain, batchSize, latencyMillis)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	for _, matchResponse := range matchResponses {
		answer := matchResponse.GetAnswer()
		log.Printf("database schema: %v", answer.String())
	}

	return err
}
