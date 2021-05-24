package attribute_type

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

func setAttributeTypeGetOwnersReq(typeReq *grakn_protocol.Type_Req, attributeTypeGetOwnersReq *grakn_protocol.AttributeType_GetOwners_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_AttributeTypeGetOwnersReq{
		AttributeTypeGetOwnersReq: attributeTypeGetOwnersReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetOwnersReq(label common.Label, onlyKey bool) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypeGetOwnersReq(newTypeRequest(label),
		&grakn_protocol.AttributeType_GetOwners_Req{
			OnlyKey: onlyKey,
		}))
}

func setAttributeTypePutReq(typeReq *grakn_protocol.Type_Req, attributeTypePutReq *grakn_protocol.AttributeType_Put_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_AttributeTypePutReq{
		AttributeTypePutReq: attributeTypePutReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func PutReq(label common.Label, value *grakn_protocol.Attribute_Value) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypePutReq(newTypeRequest(label),
		&grakn_protocol.AttributeType_Put_Req{
			Value: value,
		}))
}

func setAttributeTypeGetReq(typeReq *grakn_protocol.Type_Req, attributeTypeGetReq *grakn_protocol.AttributeType_Get_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_AttributeTypeGetReq{
		AttributeTypeGetReq: attributeTypeGetReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetReq(label common.Label, value *grakn_protocol.Attribute_Value) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypeGetReq(newTypeRequest(label),
		&grakn_protocol.AttributeType_Get_Req{
			Value: value,
		},
	))
}

func setAttributeTypeGetRegexReq(typeReq *grakn_protocol.Type_Req, attributeTypeGetRegexReq *grakn_protocol.AttributeType_GetRegex_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_AttributeTypeGetRegexReq{
		AttributeTypeGetRegexReq: attributeTypeGetRegexReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetRegexReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypeGetRegexReq(newTypeRequest(label), &grakn_protocol.AttributeType_GetRegex_Req{}))
}

func setAttributeTypeSetRegexReq(typeReq *grakn_protocol.Type_Req, attributeTypeSetRegexReq *grakn_protocol.AttributeType_SetRegex_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_AttributeTypeSetRegexReq{
		AttributeTypeSetRegexReq: attributeTypeSetRegexReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetRegexReq(label common.Label, regex string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeTypeSetRegexReq(newTypeRequest(label),
		&grakn_protocol.AttributeType_SetRegex_Req{
			Regex: regex,
		}))
}
