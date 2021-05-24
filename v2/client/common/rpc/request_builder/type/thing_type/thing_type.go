package thing_type

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

func protoThingType(label common.Label, encoding grakn_protocol.Type_Encoding) *grakn_protocol.Type {
	return &grakn_protocol.Type{
		Label:    label.Name,
		Scope:    label.Scope,
		Encoding: encoding,
	}
}

func setThingTypeSetAbstractReq(typeReq *grakn_protocol.Type_Req, thingTypeSetAbstractReq *grakn_protocol.ThingType_SetAbstract_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_ThingTypeSetAbstractReq{
		ThingTypeSetAbstractReq: thingTypeSetAbstractReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetAbstractReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeSetAbstractReq(newTypeRequest(label), &grakn_protocol.ThingType_SetAbstract_Req{}))
}

func setThingTypeUnsetAbstractReq(typeReq *grakn_protocol.Type_Req, thingTypeUnsetAbstractReq *grakn_protocol.ThingType_UnsetAbstract_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_ThingTypeUnsetAbstractReq{
		ThingTypeUnsetAbstractReq: thingTypeUnsetAbstractReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func UnsetAbstractReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeUnsetAbstractReq(newTypeRequest(label), &grakn_protocol.ThingType_UnsetAbstract_Req{}))
}

func setTypeSetSupertypeReq(typeReq *grakn_protocol.Type_Req, typeSetSupertypeReq *grakn_protocol.Type_SetSupertype_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_TypeSetSupertypeReq{
		TypeSetSupertypeReq: typeSetSupertypeReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetSupertypeReq(label common.Label, supertype *grakn_protocol.Type) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setTypeSetSupertypeReq(newTypeRequest(label),
		&grakn_protocol.Type_SetSupertype_Req{
			Type: supertype,
		}))
}

func setThingTypeGetPlaysReq(typeReq *grakn_protocol.Type_Req, thingTypeGetPlaysReq *grakn_protocol.ThingType_GetPlays_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_ThingTypeGetPlaysReq{
		ThingTypeGetPlaysReq: thingTypeGetPlaysReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetPlaysReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeGetPlaysReq(newTypeRequest(label), &grakn_protocol.ThingType_GetPlays_Req{}))
}

func setThingTypeSetPlaysReq(typeReq *grakn_protocol.Type_Req, thingTypeSetPlaysReq *grakn_protocol.ThingType_SetPlays_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_ThingTypeSetPlaysReq{
		ThingTypeSetPlaysReq: thingTypeSetPlaysReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetPlaysReq(label common.Label, roleType *grakn_protocol.Type) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeSetPlaysReq(newTypeRequest(label),
		&grakn_protocol.ThingType_SetPlays_Req{
			Role: roleType,
		}))
}

func SetPlaysOverriddenReq(label common.Label, roleType *grakn_protocol.Type, overriddenRoleType *grakn_protocol.ThingType_SetPlays_Req_OverriddenRole) *grakn_protocol.Transaction_Req {
return newTransactionRequest(setThingTypeSetPlaysReq(newTypeRequest(label),
	&grakn_protocol.ThingType_SetPlays_Req{
		Role: roleType,
		Overridden: overriddenRoleType,
	}))
}

func setThingTypeUnsetPlaysReq(typeReq *grakn_protocol.Type_Req, thingTypeUnsetPlaysReq *grakn_protocol.ThingType_UnsetPlays_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_ThingTypeUnsetPlaysReq{
		ThingTypeUnsetPlaysReq: thingTypeUnsetPlaysReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func UnsetPlaysReq(label common.Label, roleType *grakn_protocol.Type) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeUnsetPlaysReq(newTypeRequest(label),
		&grakn_protocol.ThingType_UnsetPlays_Req{
			Role: roleType,
		}))
}

func setThingTypeGetOwnsReq(typeReq *grakn_protocol.Type_Req, thingTypeGetOwnsReq *grakn_protocol.ThingType_GetOwns_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_ThingTypeGetOwnsReq{
		ThingTypeGetOwnsReq: thingTypeGetOwnsReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetOwnsReq(label common.Label, keysOnly bool) *grakn_protocol.Transaction_Req {
return newTransactionRequest(setThingTypeGetOwnsReq(newTypeRequest(label),
	&grakn_protocol.ThingType_GetOwns_Req{
		KeysOnly: keysOnly,
	}))
}

func GetOwnsByTypeReq(label common.Label, valueType *grakn_protocol.ThingType_GetOwns_Req_ValueType, keysOnly bool) *grakn_protocol.Transaction_Req {
return newTransactionRequest(setThingTypeGetOwnsReq(newTypeRequest(label),
	&grakn_protocol.ThingType_GetOwns_Req{
		KeysOnly: keysOnly,
		Filter: valueType,
	}))
}

func setThingTypeSetOwnsReq(typeReq *grakn_protocol.Type_Req, thingTypeSetOwnsReq *grakn_protocol.ThingType_SetOwns_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_ThingTypeSetOwnsReq{
		ThingTypeSetOwnsReq: thingTypeSetOwnsReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetOwnsReq(label common.Label, attributeType *grakn_protocol.Type, isKey bool) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeSetOwnsReq(newTypeRequest(label),
		&grakn_protocol.ThingType_SetOwns_Req{
			AttributeType: attributeType,
			IsKey:         isKey,
		}))
}

func SetOwnsOverriddenReq(label common.Label, attributeType *grakn_protocol.Type, overriddenType *grakn_protocol.ThingType_SetOwns_Req_OverriddenType, isKey bool) *grakn_protocol.Transaction_Req {
return newTransactionRequest(setThingTypeSetOwnsReq(newTypeRequest(label),
	&grakn_protocol.ThingType_SetOwns_Req{
		AttributeType: attributeType,
		Overridden: overriddenType,
		IsKey: isKey,
	}))
}

func setThingTypeUnsetOwnsReq(typeReq *grakn_protocol.Type_Req, thingTypeUnsetOwnsReq *grakn_protocol.ThingType_UnsetOwns_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_ThingTypeUnsetOwnsReq{
		ThingTypeUnsetOwnsReq: thingTypeUnsetOwnsReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func UnsetOwnsReq(label common.Label, attributeType *grakn_protocol.Type) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeUnsetOwnsReq(newTypeRequest(label),
		&grakn_protocol.ThingType_UnsetOwns_Req{
			AttributeType: attributeType,
		}))
}

func setThingTypeGetInstancesReq(typeReq *grakn_protocol.Type_Req, thingTypeGetInstancesReq *grakn_protocol.ThingType_GetInstances_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_ThingTypeGetInstancesReq{
		ThingTypeGetInstancesReq: thingTypeGetInstancesReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetInstancesReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeGetInstancesReq(newTypeRequest(label), &grakn_protocol.ThingType_GetInstances_Req{}))
}
