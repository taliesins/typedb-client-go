package role_type

import (
	"github.com/taliesins/typedb-client-go/v2/client/common"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

func ProtoRoleType(label common.Label, encoding typedb_protocol.Type_Encoding) *typedb_protocol.Type {
	return &typedb_protocol.Type{
		Scope: label.Scope,
		Label: label.Name,
		Encoding: encoding,
	}
}

func GetRelationTypesReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRoleTypeGetRelationTypesReq(newTypeRequest(label), &typedb_protocol.RoleType_GetRelationTypes_Req{}))
}

func GetPlayersReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRoleTypeGetPlayersReq(newTypeRequest(label), &typedb_protocol.RoleType_GetPlayers_Req{}))
}

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

func setRoleTypeGetRelationTypesReq(typeReq *typedb_protocol.Type_Req, roleTypeGetRelationTypesReq *typedb_protocol.RoleType_GetRelationTypes_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_RoleTypeGetRelationTypesReq{
		RoleTypeGetRelationTypesReq: roleTypeGetRelationTypesReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func setRoleTypeGetPlayersReq(typeReq *typedb_protocol.Type_Req, roleTypeGetPlayersReq *typedb_protocol.RoleType_GetPlayers_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_RoleTypeGetPlayersReq{
		RoleTypeGetPlayersReq: roleTypeGetPlayersReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}
