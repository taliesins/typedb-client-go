package server_manager

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func AllReq() *grakn_protocol.ServerManager_All_Req {
	return &grakn_protocol.ServerManager_All_Req{}
}
