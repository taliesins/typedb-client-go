package session

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func OpenReq(database string, sessionType typedb_protocol.Session_Type, options *typedb_protocol.Options) *typedb_protocol.Session_Open_Req {
	return &typedb_protocol.Session_Open_Req{
		Type: sessionType,
		Options: options,
		Database: database,
	}
}

func CloseReq(id []byte) *typedb_protocol.Session_Close_Req {
	return &typedb_protocol.Session_Close_Req{
		SessionId: id,
	}
}

func PulseReq(id []byte) *typedb_protocol.Session_Pulse_Req {
	return &typedb_protocol.Session_Pulse_Req{
		SessionId: id,
	}
}
