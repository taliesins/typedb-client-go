package relation

import (
	"encoding/hex"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

func AddPlayerReq(iid string, roleType *typedb_protocol.Type, player *typedb_protocol.Thing) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationAddPlayerReq(newThingRequest(iid),
		&typedb_protocol.Relation_AddPlayer_Req{
			RoleType: roleType,
			Player:   player,
		}))
}

func RemovePlayerReq(iid string, roleType *typedb_protocol.Type, player *typedb_protocol.Thing) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationRemovePlayerReq(newThingRequest(iid),
		&typedb_protocol.Relation_RemovePlayer_Req{
			RoleType: roleType,
			Player:   player,
		}))
}

func GetPlayersReq(iid string, roleTypes []*typedb_protocol.Type) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationGetPlayersReq(newThingRequest(iid),
		&typedb_protocol.Relation_GetPlayers_Req{
			RoleTypes: roleTypes,
		}))
}

func GetPlayersByRoleTypeReq(iid string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationGetPlayersByRoleTypeReq(newThingRequest(iid),
		&typedb_protocol.Relation_GetPlayersByRoleType_Req{
		}))
}

func GetRelatingReq(iid string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationGetRelatingReq(newThingRequest(iid),
		&typedb_protocol.Relation_GetRelating_Req{
		}))
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

func setRelationAddPlayerReq(thingReq *typedb_protocol.Thing_Req, relationAddPlayerReq *typedb_protocol.Relation_AddPlayer_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_RelationAddPlayerReq{
		RelationAddPlayerReq: relationAddPlayerReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func setRelationRemovePlayerReq(thingReq *typedb_protocol.Thing_Req, relationRemovePlayerReq *typedb_protocol.Relation_RemovePlayer_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_RelationRemovePlayerReq{
		RelationRemovePlayerReq: relationRemovePlayerReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func setRelationGetPlayersReq(thingReq *typedb_protocol.Thing_Req, relationGetPlayersReq *typedb_protocol.Relation_GetPlayers_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_RelationGetPlayersReq{
		RelationGetPlayersReq: relationGetPlayersReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func setRelationGetPlayersByRoleTypeReq(thingReq *typedb_protocol.Thing_Req, relationGetPlayersByRoleTypeReq *typedb_protocol.Relation_GetPlayersByRoleType_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_RelationGetPlayersByRoleTypeReq{
		RelationGetPlayersByRoleTypeReq: relationGetPlayersByRoleTypeReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func setRelationGetRelatingReq(thingReq *typedb_protocol.Thing_Req, relationGetRelatingReq *typedb_protocol.Relation_GetRelating_Req) *typedb_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &typedb_protocol.Thing_Req_RelationGetRelatingReq{
		RelationGetRelatingReq: relationGetRelatingReq,
	}

	return &typedb_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}
