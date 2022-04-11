package user_token

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func RenewReq(username string) *typedb_protocol.ClusterUserToken_Renew_Req {
	return &typedb_protocol.ClusterUserToken_Renew_Req{
		Username: username,
	}
}