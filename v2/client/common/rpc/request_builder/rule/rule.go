package rule

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func SetLabelReq(currentLabel string, newLabel string) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_RuleReq{
			RuleReq: &typedb_protocol.Rule_Req{
				Req: &typedb_protocol.Rule_Req_RuleSetLabelReq{
					RuleSetLabelReq: &typedb_protocol.Rule_SetLabel_Req{
						Label: newLabel,
					},
				},
				Label: currentLabel,
			},
		},
	}
}

func DeleteLabelReq(label string) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_RuleReq{
			RuleReq: &typedb_protocol.Rule_Req{
				Req: &typedb_protocol.Rule_Req_RuleDeleteReq{
					RuleDeleteReq: &typedb_protocol.Rule_Delete_Req{
					},
				},
				Label: label,
			},
		},
	}
}
