package relation

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

func setRelationAddPlayerReq(thingReq *grakn_protocol.Thing_Req, relationAddPlayerReq *grakn_protocol.Relation_AddPlayer_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_RelationAddPlayerReq{
		RelationAddPlayerReq: relationAddPlayerReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func AddPlayerReq(iid string, roleType *grakn_protocol.Type, player *grakn_protocol.Thing) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationAddPlayerReq(newThingRequest(iid),
		&grakn_protocol.Relation_AddPlayer_Req{
			RoleType: roleType,
			Player:   player,
		}))
}

func setRelationRemovePlayerReq(thingReq *grakn_protocol.Thing_Req, relationRemovePlayerReq *grakn_protocol.Relation_RemovePlayer_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_RelationRemovePlayerReq{
		RelationRemovePlayerReq: relationRemovePlayerReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func RemovePlayerReq(iid string, roleType *grakn_protocol.Type, player *grakn_protocol.Thing) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationRemovePlayerReq(newThingRequest(iid),
		&grakn_protocol.Relation_RemovePlayer_Req{
			RoleType: roleType,
			Player:   player,
		}))
}

func setRelationGetPlayersReq(thingReq *grakn_protocol.Thing_Req, relationGetPlayersReq *grakn_protocol.Relation_GetPlayers_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_RelationGetPlayersReq{
		RelationGetPlayersReq: relationGetPlayersReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func GetPlayersReq(iid string, roleTypes []*grakn_protocol.Type) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationGetPlayersReq(newThingRequest(iid),
		&grakn_protocol.Relation_GetPlayers_Req{
			RoleTypes: roleTypes,
		}))
}

func setRelationGetPlayersByRoleTypeReq(thingReq *grakn_protocol.Thing_Req, relationGetPlayersByRoleTypeReq *grakn_protocol.Relation_GetPlayersByRoleType_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_RelationGetPlayersByRoleTypeReq{
		RelationGetPlayersByRoleTypeReq: relationGetPlayersByRoleTypeReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func GetPlayersByRoleTypeReq(iid string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationGetPlayersByRoleTypeReq(newThingRequest(iid),
		&grakn_protocol.Relation_GetPlayersByRoleType_Req{
		}))
}

func setRelationGetRelatingReq(thingReq *grakn_protocol.Thing_Req, relationGetRelatingReq *grakn_protocol.Relation_GetRelating_Req) *grakn_protocol.Transaction_Req_ThingReq {
	thingReq.Req = &grakn_protocol.Thing_Req_RelationGetRelatingReq{
		RelationGetRelatingReq: relationGetRelatingReq,
	}

	return &grakn_protocol.Transaction_Req_ThingReq{
		ThingReq: thingReq,
	}
}

func GetRelatingReq(iid string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationGetRelatingReq(newThingRequest(iid),
		&grakn_protocol.Relation_GetRelating_Req{
		}))
}
