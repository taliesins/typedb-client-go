package transaction

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func ClientReq(reqs []*typedb_protocol.Transaction_Req) *typedb_protocol.Transaction_Client {
	return &typedb_protocol.Transaction_Client{
		Reqs: reqs,
	}
}

func OpenReq(sessionId []byte, transactionType typedb_protocol.Transaction_Type, options *typedb_protocol.Options, latencyMillis int32) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_OpenReq{
			OpenReq: &typedb_protocol.Transaction_Open_Req{
				SessionId: sessionId,
				Type: transactionType,
				Options: options,
				NetworkLatencyMillis: latencyMillis,
			},
		},
	}
}

func CommitReq() *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_CommitReq{
			CommitReq: &typedb_protocol.Transaction_Commit_Req{
			},
		},
	}
}

func RollbackReq() *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_RollbackReq{
			RollbackReq: &typedb_protocol.Transaction_Rollback_Req{
			},
		},
	}
}

func StreamReq(requestId []byte) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_StreamReq{
			StreamReq: &typedb_protocol.Transaction_Stream_Req{
			},
		},
		ReqId: requestId,
	}
}

