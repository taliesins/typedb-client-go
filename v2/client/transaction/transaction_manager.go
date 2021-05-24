package transaction

import (
	"context"
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/taliesins/typedb-client-go/v2/client/common/rpc/request_builder/transaction"
	grakn "github.com/taliesins/typedb-client-go/v2/grakn_protocol"
)

type TransactionManager interface {
	OpenTransaction(sessionId []byte, transactionType grakn.Transaction_Type, options *grakn.Options, latencyMillis int32, metadata map[string]string) error
	CommitTransaction(metadata map[string]string) error
	RollbackTransaction(metadata map[string]string) error
	CloseTransaction() error
	Stream(requestId []byte, metadata map[string]string) (*grakn.Transaction_Res, error)
}

type transactionManagerImpl struct {
	TransactionClient grakn.GraknCore_TransactionClient
	Context context.Context
}

func NewTransactionManager(transactionClient grakn.GraknCore_TransactionClient, ctx context.Context) *transactionManagerImpl {
	return &transactionManagerImpl{
		TransactionClient: transactionClient,
		Context: ctx,
	}
}

func (q *transactionManagerImpl) OpenTransaction(sessionId []byte, transactionType grakn.Transaction_Type, options *grakn.Options, latencyMillis int32, metadata map[string]string) error {
	request := transaction.OpenReq(sessionId, transactionType, options, latencyMillis)
	request.Metadata = metadata

	response, err := q.query(request)
	if err != nil {
		return err
	}

	response.GetOpenRes()
	return nil
}

func (q *transactionManagerImpl) CommitTransaction(metadata map[string]string) error {
	request := transaction.CommitReq()
	request.Metadata = metadata

	response, err := q.query(request)
	if err != nil {
		return err
	}

	response.GetCommitRes()
	return nil
}

func (q *transactionManagerImpl) RollbackTransaction(metadata map[string]string) error {
	request := transaction.RollbackReq()
	request.Metadata = metadata

	response, err := q.query(request)
	if err != nil {
		return err
	}

	response.GetRollbackRes()
	return nil
}

func (q *transactionManagerImpl) CloseTransaction() error {
	return q.TransactionClient.CloseSend()
}

func (q *transactionManagerImpl) Stream(requestId []byte, metadata map[string]string) (*grakn.Transaction_Res, error) {
	request := transaction.StreamReq(requestId)
	request.Metadata = metadata

	response, err := q.query(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (q *transactionManagerImpl) query(req *grakn.Transaction_Req) (res *grakn.Transaction_Res, err error) { //Promise<QueryProto.Res> {
	requestId := ksuid.New().String()
	req.ReqId = []byte(requestId)

	err = q.TransactionClient.Send(
		&grakn.Transaction_Client{
			Reqs: []*grakn.Transaction_Req{
				req,
			},
		},
	)

	transactionServer, err := q.TransactionClient.Recv()
	if err != nil {
		return nil, fmt.Errorf("could not receive query response: %w", err)
	}
	return transactionServer.GetRes(), err
}