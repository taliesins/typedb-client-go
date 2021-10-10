package entity_type

import (
	"github.com/taliesins/typedb-client-go/v2/client/common"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

func newTransactionRequest(req *typedb_protocol.Transaction_Req_TypeReq) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: req,
	}
}

func newTypeRequest(label common.Label) *typedb_protocol.Type_Req {
	builder := &typedb_protocol.Type_Req{
		Label: label.Name,
		Scope: label.Scope,
	}
	return builder
}

func setEntityTypeCreateReq(typeReq *typedb_protocol.Type_Req, entityTypeCreateReq *typedb_protocol.EntityType_Create_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_EntityTypeCreateReq{
		EntityTypeCreateReq: entityTypeCreateReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func CreateReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setEntityTypeCreateReq(newTypeRequest(label),
		&typedb_protocol.EntityType_Create_Req{}))
}
