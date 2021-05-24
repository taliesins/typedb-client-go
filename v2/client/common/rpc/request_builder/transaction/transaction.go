package transaction

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func CommitReq() *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_CommitReq{
			CommitReq: &grakn_protocol.Transaction_Commit_Req{
			},
		},
	}
}

func RollbackReq() *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_RollbackReq{
			RollbackReq: &grakn_protocol.Transaction_Rollback_Req{
			},
		},
	}
}

func StreamReq(requestId []byte) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_StreamReq{
			StreamReq: &grakn_protocol.Transaction_Stream_Req{
			},
		},
		ReqId: requestId,
	}
}

func OpenReq(sessionId []byte, transactionType grakn_protocol.Transaction_Type, options *grakn_protocol.Options, latencyMillis int32) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_OpenReq{
			OpenReq: &grakn_protocol.Transaction_Open_Req{
				SessionId: sessionId,
				Type: transactionType,
				Options: options,
				NetworkLatencyMillis: latencyMillis,
			},
		},
	}
}