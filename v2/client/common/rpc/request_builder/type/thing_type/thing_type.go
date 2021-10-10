package thing_type

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

func protoThingType(label common.Label, encoding typedb_protocol.Type_Encoding) *typedb_protocol.Type {
	return &typedb_protocol.Type{
		Label:    label.Name,
		Scope:    label.Scope,
		Encoding: encoding,
	}
}

func setThingTypeSetAbstractReq(typeReq *typedb_protocol.Type_Req, thingTypeSetAbstractReq *typedb_protocol.ThingType_SetAbstract_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_ThingTypeSetAbstractReq{
		ThingTypeSetAbstractReq: thingTypeSetAbstractReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetAbstractReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeSetAbstractReq(newTypeRequest(label), &typedb_protocol.ThingType_SetAbstract_Req{}))
}

func setThingTypeUnsetAbstractReq(typeReq *typedb_protocol.Type_Req, thingTypeUnsetAbstractReq *typedb_protocol.ThingType_UnsetAbstract_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_ThingTypeUnsetAbstractReq{
		ThingTypeUnsetAbstractReq: thingTypeUnsetAbstractReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func UnsetAbstractReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeUnsetAbstractReq(newTypeRequest(label), &typedb_protocol.ThingType_UnsetAbstract_Req{}))
}

func setTypeSetSupertypeReq(typeReq *typedb_protocol.Type_Req, typeSetSupertypeReq *typedb_protocol.Type_SetSupertype_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_TypeSetSupertypeReq{
		TypeSetSupertypeReq: typeSetSupertypeReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetSupertypeReq(label common.Label, supertype *typedb_protocol.Type) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setTypeSetSupertypeReq(newTypeRequest(label),
		&typedb_protocol.Type_SetSupertype_Req{
			Type: supertype,
		}))
}

func setThingTypeGetPlaysReq(typeReq *typedb_protocol.Type_Req, thingTypeGetPlaysReq *typedb_protocol.ThingType_GetPlays_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_ThingTypeGetPlaysReq{
		ThingTypeGetPlaysReq: thingTypeGetPlaysReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetPlaysReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeGetPlaysReq(newTypeRequest(label), &typedb_protocol.ThingType_GetPlays_Req{}))
}

func setThingTypeSetPlaysReq(typeReq *typedb_protocol.Type_Req, thingTypeSetPlaysReq *typedb_protocol.ThingType_SetPlays_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_ThingTypeSetPlaysReq{
		ThingTypeSetPlaysReq: thingTypeSetPlaysReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetPlaysReq(label common.Label, roleType *typedb_protocol.Type) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeSetPlaysReq(newTypeRequest(label),
		&typedb_protocol.ThingType_SetPlays_Req{
			Role: roleType,
		}))
}

func SetPlaysOverriddenReq(label common.Label, roleType *typedb_protocol.Type, overriddenRoleType *typedb_protocol.ThingType_SetPlays_Req_OverriddenRole) *typedb_protocol.Transaction_Req {
return newTransactionRequest(setThingTypeSetPlaysReq(newTypeRequest(label),
	&typedb_protocol.ThingType_SetPlays_Req{
		Role: roleType,
		Overridden: overriddenRoleType,
	}))
}

func setThingTypeUnsetPlaysReq(typeReq *typedb_protocol.Type_Req, thingTypeUnsetPlaysReq *typedb_protocol.ThingType_UnsetPlays_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_ThingTypeUnsetPlaysReq{
		ThingTypeUnsetPlaysReq: thingTypeUnsetPlaysReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func UnsetPlaysReq(label common.Label, roleType *typedb_protocol.Type) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeUnsetPlaysReq(newTypeRequest(label),
		&typedb_protocol.ThingType_UnsetPlays_Req{
			Role: roleType,
		}))
}

func setThingTypeGetOwnsReq(typeReq *typedb_protocol.Type_Req, thingTypeGetOwnsReq *typedb_protocol.ThingType_GetOwns_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_ThingTypeGetOwnsReq{
		ThingTypeGetOwnsReq: thingTypeGetOwnsReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetOwnsReq(label common.Label, keysOnly bool) *typedb_protocol.Transaction_Req {
return newTransactionRequest(setThingTypeGetOwnsReq(newTypeRequest(label),
	&typedb_protocol.ThingType_GetOwns_Req{
		KeysOnly: keysOnly,
	}))
}

func GetOwnsByTypeReq(label common.Label, valueType *typedb_protocol.ThingType_GetOwns_Req_ValueType, keysOnly bool) *typedb_protocol.Transaction_Req {
return newTransactionRequest(setThingTypeGetOwnsReq(newTypeRequest(label),
	&typedb_protocol.ThingType_GetOwns_Req{
		KeysOnly: keysOnly,
		Filter: valueType,
	}))
}

func setThingTypeSetOwnsReq(typeReq *typedb_protocol.Type_Req, thingTypeSetOwnsReq *typedb_protocol.ThingType_SetOwns_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_ThingTypeSetOwnsReq{
		ThingTypeSetOwnsReq: thingTypeSetOwnsReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetOwnsReq(label common.Label, attributeType *typedb_protocol.Type, isKey bool) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeSetOwnsReq(newTypeRequest(label),
		&typedb_protocol.ThingType_SetOwns_Req{
			AttributeType: attributeType,
			IsKey:         isKey,
		}))
}

func SetOwnsOverriddenReq(label common.Label, attributeType *typedb_protocol.Type, overriddenType *typedb_protocol.ThingType_SetOwns_Req_OverriddenType, isKey bool) *typedb_protocol.Transaction_Req {
return newTransactionRequest(setThingTypeSetOwnsReq(newTypeRequest(label),
	&typedb_protocol.ThingType_SetOwns_Req{
		AttributeType: attributeType,
		Overridden: overriddenType,
		IsKey: isKey,
	}))
}

func setThingTypeUnsetOwnsReq(typeReq *typedb_protocol.Type_Req, thingTypeUnsetOwnsReq *typedb_protocol.ThingType_UnsetOwns_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_ThingTypeUnsetOwnsReq{
		ThingTypeUnsetOwnsReq: thingTypeUnsetOwnsReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func UnsetOwnsReq(label common.Label, attributeType *typedb_protocol.Type) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeUnsetOwnsReq(newTypeRequest(label),
		&typedb_protocol.ThingType_UnsetOwns_Req{
			AttributeType: attributeType,
		}))
}

func setThingTypeGetInstancesReq(typeReq *typedb_protocol.Type_Req, thingTypeGetInstancesReq *typedb_protocol.ThingType_GetInstances_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_ThingTypeGetInstancesReq{
		ThingTypeGetInstancesReq: thingTypeGetInstancesReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetInstancesReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingTypeGetInstancesReq(newTypeRequest(label), &typedb_protocol.ThingType_GetInstances_Req{}))
}
