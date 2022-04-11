package attribute

import (
	"encoding/hex"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
	"time"
)

func GetOwnersReq(iid string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeGetOwnersReq(newThingRequest(iid),
		&typedb_protocol.Attribute_GetOwners_Req{}))
}

func GetOwnersByTypeReq(iid string, ownerType *typedb_protocol.Type) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeGetOwnersReq(newThingRequest(iid),
		&typedb_protocol.Attribute_GetOwners_Req{
			Filter: &typedb_protocol.Attribute_GetOwners_Req_ThingType{
				ThingType: ownerType,
			}}))
}

func AttributeValueBooleanReq(value bool) *typedb_protocol.Attribute_Value {
	return &typedb_protocol.Attribute_Value{
		Value: &typedb_protocol.Attribute_Value_Boolean{
			Boolean: value,
		},
	}
}

func AttributeValueLongReq(value int64) *typedb_protocol.Attribute_Value {
	return &typedb_protocol.Attribute_Value{
		Value: &typedb_protocol.Attribute_Value_Long{
			Long: value,
		},
	}
}

func AttributeValueDoubleReq(value float64) *typedb_protocol.Attribute_Value {
	return &typedb_protocol.Attribute_Value{
		Value: &typedb_protocol.Attribute_Value_Double{
			Double: value,
		},
	}
}

func AttributeValueStringReq(value string) *typedb_protocol.Attribute_Value {
	return &typedb_protocol.Attribute_Value{
		Value: &typedb_protocol.Attribute_Value_String_{
			String_: value,
		},
	}
}

func AttributeValueDateTimeReq(value time.Time) *typedb_protocol.Attribute_Value {
	return &typedb_protocol.Attribute_Value{
		Value: &typedb_protocol.Attribute_Value_DateTime{
			DateTime: value.Unix(),
		},
	}
}

func newTransactionRequest(req *typedb_protocol.Transaction_Req_ThingReq) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: req,
	}
}

func newThingRequest(iid string) *typedb_protocol.Thing_Req {
	data, err := hex.DecodeString(iid)
	if err != nil {
		panic(err)
	}

	return &typedb_protocol.Thing_Req{
		Iid: data,
	}
}

func setAttributeGetOwnersReq(thingReq *typedb_protocol.Thing_Req, attributeGetOwnersReq *typedb_protocol.Attribute_GetOwners_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_AttributeGetOwnersReq{
		AttributeGetOwnersReq: attributeGetOwnersReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}
