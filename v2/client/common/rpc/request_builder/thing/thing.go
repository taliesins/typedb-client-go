package thing

import (
	"encoding/hex"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

func ThingReq(req *typedb_protocol.Thing_Req) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_ThingReq{
			ThingReq: req,
		},
	}
}

func ProtoThing(iid string) *typedb_protocol.Thing_Req {
	data, err := hex.DecodeString(iid)
	if err != nil {
		panic(err)
	}

	return &typedb_protocol.Thing_Req{
		Iid: data,
	}
}

func GetHasReq(iid string, onlyKey bool) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingGetHasReq(ProtoThing(iid),
		&typedb_protocol.Thing_GetHas_Req{
			KeysOnly: onlyKey,
		}))
}

func GetHasByTypeReq(iid string, attributeTypes []*typedb_protocol.Type) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingGetHasReq(ProtoThing(iid),
		&typedb_protocol.Thing_GetHas_Req{
			AttributeTypes: attributeTypes,
		}))
}

func SetHasReq(iid string, attribute *typedb_protocol.Thing) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingSetHasReq(ProtoThing(iid),
		&typedb_protocol.Thing_SetHas_Req{
			Attribute: attribute,
		}))
}

func UnsetHasReq(iid string, attribute *typedb_protocol.Thing) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingUnsetHasReq(ProtoThing(iid),
		&typedb_protocol.Thing_UnsetHas_Req{
			Attribute: attribute,
		}))
}

func GetPlayingReq(iid string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingGetPlayingReq(ProtoThing(iid),
		&typedb_protocol.Thing_GetPlaying_Req{}))
}

func GetRelationsReq(iid string, roleTypes []*typedb_protocol.Type) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingGetRelationsReq(ProtoThing(iid),
		&typedb_protocol.Thing_GetRelations_Req{
			RoleTypes: roleTypes,
		}))
}

func DeleteReq(iid string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setThingDeleteReq(ProtoThing(iid),
		&typedb_protocol.Thing_Delete_Req{}))
}


func newTransactionRequest(req *typedb_protocol.Transaction_Req_ThingReq) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: req,
	}
}

func setThingGetHasReq(thingReq *typedb_protocol.Thing_Req, thingGetHasReq *typedb_protocol.Thing_GetHas_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_ThingGetHasReq{
		ThingGetHasReq: thingGetHasReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func setThingSetHasReq(thingReq *typedb_protocol.Thing_Req, thingSetHasReq *typedb_protocol.Thing_SetHas_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_ThingSetHasReq{
		ThingSetHasReq: thingSetHasReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func setThingUnsetHasReq(thingReq *typedb_protocol.Thing_Req, thingUnsetHasReq *typedb_protocol.Thing_UnsetHas_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_ThingUnsetHasReq{
		ThingUnsetHasReq: thingUnsetHasReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func setThingGetPlayingReq(thingReq *typedb_protocol.Thing_Req, thingGetPlayingReq *typedb_protocol.Thing_GetPlaying_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_ThingGetPlayingReq{
		ThingGetPlayingReq: thingGetPlayingReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func setThingGetRelationsReq(thingReq *typedb_protocol.Thing_Req, thingGetRelationsReq *typedb_protocol.Thing_GetRelations_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_ThingGetRelationsReq{
		ThingGetRelationsReq: thingGetRelationsReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func setThingDeleteReq(thingReq *typedb_protocol.Thing_Req, thingDeleteReq *typedb_protocol.Thing_Delete_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_ThingDeleteReq{
		ThingDeleteReq: thingDeleteReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}
