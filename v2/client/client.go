package client

import (
	"fmt"
	grakn "github.com/taliesins/client-go/v2/grakn_protocol"
)

func RunInsertQuery(transactionClient grakn.Grakn_TransactionClient, requestId string, metadata map[string]string, query string, explain bool, batchSize int32, latencyMillis int32) (insertResponses []*grakn.Query_Insert_Res, err error){
	if batchSize == 0 {
		batchSize = 2147483647
	}

	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_QueryReq{
			QueryReq: &grakn.Query_Req{
				Req: &grakn.Query_Req_InsertReq {
					InsertReq: &grakn.Query_Insert_Req{
						Query: query,
					},
				},
				Options: &grakn.Options{
					BatchSizeOpt: &grakn.Options_BatchSize{
						BatchSize: batchSize,
					},
					ExplainOpt: &grakn.Options_Explain{
						Explain: explain,
					},
				},
			},
		},
		Metadata:      metadata,
		Id:            requestId,
		LatencyMillis: latencyMillis,
	})

	for {
		transactionResponse, err := transactionClient.Recv()
		if err != nil {
			return nil, fmt.Errorf("could not receive query response: %w", err)
		}

		if transactionResponse.GetContinue() {
			err = transactionClient.Send(&grakn.Transaction_Req{
				Req: &grakn.Transaction_Req_Continue{
					Continue: true,
				},
				Metadata:      metadata,
				Id:            requestId,
				LatencyMillis: latencyMillis,
			})
			if err != nil {
				return nil, fmt.Errorf("could not send query request iterator: %w", err)
			}
		}  else if transactionResponse.GetDone()  {
			break
		} else {
			queryResponse := transactionResponse.GetQueryRes()
			if queryResponse != nil {
				insertResponse := queryResponse.GetInsertRes()
				insertResponses = append(insertResponses, insertResponse)
			}
		}
	}

	return insertResponses, err
}

func RunDefineQuery(transactionClient grakn.Grakn_TransactionClient, requestId string, metadata map[string]string, query string, explain bool, batchSize int32, latencyMillis int32) (err error){
	if batchSize == 0 {
		batchSize = 2147483647
	}

	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_QueryReq{
			QueryReq: &grakn.Query_Req{
				Req: &grakn.Query_Req_DefineReq {
					DefineReq: &grakn.Query_Define_Req {
						Query: query,
					},
				},
				Options: &grakn.Options{
					BatchSizeOpt: &grakn.Options_BatchSize{
						BatchSize: batchSize,
					},
					ExplainOpt: &grakn.Options_Explain{
						Explain: explain,
					},
				},
			},
		},
		Metadata:      metadata,
		Id:            requestId,
		LatencyMillis: latencyMillis,
	})

	_, err = transactionClient.Recv()
	if err != nil {
		return fmt.Errorf("could not receive query response: %w", err)
	}

	return err
}

func RunDeleteQuery(transactionClient grakn.Grakn_TransactionClient, requestId string, metadata map[string]string, query string, explain bool, batchSize int32, latencyMillis int32) (err error) {
	if batchSize == 0 {
		batchSize = 2147483647
	}

	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_QueryReq{
			QueryReq: &grakn.Query_Req{
				Req: &grakn.Query_Req_DeleteReq{
					DeleteReq: &grakn.Query_Delete_Req{
						Query: query,
					},
				},
				Options: &grakn.Options{
					BatchSizeOpt: &grakn.Options_BatchSize{
						BatchSize: batchSize,
					},
					ExplainOpt: &grakn.Options_Explain{
						Explain: explain,
					},
				},
			},
		},
		Metadata:      metadata,
		Id:            requestId,
		LatencyMillis: latencyMillis,
	})

	_, err = transactionClient.Recv()
	if err != nil {
		return fmt.Errorf("could not receive query response: %w", err)
	}

	return err
}

func RunMatchQuery(transactionClient grakn.Grakn_TransactionClient, requestId string, metadata map[string]string, query string, explain bool, batchSize int32, latencyMillis int32) (matchResponses []*grakn.Query_Match_Res, err error) {
	if batchSize == 0 {
		batchSize = 2147483647
	}

	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_QueryReq{
			QueryReq: &grakn.Query_Req {
				Req: &grakn.Query_Req_MatchReq{
					MatchReq: &grakn.Query_Match_Req{
						Query: query,
					},
				},
				Options: &grakn.Options{
					BatchSizeOpt: &grakn.Options_BatchSize{
						BatchSize: batchSize,
					},
					ExplainOpt: &grakn.Options_Explain{
						Explain: explain,
					},
				},
			},
		},
		Metadata:      metadata,
		Id:            requestId,
		LatencyMillis: latencyMillis,
	})

	if err != nil {
		return nil, fmt.Errorf("could not send query request: %w", err)
	}

	for {
		transactionResponse, err := transactionClient.Recv()
		if err != nil {
			return nil, fmt.Errorf("could not receive query response: %w", err)
		}

		if transactionResponse.GetContinue() {
			err = transactionClient.Send(&grakn.Transaction_Req{
				Req: &grakn.Transaction_Req_Continue{
					Continue: true,
				},
				Metadata:      metadata,
				Id:            requestId,
				LatencyMillis: latencyMillis,
			})
			if err != nil {
				return nil, fmt.Errorf("could not send query request iterator: %w", err)
			}
		}  else if transactionResponse.GetDone()  {
			break
		} else {
			queryResponse := transactionResponse.GetQueryRes()
			if queryResponse != nil {
				matchResponse := queryResponse.GetMatchRes()
				matchResponses = append(matchResponses, matchResponse)
			}
		}
	}

	return matchResponses, err
}

func RunMatchGroupQuery(transactionClient grakn.Grakn_TransactionClient, requestId string, metadata map[string]string, query string, explain bool, batchSize int32, latencyMillis int32) (matchGroupResponses []*grakn.Query_MatchGroup_Res, err error) {
	if batchSize == 0 {
		batchSize = 2147483647
	}

	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_QueryReq{
			QueryReq: &grakn.Query_Req{
				Req: &grakn.Query_Req_MatchGroupReq{
					MatchGroupReq: &grakn.Query_MatchGroup_Req{
						Query: query,
					},
				},
				Options: &grakn.Options{
					BatchSizeOpt: &grakn.Options_BatchSize{
						BatchSize: batchSize,
					},
					ExplainOpt: &grakn.Options_Explain{
						Explain: explain,
					},
				},
			},
		},
		Metadata:      metadata,
		Id:            requestId,
		LatencyMillis: latencyMillis,
	})

	if err != nil {
		return nil, fmt.Errorf("could not send query request: %w", err)
	}

	for {
		transactionResponse, err := transactionClient.Recv()
		if err != nil {
			return nil, fmt.Errorf("could not receive query response: %w", err)
		}

		if transactionResponse.GetContinue() {
			err = transactionClient.Send(&grakn.Transaction_Req{
				Req: &grakn.Transaction_Req_Continue{
					Continue: true,
				},
				Metadata:      metadata,
				Id:            requestId,
				LatencyMillis: latencyMillis,
			})
			if err != nil {
				return nil, fmt.Errorf("could not send query request iterator: %w", err)
			}
		} else if transactionResponse.GetDone() {
			break
		} else {
			queryResponse := transactionResponse.GetQueryRes()
			if queryResponse != nil {
				matchGroupResponse := queryResponse.GetMatchGroupRes()
				matchGroupResponses = append(matchGroupResponses, matchGroupResponse)
			}
		}
	}
	return matchGroupResponses, err
}

func RunMatchGroupAggregateQuery(transactionClient grakn.Grakn_TransactionClient, requestId string, metadata map[string]string, query string, explain bool, batchSize int32, latencyMillis int32) (matchGroupAggregateResponses []*grakn.Query_MatchGroupAggregate_Res, err error) {
	if batchSize == 0 {
		batchSize = 2147483647
	}

	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_QueryReq{
			QueryReq: &grakn.Query_Req{
				Req: &grakn.Query_Req_MatchGroupAggregateReq{
					MatchGroupAggregateReq: &grakn.Query_MatchGroupAggregate_Req{
						Query: query,
					},
				},
				Options: &grakn.Options{
					BatchSizeOpt: &grakn.Options_BatchSize{
						BatchSize: batchSize,
					},
					ExplainOpt: &grakn.Options_Explain{
						Explain: explain,
					},
				},
			},
		},
		Metadata:      metadata,
		Id:            requestId,
		LatencyMillis: latencyMillis,
	})

	if err != nil {
		return nil, fmt.Errorf("could not send query request: %w", err)
	}

	for {
		transactionResponse, err := transactionClient.Recv()
		if err != nil {
			return nil, fmt.Errorf("could not receive query response: %w", err)
		}

		if transactionResponse.GetContinue() {
			err = transactionClient.Send(&grakn.Transaction_Req{
				Req: &grakn.Transaction_Req_Continue{
					Continue: true,
				},
				Metadata:      metadata,
				Id:            requestId,
				LatencyMillis: latencyMillis,
			})
			if err != nil {
				return nil, fmt.Errorf("could not send query request iterator: %w", err)
			}
		} else if transactionResponse.GetDone() {
			break
		} else {
			queryResponse := transactionResponse.GetQueryRes()
			if queryResponse != nil {
				matchGroupAggregateResponse := queryResponse.GetMatchGroupAggregateRes()
				matchGroupAggregateResponses = append(matchGroupAggregateResponses, matchGroupAggregateResponse)
			}
		}
	}
	return matchGroupAggregateResponses, err
}

func RunMatchAggregateQuery(transactionClient grakn.Grakn_TransactionClient, requestId string, metadata map[string]string, query string, explain bool, batchSize int32, latencyMillis int32) (matchAggregateResponses []*grakn.Query_MatchAggregate_Res, err error) {
	if batchSize == 0 {
		batchSize = 2147483647
	}

	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_QueryReq{
			QueryReq: &grakn.Query_Req{
				Req: &grakn.Query_Req_MatchAggregateReq{
					MatchAggregateReq: &grakn.Query_MatchAggregate_Req{
						Query: query,
					},
				},
				Options: &grakn.Options{
					BatchSizeOpt: &grakn.Options_BatchSize{
						BatchSize: batchSize,
					},
					ExplainOpt: &grakn.Options_Explain{
						Explain: explain,
					},
				},
			},
		},
		Metadata:      metadata,
		Id:            requestId,
		LatencyMillis: latencyMillis,
	})

	if err != nil {
		return nil, fmt.Errorf("could not send query request: %w", err)
	}

	transactionResponse, err := transactionClient.Recv()
	if err != nil {
		return nil, fmt.Errorf("could not receive query response: %w", err)
	}

	queryResponse := transactionResponse.GetQueryRes()
	if queryResponse != nil {
		matchAggregateResponse := queryResponse.GetMatchAggregateRes()
		matchAggregateResponses = append(matchAggregateResponses, matchAggregateResponse)
	}

	return matchAggregateResponses, err
}

func OpenTransaction(transactionClient grakn.Grakn_TransactionClient, sessionId []byte, transactionId string, transactionType grakn.Transaction_Type, metadata map[string]string, infer bool, latencyMillis int32) (err error) {
	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_OpenReq{
			OpenReq: &grakn.Transaction_Open_Req{
				SessionId: sessionId,
				Type:      transactionType,
				Options: &grakn.Options{
					InferOpt:&grakn.Options_Infer{
						Infer: infer,
					},
				},
			},
		},
		Metadata:      metadata,
		LatencyMillis: latencyMillis,
		Id:            transactionId,
	})
	if err != nil {
		return fmt.Errorf("could not send open transaction request: %w", err)
	}

	_, err = transactionClient.Recv()
	if err != nil {
		return fmt.Errorf("could not receive open transaction response: %w", err)
	}
	return err
}

func CommitTransaction(transactionClient grakn.Grakn_TransactionClient, transactionId string, metadata map[string]string, latencyMillis int32) (err error) {
	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_CommitReq{
			CommitReq: &grakn.Transaction_Commit_Req{
			},
		},
		Metadata:      metadata,
		Id:            transactionId,
		LatencyMillis: latencyMillis,
	})
	if err != nil {
		return fmt.Errorf("could not send commit transaction request: %w", err)
	}

	_, err = transactionClient.Recv()
	if err != nil {
		return fmt.Errorf("could not receive commit transaction response: %w", err)
	}
	return err
}

func RollbackTransaction(transactionClient grakn.Grakn_TransactionClient, transactionId string, metadata map[string]string, latencyMillis int32) (err error) {
	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_RollbackReq{
			RollbackReq: &grakn.Transaction_Rollback_Req{
			},
		},
		Metadata:      metadata,
		Id:            transactionId,
		LatencyMillis: latencyMillis,
	})
	if err != nil {
		return fmt.Errorf("could not send rollback transaction request: %w", err)
	}

	_, err = transactionClient.Recv()
	if err != nil {
		return fmt.Errorf("could not receive rollback transaction response: %w", err)
	}
	return err
}

func CloseTransaction(transactionClient grakn.Grakn_TransactionClient) (err error) {
	err = transactionClient.CloseSend()
	if err != nil {
		return fmt.Errorf("could not close send: %w", err)
	}
	return err
}
