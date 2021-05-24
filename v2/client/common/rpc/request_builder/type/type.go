package _type

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

func setTypeIsAbstractReq(typeReq *grakn_protocol.Type_Req, typeIsAbstractReq *grakn_protocol.Type_IsAbstract_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_TypeIsAbstractReq{
		TypeIsAbstractReq: typeIsAbstractReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func IsAbstractReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setTypeIsAbstractReq(newTypeRequest(label), &grakn_protocol.Type_IsAbstract_Req{}))
}

func setTypeSetLabelReq(typeReq *grakn_protocol.Type_Req, typeSetLabelReq *grakn_protocol.Type_SetLabel_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_TypeSetLabelReq{
		TypeSetLabelReq: typeSetLabelReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetLabelReq(label common.Label, newLabel string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setTypeSetLabelReq(newTypeRequest(label),
		&grakn_protocol.Type_SetLabel_Req{
			Label: newLabel,
		},
	))
}

func getTypeGetSupertypesReq(typeReq *grakn_protocol.Type_Req, typeGetSupertypeReq *grakn_protocol.Type_GetSupertype_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_TypeGetSupertypeReq{
		TypeGetSupertypeReq: typeGetSupertypeReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetSupertypesReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(getTypeGetSupertypesReq(newTypeRequest(label), &grakn_protocol.Type_GetSupertype_Req{}))
}

func setTypeGetSubtypesReq(typeReq *grakn_protocol.Type_Req, typeGetSubtypesReq *grakn_protocol.Type_GetSubtypes_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_TypeGetSubtypesReq{
		TypeGetSubtypesReq: typeGetSubtypesReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetSubtypesReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setTypeGetSubtypesReq(newTypeRequest(label), &grakn_protocol.Type_GetSubtypes_Req{}))
}

func setTypeGetSupertypeReq(typeReq *grakn_protocol.Type_Req, typeSetSupertypeReq *grakn_protocol.Type_SetSupertype_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_TypeSetSupertypeReq{
		TypeSetSupertypeReq: typeSetSupertypeReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetSupertypeReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setTypeGetSupertypeReq(newTypeRequest(label), &grakn_protocol.Type_SetSupertype_Req{}))
}

func setTypeDeleteReq(typeReq *grakn_protocol.Type_Req, typeDeleteReq *grakn_protocol.Type_Delete_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_TypeDeleteReq{
		TypeDeleteReq: typeDeleteReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func DeleteReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setTypeDeleteReq(newTypeRequest(label), &grakn_protocol.Type_Delete_Req{}))
}