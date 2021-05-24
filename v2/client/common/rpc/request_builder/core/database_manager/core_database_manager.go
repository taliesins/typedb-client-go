package database_manager

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func CreateReq(name string) *grakn_protocol.CoreDatabaseManager_Create_Req {
	return &grakn_protocol.CoreDatabaseManager_Create_Req{
		Name: name,
	}
}

func ContainsReq(name string) *grakn_protocol.CoreDatabaseManager_Create_Req {
	return &grakn_protocol.CoreDatabaseManager_Create_Req{
		Name: name,
	}
}

func AllReq() *grakn_protocol.CoreDatabaseManager_All_Req {
	return &grakn_protocol.CoreDatabaseManager_All_Req{}
}