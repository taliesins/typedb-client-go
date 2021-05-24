package relation_type

import (
	"github.com/taliesins/typedb-client-go/v2/client/common"
	"github.com/taliesins/typedb-client-go/v2/grakn_protocol"
)

func newTransactionRequest(req *grakn_protocol.Transaction_Req_TypeReq) *grakn_protocol.Transaction_Req {
	return &grakn_protocol.Transaction_Req{
		Req: req,
	}
}

func newTypeRequest(label common.Label) *grakn_protocol.Type_Req {
	builder := &grakn_protocol.Type_Req{
		Label: label.Name,
		Scope: label.Scope,
	}
	return builder
}

func setRelationTypeCreateReq(typeReq *grakn_protocol.Type_Req, relationTypeCreateReq *grakn_protocol.RelationType_Create_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_RelationTypeCreateReq{
		RelationTypeCreateReq: relationTypeCreateReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func CreateReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeCreateReq(newTypeRequest(label), &grakn_protocol.RelationType_Create_Req{}))
}

func setRelationTypeGetRelatesReq(typeReq *grakn_protocol.Type_Req, relationTypeGetRelatesReq *grakn_protocol.RelationType_GetRelates_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_RelationTypeGetRelatesReq{
		RelationTypeGetRelatesReq: relationTypeGetRelatesReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetRelatesReq(label common.Label) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeGetRelatesReq(newTypeRequest(label), &grakn_protocol.RelationType_GetRelates_Req{}))
}

func setRelationTypeGetRelatesForRoleLabelReq(typeReq *grakn_protocol.Type_Req, relationTypeGetRelatesForRoleLabelReq *grakn_protocol.RelationType_GetRelatesForRoleLabel_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_RelationTypeGetRelatesForRoleLabelReq{
		RelationTypeGetRelatesForRoleLabelReq: relationTypeGetRelatesForRoleLabelReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetRelatesByRoleReq(label common.Label, roleLabel string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeGetRelatesForRoleLabelReq(newTypeRequest(label), &grakn_protocol.RelationType_GetRelatesForRoleLabel_Req{Label: roleLabel}))
}

func setRelationTypeSetRelatesReq(typeReq *grakn_protocol.Type_Req, relationTypeSetRelatesReq *grakn_protocol.RelationType_SetRelates_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_RelationTypeSetRelatesReq{
		RelationTypeSetRelatesReq: relationTypeSetRelatesReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetRelatesReq(label common.Label, roleLabel string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeSetRelatesReq(newTypeRequest(label), &grakn_protocol.RelationType_SetRelates_Req{Label: roleLabel}))
}

func SetRelatesOverriddenReq(label common.Label, roleLabel string, overriddenLabel string) *grakn_protocol.Transaction_Req {
return newTransactionRequest(setRelationTypeSetRelatesReq(newTypeRequest(label),
	&grakn_protocol.RelationType_SetRelates_Req{
		Label: roleLabel,
		Overridden: &grakn_protocol.RelationType_SetRelates_Req_OverriddenLabel{
			OverriddenLabel: overriddenLabel,
		},
	}))
}

func setRelationTypeUnsetRelatesReq(typeReq *grakn_protocol.Type_Req, relationTypeUnsetRelatesReq *grakn_protocol.RelationType_UnsetRelates_Req) *grakn_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &grakn_protocol.Type_Req_RelationTypeUnsetRelatesReq{
		RelationTypeUnsetRelatesReq: relationTypeUnsetRelatesReq,
	}

	return &grakn_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func UnsetRelatesReq(label common.Label, roleLabel string) *grakn_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeUnsetRelatesReq(newTypeRequest(label),
		&grakn_protocol.RelationType_UnsetRelates_Req{
			Label: roleLabel,
		}))
}
