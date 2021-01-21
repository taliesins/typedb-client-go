package client

import (
	"fmt"
	grakn "github.com/taliesins/client-go/v1/grakn_protocol"
	"log"
)

func GetById(transactionClient grakn.SessionService_TransactionClient, metadata map[string]string, id string) (concept *grakn.Concept, err error){
	if id != "" {
		err = transactionClient.Send(&grakn.Transaction_Req{
			Req: &grakn.Transaction_Req_GetConceptReq{
				GetConceptReq: &grakn.Transaction_GetConcept_Req{
					Id: id,
				},
			},
			Metadata: metadata,
		})

		if err != nil {
			return nil, fmt.Errorf("could not send get concept request: %w", err)
		}

		transactionResponse, err := transactionClient.Recv()
		if err != nil {
			return nil, fmt.Errorf("could not receive get concept response: %w", err)
		}

		getConceptResponse := transactionResponse.GetGetConceptRes()

		if getConceptResponse == nil {
			return nil, nil
		}

		concept = getConceptResponse.GetConcept()
		return concept, err
	}

	return nil, nil
}

func GetAttributesByValue(transactionClient grakn.SessionService_TransactionClient, metadata map[string]string, batchSize int32, value *grakn.ValueObject) (concepts []*grakn.Concept, err error){
	var options grakn.Transaction_Iter_Req_Options
	if batchSize == 0 {
		options = grakn.Transaction_Iter_Req_Options{
			BatchSize: &grakn.Transaction_Iter_Req_Options_All{
				All: true,
			},
		}
	} else {
		options = grakn.Transaction_Iter_Req_Options{
			BatchSize: &grakn.Transaction_Iter_Req_Options_Number{
				Number: batchSize,
			},
		}
	}

	if value != nil {
		err = transactionClient.Send(&grakn.Transaction_Req{
			Req: &grakn.Transaction_Req_IterReq{
				IterReq: &grakn.Transaction_Iter_Req{
					Req: &grakn.Transaction_Iter_Req_GetAttributesIterReq{
						GetAttributesIterReq: &grakn.Transaction_GetAttributes_Iter_Req{
							Value: value,
						},
					},
				},
			},
			Metadata: metadata,
		})

		if err != nil {
			return nil, fmt.Errorf("could not send get attribute by value request: %w", err)
		}

		for {
			transactionResponse, err := transactionClient.Recv()
			if err != nil {
				return nil, fmt.Errorf("could not receive get attribute by value response: %w", err)
			}

			queryIteratorResponse := transactionResponse.GetIterRes()
			if queryIteratorResponse.GetDone() {
				break
			}

			getAttributesResponse := queryIteratorResponse.GetGetAttributesIterRes()

			if getAttributesResponse != nil {
				attribute := getAttributesResponse.GetAttribute()
				concepts = append(concepts, attribute)
			} else {
				//Batch size hit
				iteratorId := queryIteratorResponse.GetIteratorId()
				if iteratorId > 0 {
					log.Printf("iteratorId: %v", iteratorId)

					err = transactionClient.Send(&grakn.Transaction_Req{
						Req: &grakn.Transaction_Req_IterReq{
							IterReq: &grakn.Transaction_Iter_Req{
								Req: &grakn.Transaction_Iter_Req_IteratorId{
									IteratorId: iteratorId,
								},
								Options: &options,
							},
						},
						Metadata: metadata,
					})

					if err != nil {
						return nil, fmt.Errorf("could not send get attribute request iterator: %w", err)
					}
				}
			}
		}
	}

	return concepts, err
}

func RunQuery(transactionClient grakn.SessionService_TransactionClient, metadata map[string]string, query string, infer bool, explain bool, batchSize int32) (answers []*grakn.Answer, err error) {
	var options grakn.Transaction_Iter_Req_Options
	if batchSize == 0 {
		options = grakn.Transaction_Iter_Req_Options{
			BatchSize: &grakn.Transaction_Iter_Req_Options_All{
				All: true,
			},
		}
	} else {
		options = grakn.Transaction_Iter_Req_Options{
			BatchSize: &grakn.Transaction_Iter_Req_Options_Number{
				Number: batchSize,
			},
		}
	}

	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_IterReq{
			IterReq: &grakn.Transaction_Iter_Req{
				Req: &grakn.Transaction_Iter_Req_QueryIterReq{
					QueryIterReq: &grakn.Transaction_Query_Iter_Req{
						Query: query,
						Options: &grakn.Transaction_Query_Options{
							Infer: &grakn.Transaction_Query_Options_InferFlag{
								InferFlag: infer,
							},
							Explain: &grakn.Transaction_Query_Options_ExplainFlag{
								ExplainFlag: explain,
							},
						},
					},
				},
			},
		},
		Metadata: metadata,
	})

	if err != nil {
		return nil, fmt.Errorf("could not send query request: %w", err)
	}

	for {
		transactionResponse, err := transactionClient.Recv()
		if err != nil {
			return nil, fmt.Errorf("could not receive query response: %w", err)
		}

		queryIteratorResponse := transactionResponse.GetIterRes()
		if queryIteratorResponse.GetDone() {
			break
		}

		queryResponse := queryIteratorResponse.GetQueryIterRes()

		if queryResponse != nil {
			answer := queryResponse.GetAnswer()
			answers = append(answers, answer)
		} else {
			//Batch size hit
			iteratorId := queryIteratorResponse.GetIteratorId()
			if iteratorId > 0 {
				log.Printf("iteratorId: %v", iteratorId)

				err = transactionClient.Send(&grakn.Transaction_Req{
					Req: &grakn.Transaction_Req_IterReq{
						IterReq: &grakn.Transaction_Iter_Req{
							Req: &grakn.Transaction_Iter_Req_IteratorId{
								IteratorId: iteratorId,
							},
							Options: &options,
						},
					},
					Metadata: metadata,
				})

				if err != nil {
					return nil, fmt.Errorf("could not send query request iterator: %w", err)
				}
			}
		}
	}

	return answers, err
}

func OpenTransaction(transactionClient grakn.SessionService_TransactionClient, sessionId string, transactionType grakn.Transaction_Type, metadata map[string]string) (err error) {
	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_OpenReq{
			OpenReq: &grakn.Transaction_Open_Req{
				SessionId: sessionId,
				Type:      transactionType,
			},
		},
		Metadata: metadata,
	})

	_, err = transactionClient.Recv()
	if err != nil {
		return fmt.Errorf("could not receive transaction response: %w", err)
	}
	return err
}

func CommitTransaction(transactionClient grakn.SessionService_TransactionClient, metadata map[string]string) (err error) {
	err = transactionClient.Send(&grakn.Transaction_Req{
		Req: &grakn.Transaction_Req_CommitReq{
			CommitReq: &grakn.Transaction_Commit_Req{
			},
		},
		Metadata: metadata,
	})

	_, err = transactionClient.Recv()
	if err != nil {
		return fmt.Errorf("could not receive transaction response: %w", err)
	}
	return err
}

func CloseTransaction(transactionClient grakn.SessionService_TransactionClient) (err error) {
	err = transactionClient.CloseSend()
	if err != nil {
		return fmt.Errorf("could not close send: %w", err)
	}
	return err
}
