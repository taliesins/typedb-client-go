package entity_type

import (
	"github.com/taliesins/typedb-client-go/v2/client/common"
	"github.com/taliesins/typedb-client-go/v2/grakn_protocol"
)

func newTransactionRequest(req *grakn_protocol.Transaction_Req_TypeReq) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: req,
	}
}

func newTypeRequest(label common.Label) *grakn_protocol.Type_Req {
	builder := &grakn_protocol.Type_Req{
		Label: label.Name,
		Scope: label.Scope,
	}
	return builder
}

func setEntityTypeCreateReq(typeReq *grakn_protocol.Type_Req, entityTypeCreateReq *grakn_protocol.EntityType_Create_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_EntityTypeCreateReq{
		EntityTypeCreateReq: entityTypeCreateReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func CreateReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setEntityTypeCreateReq(newTypeRequest(label),
		&grakn_protocol.EntityType_Create_Req{}))
}
