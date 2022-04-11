package user

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func PasswordReq(name string, password string) *typedb_protocol.ClusterUser_Password_Req {
	return &typedb_protocol.ClusterUser_Password_Req{
		Username: name,
		Password: password,
	}
}

func DeleteReq(name string) *typedb_protocol.ClusterUser_Delete_Req {
	return &typedb_protocol.ClusterUser_Delete_Req{
		Username: name,
	}
}
