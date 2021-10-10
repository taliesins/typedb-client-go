package query_manager

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func DefineReq(query string, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_DefineReq{
					DefineReq: &typedb_protocol.QueryManager_Define_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func UndefineReq(query string, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_UndefineReq{
					UndefineReq: &typedb_protocol.QueryManager_Undefine_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func MatchReq(query string, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_MatchReq{
					MatchReq: &typedb_protocol.QueryManager_Match_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func MatchAggregateReq(query string, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_MatchAggregateReq{
					MatchAggregateReq: &typedb_protocol.QueryManager_MatchAggregate_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func MatchGroupReq(query string, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_MatchGroupReq{
					MatchGroupReq: &typedb_protocol.QueryManager_MatchGroup_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func MatchGroupAggregateReq(query string, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_MatchGroupAggregateReq{
					MatchGroupAggregateReq: &typedb_protocol.QueryManager_MatchGroupAggregate_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func InsertReq(query string, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_InsertReq{
					InsertReq: &typedb_protocol.QueryManager_Insert_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func DeleteReq(query string, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_DeleteReq{
					DeleteReq: &typedb_protocol.QueryManager_Delete_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func UpdateReq(query string, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_UpdateReq{
					UpdateReq: &typedb_protocol.QueryManager_Update_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func ExplainReq(id int64, options *typedb_protocol.Options) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &typedb_protocol.QueryManager_Req{
				Req: &typedb_protocol.QueryManager_Req_ExplainReq{
					ExplainReq: &typedb_protocol.QueryManager_Explain_Req{
						ExplainableId: id,
					},
				},
				Options: options,
			},
		},
	}
}
