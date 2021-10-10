package attribute_type

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

func setAttributeTypeGetOwnersReq(typeReq *typedb_protocol.Type_Req, attributeTypeGetOwnersReq *typedb_protocol.AttributeType_GetOwners_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_AttributeTypeGetOwnersReq{
		AttributeTypeGetOwnersReq: attributeTypeGetOwnersReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetOwnersReq(label common.Label, onlyKey bool) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypeGetOwnersReq(newTypeRequest(label),
		&typedb_protocol.AttributeType_GetOwners_Req{
			OnlyKey: onlyKey,
		}))
}

func setAttributeTypePutReq(typeReq *typedb_protocol.Type_Req, attributeTypePutReq *typedb_protocol.AttributeType_Put_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_AttributeTypePutReq{
		AttributeTypePutReq: attributeTypePutReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func PutReq(label common.Label, value *typedb_protocol.Attribute_Value) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypePutReq(newTypeRequest(label),
		&typedb_protocol.AttributeType_Put_Req{
			Value: value,
		}))
}

func setAttributeTypeGetReq(typeReq *typedb_protocol.Type_Req, attributeTypeGetReq *typedb_protocol.AttributeType_Get_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_AttributeTypeGetReq{
		AttributeTypeGetReq: attributeTypeGetReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetReq(label common.Label, value *typedb_protocol.Attribute_Value) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypeGetReq(newTypeRequest(label),
		&typedb_protocol.AttributeType_Get_Req{
			Value: value,
		},
	))
}

func setAttributeTypeGetRegexReq(typeReq *typedb_protocol.Type_Req, attributeTypeGetRegexReq *typedb_protocol.AttributeType_GetRegex_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_AttributeTypeGetRegexReq{
		AttributeTypeGetRegexReq: attributeTypeGetRegexReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetRegexReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypeGetRegexReq(newTypeRequest(label), &typedb_protocol.AttributeType_GetRegex_Req{}))
}

func setAttributeTypeSetRegexReq(typeReq *typedb_protocol.Type_Req, attributeTypeSetRegexReq *typedb_protocol.AttributeType_SetRegex_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_AttributeTypeSetRegexReq{
		AttributeTypeSetRegexReq: attributeTypeSetRegexReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetRegexReq(label common.Label, regex string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypeSetRegexReq(newTypeRequest(label),
		&typedb_protocol.AttributeType_SetRegex_Req{
			Regex: regex,
		}))
}
