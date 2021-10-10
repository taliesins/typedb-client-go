package _type

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

func setTypeIsAbstractReq(typeReq *typedb_protocol.Type_Req, typeIsAbstractReq *typedb_protocol.Type_IsAbstract_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_TypeIsAbstractReq{
		TypeIsAbstractReq: typeIsAbstractReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func IsAbstractReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setTypeIsAbstractReq(newTypeRequest(label), &typedb_protocol.Type_IsAbstract_Req{}))
}

func setTypeSetLabelReq(typeReq *typedb_protocol.Type_Req, typeSetLabelReq *typedb_protocol.Type_SetLabel_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_TypeSetLabelReq{
		TypeSetLabelReq: typeSetLabelReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetLabelReq(label common.Label, newLabel string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setTypeSetLabelReq(newTypeRequest(label),
		&typedb_protocol.Type_SetLabel_Req{
			Label: newLabel,
		},
	))
}

func getTypeGetSupertypesReq(typeReq *typedb_protocol.Type_Req, typeGetSupertypeReq *typedb_protocol.Type_GetSupertype_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_TypeGetSupertypeReq{
		TypeGetSupertypeReq: typeGetSupertypeReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetSupertypesReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(getTypeGetSupertypesReq(newTypeRequest(label), &typedb_protocol.Type_GetSupertype_Req{}))
}

func setTypeGetSubtypesReq(typeReq *typedb_protocol.Type_Req, typeGetSubtypesReq *typedb_protocol.Type_GetSubtypes_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_TypeGetSubtypesReq{
		TypeGetSubtypesReq: typeGetSubtypesReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetSubtypesReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setTypeGetSubtypesReq(newTypeRequest(label), &typedb_protocol.Type_GetSubtypes_Req{}))
}

func setTypeGetSupertypeReq(typeReq *typedb_protocol.Type_Req, typeSetSupertypeReq *typedb_protocol.Type_SetSupertype_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_TypeSetSupertypeReq{
		TypeSetSupertypeReq: typeSetSupertypeReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetSupertypeReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setTypeGetSupertypeReq(newTypeRequest(label), &typedb_protocol.Type_SetSupertype_Req{}))
}

func setTypeDeleteReq(typeReq *typedb_protocol.Type_Req, typeDeleteReq *typedb_protocol.Type_Delete_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_TypeDeleteReq{
		TypeDeleteReq: typeDeleteReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func DeleteReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setTypeDeleteReq(newTypeRequest(label), &typedb_protocol.Type_Delete_Req{}))
}