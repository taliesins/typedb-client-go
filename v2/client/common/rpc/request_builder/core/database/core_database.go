package database

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func SchemaReq(name string) *grakn_protocol.CoreDatabase_Schema_Req {
	return &grakn_protocol.CoreDatabase_Schema_Req{
		Name: name,
	}
}

func DeleteReq(name string) *grakn_protocol.CoreDatabase_Delete_Req {
	return &grakn_protocol.CoreDatabase_Delete_Req{
		Name: name,
	}
}
