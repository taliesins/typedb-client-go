package thing

import (
	"encoding/hex"
	"github.com/taliesins/typedb-client-go/v2/grakn_protocol"
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

func setThingGetHasReq(thingReq *grakn_protocol.Thing_Req, thingGetHasReq *grakn_protocol.Thing_GetHas_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_ThingGetHasReq{
		ThingGetHasReq: thingGetHasReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func GetHasReq(iid string, onlyKey bool) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingGetHasReq(newThingRequest(iid),
		&grakn_protocol.Thing_GetHas_Req{
			KeysOnly: onlyKey,
		}))
}

func GetHasByTypeReq(iid string, attributeTypes []*grakn_protocol.Type) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingGetHasReq(newThingRequest(iid),
		&grakn_protocol.Thing_GetHas_Req{
			AttributeTypes: attributeTypes,
		}))
}

func setThingSetHasReq(thingReq *grakn_protocol.Thing_Req, thingSetHasReq *grakn_protocol.Thing_SetHas_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_ThingSetHasReq{
		ThingSetHasReq: thingSetHasReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func SetHasReq(iid string, attribute *grakn_protocol.Thing) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingSetHasReq(newThingRequest(iid),
		&grakn_protocol.Thing_SetHas_Req{
			Attribute: attribute,
		}))
}

func setThingUnsetHasReq(thingReq *grakn_protocol.Thing_Req, thingUnsetHasReq *grakn_protocol.Thing_UnsetHas_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_ThingUnsetHasReq{
		ThingUnsetHasReq: thingUnsetHasReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func UnsetHasReq(iid string, attribute *grakn_protocol.Thing) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingUnsetHasReq(newThingRequest(iid),
		&grakn_protocol.Thing_UnsetHas_Req{
			Attribute: attribute,
		}))
}

func setThingGetPlayingReq(thingReq *grakn_protocol.Thing_Req, thingGetPlayingReq *grakn_protocol.Thing_GetPlaying_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_ThingGetPlayingReq{
		ThingGetPlayingReq: thingGetPlayingReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func GetPlayingReq(iid string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingGetPlayingReq(newThingRequest(iid),
		&grakn_protocol.Thing_GetPlaying_Req{}))
}

func setThingGetRelationsReq(thingReq *grakn_protocol.Thing_Req, thingGetRelationsReq *grakn_protocol.Thing_GetRelations_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_ThingGetRelationsReq{
		ThingGetRelationsReq: thingGetRelationsReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func GetRelationsReq(iid string, roleTypes []*grakn_protocol.Type) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingGetRelationsReq(newThingRequest(iid),
		&grakn_protocol.Thing_GetRelations_Req{
			RoleTypes: roleTypes,
		}))
}

func setThingDeleteReq(thingReq *grakn_protocol.Thing_Req, thingDeleteReq *grakn_protocol.Thing_Delete_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_ThingDeleteReq{
		ThingDeleteReq: thingDeleteReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func DeleteReq(iid string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setThingDeleteReq(newThingRequest(iid),
		&grakn_protocol.Thing_Delete_Req{}))
}
