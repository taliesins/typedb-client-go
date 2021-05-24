package role_type

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

func protoRoleType(label common.Label, encoding grakn_protocol.Type_Encoding) *grakn_protocol.Type {
	return &grakn_protocol.Type{
		Scope: label.Scope,
		Label: label.Name,
		Encoding: encoding,
	}
}

func setRoleTypeGetRelationTypesReq(typeReq *grakn_protocol.Type_Req, roleTypeGetRelationTypesReq *grakn_protocol.RoleType_GetRelationTypes_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_RoleTypeGetRelationTypesReq{
		RoleTypeGetRelationTypesReq: roleTypeGetRelationTypesReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetRelationTypesReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRoleTypeGetRelationTypesReq(newTypeRequest(label), &grakn_protocol.RoleType_GetRelationTypes_Req{}))
}

func setRoleTypeGetPlayersReq(typeReq *grakn_protocol.Type_Req, roleTypeGetPlayersReq *grakn_protocol.RoleType_GetPlayers_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_RoleTypeGetPlayersReq{
		RoleTypeGetPlayersReq: roleTypeGetPlayersReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetPlayersReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRoleTypeGetPlayersReq(newTypeRequest(label), &grakn_protocol.RoleType_GetPlayers_Req{}))
}
