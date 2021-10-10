package database

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func SchemaReq(name string) *typedb_protocol.CoreDatabase_Schema_Req {
	return &typedb_protocol.CoreDatabase_Schema_Req{
		Name: name,
	}
}

func DeleteReq(name string) *typedb_protocol.CoreDatabase_Delete_Req {
	return &typedb_protocol.CoreDatabase_Delete_Req{
		Name: name,
	}
}
