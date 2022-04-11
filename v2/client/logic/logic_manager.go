package logic

import (
	"context"
	"github.com/taliesins/typedb-client-go/v2/client/common"
	"github.com/taliesins/typedb-client-go/v2/client/common/rpc/request_builder/logic_manager"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

type LogicManager interface {
	GetRule(label string) (*typedb_protocol.Rule, error)
	GetRules()([]*typedb_protocol.Rule, error)
	PutRule(label string, when string, then string) (*typedb_protocol.Rule, error)
}

type logicManagerImpl struct {
	TransactionClient typedb_protocol.TypeDB_TransactionClient
	Context           context.Context
}

func NewLogicManager(transactionClient typedb_protocol.TypeDB_TransactionClient, ctx context.Context) *logicManagerImpl {
	return &logicManagerImpl{
		TransactionClient: transactionClient,
		Context: ctx,
	}
}

func (q *logicManagerImpl) GetRule(label string, metadata map[string]string) (*typedb_protocol.Rule, error) { //Promise<Rule> {
	request := logic_manager.GetRuleReq(label)
	request.Metadata = metadata

	response, err := common.NewPromise(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	var rule *typedb_protocol.Rule = nil

	for {
		select {
		case <-response.Ctx.Done():
			return rule, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}

			rule = data.GetLogicManagerRes().GetGetRuleRes().GetRule()
		}
	}
}

func (q *logicManagerImpl) GetRules(metadata map[string]string)([]*typedb_protocol.Rule, error) { //Stream<Rule> {
	request := logic_manager.GetRulesReq()
	request.Metadata = metadata

	response, err := common.NewStream(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	rules := make([]*typedb_protocol.Rule, 0)
	for {
		select {
		case <-response.Ctx.Done():
			return rules, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}
			rules = append(rules, data.GetLogicManagerResPart().GetGetRulesResPart().GetRules()...)
		}
	}
}

func (q *logicManagerImpl) PutRule(label string, when string, then string, metadata map[string]string)(*typedb_protocol.Rule, error) { //Stream<Rule> {
	request := logic_manager.PutRuleReq(label, when, then)
	request.Metadata = metadata

	response, err := common.NewPromise(q.TransactionClient, q.Context, request)
	if err != nil {
		return nil, err
	}

	var rule *typedb_protocol.Rule = nil

	for {
		select {
		case <-response.Ctx.Done():
			return rule, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}

			rule = data.GetLogicManagerRes().GetPutRuleRes().GetRule()
		}
	}
}