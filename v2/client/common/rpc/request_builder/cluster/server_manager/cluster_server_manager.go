package server_manager

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func AllReq() *typedb_protocol.ServerManager_All_Req {
	return &typedb_protocol.ServerManager_All_Req{}
}
