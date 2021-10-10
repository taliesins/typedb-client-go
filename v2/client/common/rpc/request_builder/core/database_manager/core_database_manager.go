package database_manager

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func CreateReq(name string) *typedb_protocol.CoreDatabaseManager_Create_Req {
	return &typedb_protocol.CoreDatabaseManager_Create_Req{
		Name: name,
	}
}

func ContainsReq(name string) *typedb_protocol.CoreDatabaseManager_Create_Req {
	return &typedb_protocol.CoreDatabaseManager_Create_Req{
		Name: name,
	}
}

func AllReq() *typedb_protocol.CoreDatabaseManager_All_Req {
	return &typedb_protocol.CoreDatabaseManager_All_Req{}
}