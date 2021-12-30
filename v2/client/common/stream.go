package common

import (
	"context"
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/taliesins/typedb-client-go/v2/client/common/rpc/request_builder/transaction"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

func NewStream(transactionClient typedb_protocol.TypeDB_TransactionClient, ctx context.Context, req *typedb_protocol.Transaction_Req) (*stream, error) { //Stream<QueryProto.ResPart> {
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

	c := &stream{
		Ctx:               ctx,
		ReqId:             req.ReqId,
		TransactionClient: transactionClient,
		Data:              make(chan *typedb_protocol.Transaction_ResPart),
	}
	go c.begin()

	return c, nil
}

type stream struct {
	Ctx               context.Context
	ReqId             []byte
	TransactionClient typedb_protocol.TypeDB_TransactionClient
	Data chan *typedb_protocol.Transaction_ResPart
	Err  error
}

func (c *stream) begin() {
	for {
		transactionServer, err := c.TransactionClient.Recv()
		if err != nil {
			c.Err = fmt.Errorf("could not receive query response: %w", err)
			close(c.Data)
			return
		}

		transactionResponsePart := transactionServer.GetResPart()
		if transactionResponsePart == nil {
			transactionResponse := transactionServer.GetRes()
			if transactionResponse == nil {
				c.Err = fmt.Errorf("Unknown response type")
				close(c.Data)
				return
			} else {
				c.Err = fmt.Errorf("single response type")
				close(c.Data)
				return
			}
		} else {
			streamResPart := transactionResponsePart.GetStreamResPart()
			if streamResPart == nil {
				c.Data <- transactionResponsePart
			} else {
				state := streamResPart.GetState()
				switch state {
				case typedb_protocol.Transaction_Stream_DONE:
					close(c.Data)
					return
				case typedb_protocol.Transaction_Stream_CONTINUE:
					err = c.TransactionClient.Send(
						&typedb_protocol.Transaction_Client{
							Reqs: []*typedb_protocol.Transaction_Req{
								transaction.StreamReq(c.ReqId),
							},
						},
					)

					if err != nil {
						c.Err = err
						close(c.Data)
						return
					}
				default:
					c.Err = fmt.Errorf("unknown stream state")
					close(c.Data)
					return
				}
			}
		}
	}

	c.Err = fmt.Errorf("Exited event loop: %w", c.Err)
	close(c.Data)
	return
}


