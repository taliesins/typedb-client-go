package concept_manager

import "github.com/taliesins/typedb-client-go/v2/typedb_protocol"

func ConceptManagerReq(req *typedb_protocol.ConceptManager_Req) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: req,
		},
	}
}

func PutEntityTypeReq(label string) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &typedb_protocol.ConceptManager_Req{
				Req: &typedb_protocol.ConceptManager_Req_PutEntityTypeReq{
					PutEntityTypeReq: &typedb_protocol.ConceptManager_PutEntityType_Req{
						Label: label,
					},
				},
			},
		},
	}
}

func PutRelationTypeReq(label string) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &typedb_protocol.ConceptManager_Req{
				Req: &typedb_protocol.ConceptManager_Req_PutRelationTypeReq{
					PutRelationTypeReq: &typedb_protocol.ConceptManager_PutRelationType_Req{
						Label: label,
					},
				},
			},
		},
	}
}

func PutAttributeTypeReq(label string, valueType typedb_protocol.AttributeType_ValueType) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &typedb_protocol.ConceptManager_Req{
				Req: &typedb_protocol.ConceptManager_Req_PutAttributeTypeReq{
					PutAttributeTypeReq: &typedb_protocol.ConceptManager_PutAttributeType_Req{
						Label: label,
						ValueType: valueType,
					},
				},
			},
		},
	}
}

func GetThingTypeReq(label string) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &typedb_protocol.ConceptManager_Req{
				Req: &typedb_protocol.ConceptManager_Req_GetThingTypeReq{
					GetThingTypeReq: &typedb_protocol.ConceptManager_GetThingType_Req{
						Label: label,
					},
				},
			},
		},
	}
}

func GetThingReq(iid []byte) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: &typedb_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &typedb_protocol.ConceptManager_Req{
				Req: &typedb_protocol.ConceptManager_Req_GetThingReq{
					GetThingReq: &typedb_protocol.ConceptManager_GetThing_Req{
						Iid: iid,
					},
				},
			},
		},
	}
}
