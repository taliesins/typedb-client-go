package v1

import (
	"context"
	"fmt"
	"github.com/taliesins/client-go/phone_data"
	"github.com/taliesins/client-go/v1/client"
	grakn "github.com/taliesins/client-go/v1/grakn_protocol"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Run() {
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	databaseName := "test"
	keyspaceServiceClient := grakn.NewKeyspaceServiceClient(conn)

	err = SetupDatabase(keyspaceServiceClient, ctx, username, password, databaseName)
	if err != nil {
		log.Fatalf("Can't setup database: %v", err)
	}

	metadata := map[string]string{}

	sessionClient := grakn.NewSessionServiceClient(conn)

	sessionOpenResult, err := sessionClient.Open(ctx, &grakn.Session_Open_Req{
		Username: username,
		Password: password,
		Keyspace: databaseName,
		Metadata: metadata,
	})
	if err != nil {
		log.Fatalf("could not open session: %v", err)
	}

	sessionId := sessionOpenResult.SessionId

	defer func() {
		_, err = sessionClient.Close(ctx, &grakn.Session_Close_Req{
			SessionId: sessionId,
			Metadata:  metadata,
		})
		if err != nil {
			log.Fatalf("could not close session: %v", err)
		}
	}()

	err = ExecuteSchema(sessionClient, ctx, sessionId, metadata)
	if err != nil {
		log.Fatalf("Can't setup schema: %v", err)
	}

	err = ExecuteData(sessionClient, ctx, sessionId, metadata)
	if err != nil {
		log.Fatalf("Can't setup data: %v", err)
	}
}

func ExecuteSchema(sessionClient grakn.SessionServiceClient, ctx context.Context, sessionId string, metadata map[string]string) (err error) {
	transactionClient, err := sessionClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(transactionClient, sessionId, grakn.Transaction_WRITE, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	err = CreateTestDatabaseSchema(transactionClient, metadata)
	if err != nil {
		return fmt.Errorf("could not create database schema: %w", err)
	}

	err = client.CommitTransaction(transactionClient, metadata)
	if err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	err = client.CloseTransaction(transactionClient)
	if err != nil {
		return fmt.Errorf("could not create database schema: %w", err)
	}

	transactionClient, err = sessionClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(transactionClient, sessionId, grakn.Transaction_READ, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	err = GetTestDatabaseSchema(transactionClient, metadata)
	if err != nil {
		return fmt.Errorf("could not get database schema: %w", err)
	}

	err = client.CloseTransaction(transactionClient)
	if err != nil {
		return fmt.Errorf("could not create database schema: %w", err)
	}

	return nil
}

func ExecuteData(sessionClient grakn.SessionServiceClient, ctx context.Context, sessionId string, metadata map[string]string) (err error) {
	transactionClient, err := sessionClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(transactionClient, sessionId, grakn.Transaction_WRITE, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	err = CreateTestDatabaseData(transactionClient, metadata)
	if err != nil {
		return fmt.Errorf("could not create database data: %w", err)
	}

	err = client.CommitTransaction(transactionClient, metadata)
	if err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}

	err = client.CloseTransaction(transactionClient)
	if err != nil {
		return fmt.Errorf("could not create database schema: %w", err)
	}

	transactionClient, err = sessionClient.Transaction(ctx)
	if err != nil {
		return fmt.Errorf("could not create transaction client: %w", err)
	}

	err = client.OpenTransaction(transactionClient, sessionId, grakn.Transaction_READ, metadata)
	if err != nil {
		return fmt.Errorf("could not open transaction: %w", err)
	}

	err = GetTestDatabaseData(transactionClient, metadata)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	err = client.CloseTransaction(transactionClient)
	if err != nil {
		return fmt.Errorf("could not create database schema: %w", err)
	}
	return nil
}

func SetupDatabase(keyspaceServiceClient grakn.KeyspaceServiceClient, ctx context.Context, username string, password string, databaseName string) (err error) {
	databaseAllResult, err := keyspaceServiceClient.Retrieve(ctx, &grakn.Keyspace_Retrieve_Req{
		Username: username,
		Password: password,
	})
	if err != nil {
		return fmt.Errorf("could not list databases: %w", err)
	}
	log.Printf("databases: %v", databaseAllResult.Names)

	if contains(databaseAllResult.Names, databaseName) {
		log.Printf("database %v exists", databaseName)

		_, err = keyspaceServiceClient.Delete(ctx, &grakn.Keyspace_Delete_Req{
			Username: username,
			Password: password,
			Name: databaseName,
		})

		if err != nil {
			return fmt.Errorf("could not delete database %s: %w", databaseName, err)
		}
	}
	return nil
}

func CreateTestDatabaseSchema(transactionClient grakn.SessionService_TransactionClient, metadata map[string]string) (err error) {
	query, err := phone_data.GetPhoneCallsSchemaV1Gql()
	if err != nil {
		return fmt.Errorf("could not get phone calls gsql: %w", err)
	}

	infer := true
	explain := true
	batchSize := int32(0)

	_, err =  client.RunQuery(transactionClient, metadata, query, infer, explain, batchSize)
	if err != nil {
		return fmt.Errorf("could not get phone calls gsql: %w", err)
	}

	return err
}


func GetTestDatabaseSchema(transactionClient grakn.SessionService_TransactionClient, metadata map[string]string) (err error) {
	query := `
match 
	$x sub thing; 
get 
	$x;
`

	infer := true
	explain := true
	batchSize := int32(0)

	answers, err := client.RunQuery(transactionClient, metadata, query, infer, explain, batchSize)
	if err != nil {
		return fmt.Errorf("could not get database schema: %w", err)
	}

	for _, answer := range answers {
		log.Printf("database schema: %v", answer.String())
	}

	return err
}


func CreateTestDatabaseData(transactionClient grakn.SessionService_TransactionClient, metadata map[string]string) (err error) {
	infer := true
	explain := true
	batchSize := int32(0)

	gql, err := phone_data.GetPhoneCallsDataGql()
	if err != nil {
		return fmt.Errorf("could not get phone calls data gql: %w", err)
	}

	for _, query := range gql {
		answers, err := client.RunQuery(transactionClient, metadata, query, infer, explain, batchSize)
		if err != nil {
			return fmt.Errorf("could not insert phone calls data: %w", err)
		}

		for _, answer := range answers {
			log.Printf("inserted: %v", answer.String())
		}
	}

	return err
}

func GetTestDatabaseData(transactionClient grakn.SessionService_TransactionClient, metadata map[string]string) (err error) {
	infer := true
	explain := true
	batchSize := int32(0)

	query := `
match
	$person isa person;
get; 
count;
`

	answers, err := client.RunQuery(transactionClient, metadata, query, infer, explain, batchSize)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	for _, answer := range answers {
		log.Printf("Number of people: %v", answer.String())
	}

	query = `
match
	$company isa company;
get; 
count;
`

	answers, err = client.RunQuery(transactionClient, metadata, query, infer, explain, batchSize)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	for _, answer := range answers {
		log.Printf("Number of companies: %v", answer.String())
	}

	query = `
match
	$contract isa contract;
get; 
count;
`

	answers, err = client.RunQuery(transactionClient, metadata, query, infer, explain, batchSize)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	for _, answer := range answers {
		log.Printf("Number of contracts: %v", answer.String())
	}

	query = `
match
	$call isa call;
get; 
count;
`

	answers, err = client.RunQuery(transactionClient, metadata, query, infer, explain, batchSize)
	if err != nil {
		return fmt.Errorf("could not get database data: %w", err)
	}

	for _, answer := range answers {
		log.Printf("Number of calls: %v", answer.String())
	}

	return err
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func printAttribute(attribute *grakn.Concept) {
	attributeValueType := attribute.GetValueTypeRes().GetValueType()
	log.Printf("value type : %v", attributeValueType)
	log.Printf("label : %v", attribute.GetTypeRes().GetType().GetLabelRes().GetLabel())

	attributeValue := attribute.GetValueRes().GetValue()

	if attributeValue != nil {
		switch attributeValueType {
		case grakn.AttributeType_STRING:
			log.Printf("value : %v", attributeValue.GetString_())
		case grakn.AttributeType_BOOLEAN:
			log.Printf("value : %v", attributeValue.GetBoolean())
		case grakn.AttributeType_INTEGER:
			log.Printf("value : %v", attributeValue.GetInteger())
		case grakn.AttributeType_LONG:
			log.Printf("value : %v", attributeValue.GetLong())
		case grakn.AttributeType_FLOAT:
			log.Printf("value : %v", attributeValue.GetFloat())
		case grakn.AttributeType_DOUBLE:
			log.Printf("value : %v", attributeValue.GetDouble())
		case grakn.AttributeType_DATETIME:
			log.Printf("value : %v", attributeValue.GetDatetime())
		default:

		}
	}
}
