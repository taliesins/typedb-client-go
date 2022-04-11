package common

import (
	"context"
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

func NewPromise(transactionClient typedb_protocol.TypeDB_TransactionClient, ctx context.Context, req *typedb_protocol.Transaction_Req) (*promise, error) {
	requestId := ksuid.New().String()
	req.ReqId = []byte(requestId)

	err := transactionClient.Send(
		&typedb_protocol.Transaction_Client{
			Reqs: []*typedb_protocol.Transaction_Req{
				req,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	c := &promise{
		Ctx:               ctx,
		ReqId:             req.ReqId,
		TransactionClient: transactionClient,
		Data:              make(chan *typedb_protocol.Transaction_Res),
	}
	go c.begin()

	return c, nil
}

type promise struct {
	Ctx               context.Context
	ReqId             []byte
	TransactionClient typedb_protocol.TypeDB_TransactionClient
	Data chan *typedb_protocol.Transaction_Res
	Err  error
}

func (c *promise) begin() {
	transactionServer, err := c.TransactionClient.Recv()
	if err != nil {
		c.Err = fmt.Errorf("could not receive query response: %w", err)
		close(c.Data)
		return
	}

	transactionResponse := transactionServer.GetRes()
	c.Data <- transactionResponse
	close(c.Data)
	return
}



