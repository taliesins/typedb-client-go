package logic_manager

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func PutRuleReq(label string, when string, then string) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_LogicManagerReq{
			LogicManagerReq: &typedb_protocol.LogicManager_Req{
				Req: &typedb_protocol.LogicManager_Req_PutRuleReq{
					PutRuleReq: &typedb_protocol.LogicManager_PutRule_Req{
						Label: label,
						When:  when,
						Then:  then,
					},
				},
			},
		},
	}
}

func GetRuleReq(label string) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_LogicManagerReq{
			LogicManagerReq: &typedb_protocol.LogicManager_Req{
				Req: &typedb_protocol.LogicManager_Req_GetRuleReq{
					GetRuleReq: &typedb_protocol.LogicManager_GetRule_Req{
						Label: label,
					},
				},
			},
		},
	}
}

func GetRulesReq() *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_LogicManagerReq{
			LogicManagerReq: &typedb_protocol.LogicManager_Req{
				Req: &typedb_protocol.LogicManager_Req_GetRulesReq{
					GetRulesReq: &typedb_protocol.LogicManager_GetRules_Req{
					},
				},
			},
		},
	}
}
