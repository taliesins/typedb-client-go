package concept_manager

import "github.com/taliesins/typedb-client-go/v2/grakn_protocol"

func PutEntityTypeReq(label string) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &grakn_protocol.ConceptManager_Req{
				Req: &grakn_protocol.ConceptManager_Req_PutEntityTypeReq{
					PutEntityTypeReq: &grakn_protocol.ConceptManager_PutEntityType_Req{
						Label: label,
					},
				},
			},
		},
	}
}

func PutRelationTypeReq(label string) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &grakn_protocol.ConceptManager_Req{
				Req: &grakn_protocol.ConceptManager_Req_PutRelationTypeReq{
					PutRelationTypeReq: &grakn_protocol.ConceptManager_PutRelationType_Req{
						Label: label,
					},
				},
			},
		},
	}
}

func PutAttributeTypeReq(label string) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &grakn_protocol.ConceptManager_Req{
				Req: &grakn_protocol.ConceptManager_Req_PutAttributeTypeReq{
					PutAttributeTypeReq: &grakn_protocol.ConceptManager_PutAttributeType_Req{
						Label: label,
					},
				},
			},
		},
	}
}

func GetThingTypeReq(label string) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &grakn_protocol.ConceptManager_Req{
				Req: &grakn_protocol.ConceptManager_Req_GetThingTypeReq{
					GetThingTypeReq: &grakn_protocol.ConceptManager_GetThingType_Req{
						Label: label,
					},
				},
			},
		},
	}
}

func GetThingReq(iid []byte) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: &grakn_protocol.Transaction_Req_ConceptManagerReq{
			ConceptManagerReq: &grakn_protocol.ConceptManager_Req{
				Req: &grakn_protocol.ConceptManager_Req_GetThingReq{
					GetThingReq: &grakn_protocol.ConceptManager_GetThing_Req{
						Iid: iid,
					},
				},
			},
		},
	}
}
