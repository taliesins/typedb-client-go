package transaction

import (
	"context"
	"github.com/taliesins/typedb-client-go/v2/client/common"
	"github.com/taliesins/typedb-client-go/v2/client/common/rpc/request_builder/transaction"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

type TransactionManager interface {
	OpenTransaction(sessionId []byte, transactionType typedb_protocol.Transaction_Type, options *typedb_protocol.Options, latencyMillis int32, metadata map[string]string) error
	CommitTransaction(metadata map[string]string) error
	RollbackTransaction(metadata map[string]string) error
	CloseTransaction() error
	Stream(requestId []byte, metadata map[string]string) (*typedb_protocol.Transaction_Res, error)
}

type transactionManagerImpl struct {
	TransactionClient typedb_protocol.TypeDB_TransactionClient
	Context           context.Context
}

func NewTransactionManager(transactionClient typedb_protocol.TypeDB_TransactionClient, ctx context.Context) *transactionManagerImpl {
	return &transactionManagerImpl{
		TransactionClient: transactionClient,
		Context: ctx,
	}
}

func (q *transactionManagerImpl) OpenTransaction(sessionId []byte, transactionType typedb_protocol.Transaction_Type, options *typedb_protocol.Options, latencyMillis int32, metadata map[string]string) error {
	request := transaction.OpenReq(sessionId, transactionType, options, latencyMillis)
	request.Metadata = metadata

	response, err := common.NewPromise(q.TransactionClient, q.Context, request)
	if err != nil {
		return err
	}

	for {
		select {
		case <-response.Ctx.Done():
			return response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return response.Err
			}

			data.GetOpenRes()

			return nil
		}
	}
}

func (q *transactionManagerImpl) CommitTransaction(metadata map[string]string) error {
	request := transaction.CommitReq()
	request.Metadata = metadata

	response, err := common.NewPromise(q.TransactionClient, q.Context, request)
	if err != nil {
		return err
	}

	for {
		select {
		case <-response.Ctx.Done():
			return response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return response.Err
			}

			data.GetCommitRes()

			return nil
		}
	}
}

func (q *transactionManagerImpl) RollbackTransaction(metadata map[string]string) error {
	request := transaction.RollbackReq()
	request.Metadata = metadata

	response, err := common.NewPromise(q.TransactionClient, q.Context, request)
	if err != nil {
		return err
	}

	for {
		select {
		case <-response.Ctx.Done():
			return response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return response.Err
			}

			data.GetRollbackRes()

			return nil
		}
	}
}

func (q *transactionManagerImpl) CloseTransaction() error {
	return q.TransactionClient.CloseSend()
}

func (q *transactionManagerImpl) Stream(requestId []byte, metadata map[string]string) (*typedb_protocol.Transaction_Res, error) {
	request := transaction.StreamReq(requestId)
	request.Metadata = metadata

	response, err := common.NewPromise(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	var result *typedb_protocol.Transaction_Res = nil

	for {
		select {
		case <-response.Ctx.Done():
			return result, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}

			result = data
		}
	}
}
