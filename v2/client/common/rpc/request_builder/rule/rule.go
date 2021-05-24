package rule

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func SetLabelReq(currentLabel string, newLabel string) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_RuleReq{
			RuleReq: &grakn_protocol.Rule_Req{
				Req: &grakn_protocol.Rule_Req_RuleSetLabelReq{
					RuleSetLabelReq: &grakn_protocol.Rule_SetLabel_Req{
						Label: newLabel,
					},
				},
				Label: currentLabel,
			},
		},
	}
}

func DeleteLabelReq(label string) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_RuleReq{
			RuleReq: &grakn_protocol.Rule_Req{
				Req: &grakn_protocol.Rule_Req_RuleDeleteReq{
					RuleDeleteReq: &grakn_protocol.Rule_Delete_Req{
					},
				},
				Label: label,
			},
		},
	}
}
