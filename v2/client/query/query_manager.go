package query

import (
	"context"
	"github.com/taliesins/typedb-client-go/v2/client/common"
	"github.com/taliesins/typedb-client-go/v2/client/common/rpc/request_builder/query_manager"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

type QueryManager interface {
	Match(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.ConceptMap, error)
	MatchAggregate(query string, options *typedb_protocol.Options, metadata map[string]string) (*typedb_protocol.Numeric, error)
	MatchGroup(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.ConceptMapGroup, error)
	MatchGroupAggregate(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.NumericGroup, error)
	Insert(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.ConceptMap, error)
	Delete(query string, options *typedb_protocol.Options, metadata map[string]string) error
	Update(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.ConceptMap, error)
	Define(query string, options *typedb_protocol.Options, metadata map[string]string) error
	Undefine(query string, options *typedb_protocol.Options, metadata map[string]string) error
	Explain(explainable *typedb_protocol.Explainable, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.Explanation, error)
}

type QueryManagerAsync interface {
	MatchAsync(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.ConceptMap; Error error}, error)
	MatchGroup(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.ConceptMapGroup; Error error}, error)
	MatchGroupAggregate(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.NumericGroup; Error error}, error)
	InsertAsync(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.ConceptMap; Error error}, error)
	UpdateAsync(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.ConceptMap; Error error}, error)
}

type queryManagerImpl struct {
	TransactionClient typedb_protocol.TypeDB_TransactionClient
	Context           context.Context
	Metadata map[string]string
}

func NewQueryManager(transactionClient typedb_protocol.TypeDB_TransactionClient, ctx context.Context) *queryManagerImpl {
	return &queryManagerImpl{
		TransactionClient: transactionClient,
		Context: ctx,
	}
}

func (q *queryManagerImpl) Match(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.ConceptMap, error) { //Stream<ConceptMap> {
	request := query_manager.MatchReq(query, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make([]*typedb_protocol.ConceptMap, 0)
	for {
		select {
		case <-response.Ctx.Done():
			return answers, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetMatchResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) MatchAsync(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.ConceptMap; Error error}, error) { //Stream<ConceptMap> {
	request := query_manager.MatchReq(query, options)
	request.Metadata = metadata

	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*typedb_protocol.ConceptMap; Error error})
	go func() {
		for {
			select {
			case <-response.Ctx.Done():
				err := response.Ctx.Err()
				if err != nil {
					answers <- struct {Answers []*typedb_protocol.ConceptMap; Error error}{nil, response.Ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.Data:
				if !ok {
					err := response.Err
					if err != nil {
						answers <- struct {Answers []*typedb_protocol.ConceptMap; Error error}{nil, response.Err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*typedb_protocol.ConceptMap; Error error}{data.GetQueryManagerResPart().GetMatchResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) MatchAggregate(query string, options *typedb_protocol.Options, metadata map[string]string) (*typedb_protocol.Numeric, error) { //Promise<Numeric>{
	request := query_manager.MatchReq(query, options)
	request.Metadata = metadata
	response, err := common.NewPromise(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	var attributeType *typedb_protocol.Numeric = nil

	for {
		select {
		case <-response.Ctx.Done():
			return attributeType, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}

			attributeType = data.GetQueryManagerRes().GetMatchAggregateRes().GetAnswer()
		}
	}
}

func (q *queryManagerImpl) MatchGroup(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.ConceptMapGroup, error) { //Stream<ConceptMapGroup>{
	request := query_manager.MatchGroupReq(query, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make([]*typedb_protocol.ConceptMapGroup, 0)
	for {
		select {
		case <-response.Ctx.Done():
			return answers, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetMatchGroupResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) MatchGroupAsync(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.ConceptMapGroup; Error error}, error) { //Stream<ConceptMapGroup>{
	request := query_manager.MatchGroupReq(query, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*typedb_protocol.ConceptMapGroup; Error error})
	go func() {
		for {
			select {
			case <-response.Ctx.Done():
				err := response.Ctx.Err()
				if err != nil {
					answers <- struct {Answers []*typedb_protocol.ConceptMapGroup; Error error}{nil, response.Ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.Data:
				if !ok {
					err := response.Err
					if err != nil {
						answers <- struct {Answers []*typedb_protocol.ConceptMapGroup; Error error}{nil, response.Err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*typedb_protocol.ConceptMapGroup; Error error}{data.GetQueryManagerResPart().GetMatchGroupResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) MatchGroupAggregate(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.NumericGroup, error) { //Stream<NumericGroup>{
	request := query_manager.MatchGroupAggregateReq(query, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context,request)
	if err != nil {
		return nil, err
	}

	answers := make([]*typedb_protocol.NumericGroup, 0)
	for {
		select {
		case <-response.Ctx.Done():
			return answers, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetMatchGroupAggregateResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) MatchGroupAggregateAsync(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.NumericGroup; Error error}, error) { //Stream<NumericGroup>{
	request := query_manager.MatchGroupAggregateReq(query, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*typedb_protocol.NumericGroup; Error error})
	go func() {
		for {
			select {
			case <-response.Ctx.Done():
				err := response.Ctx.Err()
				if err != nil {
					answers <- struct {Answers []*typedb_protocol.NumericGroup; Error error}{nil, response.Ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.Data:
				if !ok {
					err := response.Err
					if err != nil {
						answers <- struct {Answers []*typedb_protocol.NumericGroup; Error error}{nil, response.Err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*typedb_protocol.NumericGroup; Error error}{data.GetQueryManagerResPart().GetMatchGroupAggregateResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) Insert(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.ConceptMap, error) { //Stream<ConceptMap>{
	request := query_manager.InsertReq(query, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context,request)
	if err != nil {
		return nil, err
	}

	answers := make([]*typedb_protocol.ConceptMap, 0)
	for {
		select {
		case <-response.Ctx.Done():
			return answers, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetInsertResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) InsertAsync(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.ConceptMap; Error error}, error) { //Stream<ConceptMap>{
	request := query_manager.InsertReq(query, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*typedb_protocol.ConceptMap; Error error})
	go func() {
		for {
			select {
			case <-response.Ctx.Done():
				err := response.Ctx.Err()
				if err != nil {
					answers <- struct {Answers []*typedb_protocol.ConceptMap; Error error}{nil, response.Ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.Data:
				if !ok {
					err := response.Err
					if err != nil {
						answers <- struct {Answers []*typedb_protocol.ConceptMap; Error error}{nil, response.Err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*typedb_protocol.ConceptMap; Error error}{data.GetQueryManagerResPart().GetInsertResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) Delete(query string, options *typedb_protocol.Options, metadata map[string]string) error { //Promise<void>{
	request := query_manager.DeleteReq(query, options)
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

			data.GetQueryManagerRes().GetDeleteRes()
		}
	}
}

func (q *queryManagerImpl) Update(query string, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.ConceptMap, error) { //Stream<ConceptMap>{
	request := query_manager.UpdateReq(query, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make([]*typedb_protocol.ConceptMap, 0)
	for {
		select {
		case <-response.Ctx.Done():
			return answers, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}
			answers = append(answers, data.GetQueryManagerResPart().GetUpdateResPart().GetAnswers()...)
		}
	}
}

func (q *queryManagerImpl) UpdateAsync(query string, options *typedb_protocol.Options, metadata map[string]string) (chan struct {Answers []*typedb_protocol.ConceptMap; Error error}, error) { //Stream<ConceptMap>{
	request := query_manager.UpdateReq(query, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	answers := make(chan struct {Answers []*typedb_protocol.ConceptMap; Error error})
	go func() {
		for {
			select {
			case <-response.Ctx.Done():
				err := response.Ctx.Err()
				if err != nil {
					answers <- struct {Answers []*typedb_protocol.ConceptMap; Error error}{nil, response.Ctx.Err()}
				}
				close(answers)
				return
			case data, ok := <-response.Data:
				if !ok {
					err := response.Err
					if err != nil {
						answers <- struct {Answers []*typedb_protocol.ConceptMap; Error error}{nil, response.Err}
					}
					close(answers)
					return
				}
				answers <- struct {Answers []*typedb_protocol.ConceptMap; Error error}{data.GetQueryManagerResPart().GetUpdateResPart().GetAnswers(), nil}
			}
		}
	}()

	return answers, nil
}

func (q *queryManagerImpl) Define(query string, options *typedb_protocol.Options, metadata map[string]string) error { //Promise<void>{
	request := query_manager.DefineReq(query, options)
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

			data.GetQueryManagerRes().GetDefineRes()
		}
	}
}

func (q *queryManagerImpl) Undefine(query string, options *typedb_protocol.Options, metadata map[string]string) error { //Promise<void>{
	request := query_manager.UndefineReq(query, options)
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

			data.GetQueryManagerRes().GetUndefineRes()
		}
	}
}

func (q *queryManagerImpl) Explain(explainable *typedb_protocol.Explainable, options *typedb_protocol.Options, metadata map[string]string) ([]*typedb_protocol.Explanation, error) { //Stream<Explanation>{
	request := query_manager.ExplainReq(explainable.Id, options)
	request.Metadata = metadata
	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	explanations := make([]*typedb_protocol.Explanation, 0)
	for {
		select {
		case <-response.Ctx.Done():
			return explanations, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}
			explanations = append(explanations, data.GetQueryManagerResPart().GetExplainResPart().GetExplanations()...)
		}
	}
}
