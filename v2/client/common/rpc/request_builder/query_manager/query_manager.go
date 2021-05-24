package query_manager

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func DefineReq(query string, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_DefineReq{
					DefineReq: &grakn_protocol.QueryManager_Define_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func UndefineReq(query string, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_UndefineReq{
					UndefineReq: &grakn_protocol.QueryManager_Undefine_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func MatchReq(query string, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_MatchReq{
					MatchReq: &grakn_protocol.QueryManager_Match_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func MatchAggregateReq(query string, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_MatchAggregateReq{
					MatchAggregateReq: &grakn_protocol.QueryManager_MatchAggregate_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func MatchGroupReq(query string, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_MatchGroupReq{
					MatchGroupReq: &grakn_protocol.QueryManager_MatchGroup_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func MatchGroupAggregateReq(query string, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_MatchGroupAggregateReq{
					MatchGroupAggregateReq: &grakn_protocol.QueryManager_MatchGroupAggregate_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func InsertReq(query string, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_InsertReq{
					InsertReq: &grakn_protocol.QueryManager_Insert_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func DeleteReq(query string, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_DeleteReq{
					DeleteReq: &grakn_protocol.QueryManager_Delete_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func UpdateReq(query string, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_UpdateReq{
					UpdateReq: &grakn_protocol.QueryManager_Update_Req{
						Query: query,
					},
				},
				Options: options,
			},
		},
	}
}

func ExplainReq(id int64, options *grakn_protocol.Options) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_QueryManagerReq{
			QueryManagerReq: &grakn_protocol.QueryManager_Req{
				Req: &grakn_protocol.QueryManager_Req_ExplainReq{
					ExplainReq: &grakn_protocol.QueryManager_Explain_Req{
						ExplainableId: id,
					},
				},
				Options: options,
			},
		},
	}
}
