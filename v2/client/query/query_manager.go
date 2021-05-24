package query

import (
	"context"
	"fmt"
	"github.com/segmentio/ksuid"
	"github.com/taliesins/typedb-client-go/v2/client/common/rpc/request_builder/query_manager"
	"github.com/taliesins/typedb-client-go/v2/client/common/rpc/request_builder/transaction"
	grakn "github.com/taliesins/typedb-client-go/v2/grakn_protocol"
)

type QueryManager interface {
	Match(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.ConceptMap, error)
	MatchAggregate(query string, options *grakn.Options, metadata map[string]string) (*grakn.Numeric, error)
	MatchGroup(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.ConceptMapGroup, error)
	MatchGroupAggregate(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.NumericGroup, error)
	Insert(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.ConceptMap, error)
	Delete(query string, options *grakn.Options, metadata map[string]string) error
	Update(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.ConceptMap, error)
	Define(query string, options *grakn.Options, metadata map[string]string) error
	Undefine(query string, options *grakn.Options, metadata map[string]string) error
	//Explain(explainable ConceptMap.Explainable, options *grakn.Options, metadata map[string]string)
}

type QueryManagerAsync interface {
	MatchAsync(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.ConceptMap; Error error}, error)
	MatchGroup(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.ConceptMapGroup; Error error}, error)
	MatchGroupAggregate(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.NumericGroup; Error error}, error)
	InsertAsync(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.ConceptMap; Error error}, error)
	UpdateAsync(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.ConceptMap; Error error}, error)
}

type queryManagerImpl struct {
	TransactionClient grakn.GraknCore_TransactionClient
	Context context.Context
}

func NewQueryManager(transactionClient grakn.GraknCore_TransactionClient, ctx context.Context) *queryManagerImpl {
	return &queryManagerImpl{
		TransactionClient: transactionClient,
		Context: ctx,
	}
}

func (q *queryManagerImpl) Match(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.ConceptMap, error) {//Stream<ConceptMap> {
	request := query_manager.MatchReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make([]*grakn.ConceptMap, 0)
	for {
		select {
		case <-response.ctx.Done():
			return answers, response.ctx.Err()
		case data, ok := <-response.data:
			if !ok {
				return nil, response.err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetMatchResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) MatchAsync(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.ConceptMap; Error error}, error) { //Stream<ConceptMap> {
	request := query_manager.MatchReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*grakn.ConceptMap; Error error})
	go func() {
		for {
			select {
			case <-response.ctx.Done():
				err := response.ctx.Err()
				if err != nil {
					answers <- struct {Answers []*grakn.ConceptMap; Error error}{nil, response.ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.data:
				if !ok {
					err := response.err
					if err != nil {
						answers <- struct {Answers []*grakn.ConceptMap; Error error}{nil, response.err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*grakn.ConceptMap; Error error}{data.GetQueryManagerResPart().GetMatchResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) MatchAggregate(query string, options *grakn.Options, metadata map[string]string) (*grakn.Numeric, error) {//Promise<Numeric>{
	request := query_manager.MatchReq(query, options)
	request.Metadata = metadata
	response, err := q.query(request)
	if err != nil {
		return nil, err
	}
	return response.GetQueryManagerRes().GetMatchAggregateRes().GetAnswer(), nil
}

func (q *queryManagerImpl) MatchGroup(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.ConceptMapGroup, error) {//Stream<ConceptMapGroup>{
	request := query_manager.MatchGroupReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make([]*grakn.ConceptMapGroup, 0)
	for {
		select {
		case <-response.ctx.Done():
			return answers, response.ctx.Err()
		case data, ok := <-response.data:
			if !ok {
				return nil, response.err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetMatchGroupResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) MatchGroupAsync(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.ConceptMapGroup; Error error}, error) { //Stream<ConceptMapGroup>{
	request := query_manager.MatchGroupReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*grakn.ConceptMapGroup; Error error})
	go func() {
		for {
			select {
			case <-response.ctx.Done():
				err := response.ctx.Err()
				if err != nil {
					answers <- struct {Answers []*grakn.ConceptMapGroup; Error error}{nil, response.ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.data:
				if !ok {
					err := response.err
					if err != nil {
						answers <- struct {Answers []*grakn.ConceptMapGroup; Error error}{nil, response.err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*grakn.ConceptMapGroup; Error error}{data.GetQueryManagerResPart().GetMatchGroupResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) MatchGroupAggregate(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.NumericGroup, error) {//Stream<NumericGroup>{
	request := query_manager.MatchGroupAggregateReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context,request)
	if err != nil {
		return nil, err
	}

	answers := make([]*grakn.NumericGroup, 0)
	for {
		select {
		case <-response.ctx.Done():
			return answers, response.ctx.Err()
		case data, ok := <-response.data:
			if !ok {
				return nil, response.err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetMatchGroupAggregateResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) MatchGroupAggregateAsync(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.NumericGroup; Error error}, error) {//Stream<NumericGroup>{
	request := query_manager.MatchGroupAggregateReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*grakn.NumericGroup; Error error})
	go func() {
		for {
			select {
			case <-response.ctx.Done():
				err := response.ctx.Err()
				if err != nil {
					answers <- struct {Answers []*grakn.NumericGroup; Error error}{nil, response.ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.data:
				if !ok {
					err := response.err
					if err != nil {
						answers <- struct {Answers []*grakn.NumericGroup; Error error}{nil, response.err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*grakn.NumericGroup; Error error}{data.GetQueryManagerResPart().GetMatchGroupAggregateResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) Insert(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.ConceptMap, error) {//Stream<ConceptMap>{
	request := query_manager.InsertReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context,request)
	if err != nil {
		return nil, err
	}

	answers := make([]*grakn.ConceptMap, 0)
	for {
		select {
		case <-response.ctx.Done():
			return answers, response.ctx.Err()
		case data, ok := <-response.data:
			if !ok {
				return nil, response.err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetInsertResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) InsertAsync(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.ConceptMap; Error error}, error) {//Stream<ConceptMap>{
	request := query_manager.InsertReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*grakn.ConceptMap; Error error})
	go func() {
		for {
			select {
			case <-response.ctx.Done():
				err := response.ctx.Err()
				if err != nil {
					answers <- struct {Answers []*grakn.ConceptMap; Error error}{nil, response.ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.data:
				if !ok {
					err := response.err
					if err != nil {
						answers <- struct {Answers []*grakn.ConceptMap; Error error}{nil, response.err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*grakn.ConceptMap; Error error}{data.GetQueryManagerResPart().GetInsertResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) Delete(query string, options *grakn.Options, metadata map[string]string) error { //Promise<void>{
	request := query_manager.DeleteReq(query, options)
	request.Metadata = metadata
	response, err := q.query(request)
	if err != nil {
		return err
	}

	response.GetQueryManagerRes().GetDeleteRes()

	return nil
}

func (q *queryManagerImpl) Update(query string, options *grakn.Options, metadata map[string]string) ([]*grakn.ConceptMap, error) { //Stream<ConceptMap>{
	request := query_manager.UpdateReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make([]*grakn.ConceptMap, 0)
	for {
		select {
		case <-response.ctx.Done():
			return answers, response.ctx.Err()
		case data, ok := <-response.data:
			if !ok {
				return nil, response.err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetUpdateResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) UpdateAsync(query string, options *grakn.Options, metadata map[string]string) (chan struct {Answers []*grakn.ConceptMap; Error error}, error) {//Stream<ConceptMap>{
	request := query_manager.UpdateReq(query, options)
	request.Metadata = metadata
	response, err := q.stream(q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*grakn.ConceptMap; Error error})
	go func() {
		for {
			select {
			case <-response.ctx.Done():
				err := response.ctx.Err()
				if err != nil {
					answers <- struct {Answers []*grakn.ConceptMap; Error error}{nil, response.ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.data:
				if !ok {
					err := response.err
					if err != nil {
						answers <- struct {Answers []*grakn.ConceptMap; Error error}{nil, response.err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*grakn.ConceptMap; Error error}{data.GetQueryManagerResPart().GetUpdateResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) Define(query string, options *grakn.Options, metadata map[string]string) error { //Promise<void>{
	request := query_manager.DefineReq(query, options)
	request.Metadata = metadata
	response, err := q.query(request)
	if err != nil {
		return err
	}

	response.GetQueryManagerRes().GetDefineRes()

	return nil
}

func (q *queryManagerImpl) Undefine(query string, options *grakn.Options, metadata map[string]string) error {//Promise<void>{
	request := query_manager.UndefineReq(query, options)
	request.Metadata = metadata
	response, err := q.query(request)
	if err != nil {
		return err
	}
	response.GetQueryManagerRes().GetUndefineRes()

	return nil
}

/*
func (q *queryManagerImpl) Explain(explainable ConceptMap.Explainable, options *grakn.Options) Stream<Explanation>{

}

 */

func (q *queryManagerImpl) query(req *grakn.Transaction_Req) (res *grakn.Transaction_Res, err error) { //Promise<QueryProto.Res> {
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

func (q *queryManagerImpl) stream(ctx context.Context, req *grakn.Transaction_Req) (*cancelableStream, error) { //Stream<QueryProto.ResPart> {
	requestId := ksuid.New().String()
	req.ReqId = []byte(requestId)

	err := q.TransactionClient.Send(
		&grakn.Transaction_Client{
			Reqs: []*grakn.Transaction_Req{
				req,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return newCancelableStream(ctx, req.ReqId, q.TransactionClient), err
}

type cancelableStream struct {
	ctx  context.Context
	reqId []byte
	transactionClient grakn.GraknCore_TransactionClient
	data chan *grakn.Transaction_ResPart
	err  error
}

func (c *cancelableStream) begin() {
	for {
		transactionServer, err := c.transactionClient.Recv()
		if err != nil {
			c.err = fmt.Errorf("could not receive query response: %w", err)
			close(c.data)
			return
		}

		transactionResponse := transactionServer.GetResPart()
		c.data <- transactionResponse

		switch transactionResponse.GetStreamResPart().GetState() {
		case grakn.Transaction_Stream_DONE:
			close(c.data)
			return
		case grakn.Transaction_Stream_CONTINUE:
			err = c.transactionClient.Send(
				&grakn.Transaction_Client{
					Reqs: []*grakn.Transaction_Req{
						transaction.StreamReq(c.reqId),
					},
				},
			)
			if err != nil {
				c.err = err
				close(c.data)
				return
			}
		default:
			c.err = fmt.Errorf("unknown stream state: %s", transactionServer.GetResPart().GetStreamResPart().GetState())
			close(c.data)
			return
		}
	}
}

func newCancelableStream(ctx context.Context, reqId []byte, transactionClient grakn.GraknCore_TransactionClient) *cancelableStream {
	c := &cancelableStream{
		ctx:  ctx,
		reqId: reqId,
		transactionClient: transactionClient,
		data: make(chan *grakn.Transaction_ResPart),
	}
	go c.begin()
	return c
}



