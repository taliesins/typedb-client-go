package session

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func OpenReq(database string, sessionType grakn_protocol.Session_Type, options *grakn_protocol.Options) *grakn_protocol.Session_Open_Req {
	return &grakn_protocol.Session_Open_Req{
		Type: sessionType,
		Options: options,
		Database: database,
	}
}

func CloseReq(id []byte) *grakn_protocol.Session_Close_Req {
	return &grakn_protocol.Session_Close_Req{
		SessionId: id,
	}
}

func PulseReq(id []byte) *grakn_protocol.Session_Pulse_Req {
	return &grakn_protocol.Session_Pulse_Req{
		SessionId: id,
	}
}
