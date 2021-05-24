package attribute

import (
	"encoding/hex"
	"github.com/taliesins/typedb-client-go/v2/grakn_protocol"
	"time"
)

func newTransactionRequest(req *grakn_protocol.Transaction_Req_ThingReq) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: req,
	}
}

func newThingRequest(iid string) *grakn_protocol.Thing_Req {
	data, err := hex.DecodeString(iid)
	if err != nil {
		panic(err)
	}

	return &grakn_protocol.Thing_Req{
		Iid: data,
	}
}

func setAttributeGetOwnersReq(thingReq *grakn_protocol.Thing_Req, attributeGetOwnersReq *grakn_protocol.Attribute_GetOwners_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_AttributeGetOwnersReq{
		AttributeGetOwnersReq: attributeGetOwnersReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func GetOwnersReq(iid string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeGetOwnersReq(newThingRequest(iid),
		&grakn_protocol.Attribute_GetOwners_Req{}))
}

func GetOwnersByTypeReq(iid string, ownerType *grakn_protocol.Type) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setAttributeGetOwnersReq(newThingRequest(iid),
		&grakn_protocol.Attribute_GetOwners_Req{
			Filter: &grakn_protocol.Attribute_GetOwners_Req_ThingType{
				ThingType: ownerType,
			}}))
}

func AttributeValueBooleanReq(value bool) *grakn_protocol.Attribute_Value {
	return &grakn_protocol.Attribute_Value{
		Value: &grakn_protocol.Attribute_Value_Boolean{
			Boolean: value,
		},
	}
}

func AttributeValueLongReq(value int64) *grakn_protocol.Attribute_Value {
	return &grakn_protocol.Attribute_Value{
		Value: &grakn_protocol.Attribute_Value_Long{
			Long: value,
		},
	}
}

func AttributeValueDoubleReq(value float64) *grakn_protocol.Attribute_Value {
	return &grakn_protocol.Attribute_Value{
		Value: &grakn_protocol.Attribute_Value_Double{
			Double: value,
		},
	}
}

func AttributeValueStringReq(value string) *grakn_protocol.Attribute_Value {
	return &grakn_protocol.Attribute_Value{
		Value: &grakn_protocol.Attribute_Value_String_{
			String_: value,
		},
	}
}

func AttributeValueDateTimeReq(value time.Time) *grakn_protocol.Attribute_Value {
	return &grakn_protocol.Attribute_Value{
		Value: &grakn_protocol.Attribute_Value_DateTime{
			DateTime: value.Unix(),
		},
	}
}
