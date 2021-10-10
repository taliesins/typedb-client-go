package database_manager

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func GetReq(name string) *typedb_protocol.ClusterDatabaseManager_Get_Req {
	return &typedb_protocol.ClusterDatabaseManager_Get_Req{
		Name: name,
	}
}

func AllReq() *typedb_protocol.ClusterDatabaseManager_All_Req {
	return &typedb_protocol.ClusterDatabaseManager_All_Req{}
}
