package logic_manager

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func PutRuleReq(label string, when string, then string) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_LogicManagerReq{
			LogicManagerReq: &grakn_protocol.LogicManager_Req{
				Req: &grakn_protocol.LogicManager_Req_PutRuleReq{
					PutRuleReq: &grakn_protocol.LogicManager_PutRule_Req{
						Label: label,
						When:  when,
						Then:  then,
					},
				},
			},
		},
	}
}

func GetRuleReq(label string) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_LogicManagerReq{
			LogicManagerReq: &grakn_protocol.LogicManager_Req{
				Req: &grakn_protocol.LogicManager_Req_GetRuleReq{
					GetRuleReq: &grakn_protocol.LogicManager_GetRule_Req{
						Label: label,
					},
				},
			},
		},
	}
}

func GetRulesReq() *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_LogicManagerReq{
			LogicManagerReq: &grakn_protocol.LogicManager_Req{
				Req: &grakn_protocol.LogicManager_Req_GetRulesReq{
					GetRulesReq: &grakn_protocol.LogicManager_GetRules_Req{
					},
				},
			},
		},
	}
}
