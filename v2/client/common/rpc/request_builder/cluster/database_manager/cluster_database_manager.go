package database_manager

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func GetReq(name string) *grakn_protocol.ClusterDatabaseManager_Get_Req {
	return &grakn_protocol.ClusterDatabaseManager_Get_Req{
		Name: name,
	}
}

func AllReq() *grakn_protocol.ClusterDatabaseManager_All_Req {
	return &grakn_protocol.ClusterDatabaseManager_All_Req{}
}
