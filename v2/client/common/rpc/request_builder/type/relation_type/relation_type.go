package relation_type

import (
	"github.com/taliesins/typedb-client-go/v2/client/common"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

func newTransactionRequest(req *typedb_protocol.Transaction_Req_TypeReq) *typedb_protocol.Transaction_Req {
	return &typedb_protocol.Transaction_Req{
		Req: req,
	}
}

func newTypeRequest(label common.Label) *typedb_protocol.Type_Req {
	builder := &typedb_protocol.Type_Req{
		Label: label.Name,
		Scope: label.Scope,
	}
	return builder
}

func setRelationTypeCreateReq(typeReq *typedb_protocol.Type_Req, relationTypeCreateReq *typedb_protocol.RelationType_Create_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_RelationTypeCreateReq{
		RelationTypeCreateReq: relationTypeCreateReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func CreateReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeCreateReq(newTypeRequest(label), &typedb_protocol.RelationType_Create_Req{}))
}

func setRelationTypeGetRelatesReq(typeReq *typedb_protocol.Type_Req, relationTypeGetRelatesReq *typedb_protocol.RelationType_GetRelates_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_RelationTypeGetRelatesReq{
		RelationTypeGetRelatesReq: relationTypeGetRelatesReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetRelatesReq(label common.Label) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeGetRelatesReq(newTypeRequest(label), &typedb_protocol.RelationType_GetRelates_Req{}))
}

func setRelationTypeGetRelatesForRoleLabelReq(typeReq *typedb_protocol.Type_Req, relationTypeGetRelatesForRoleLabelReq *typedb_protocol.RelationType_GetRelatesForRoleLabel_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_RelationTypeGetRelatesForRoleLabelReq{
		RelationTypeGetRelatesForRoleLabelReq: relationTypeGetRelatesForRoleLabelReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func GetRelatesByRoleReq(label common.Label, roleLabel string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeGetRelatesForRoleLabelReq(newTypeRequest(label), &typedb_protocol.RelationType_GetRelatesForRoleLabel_Req{Label: roleLabel}))
}

func setRelationTypeSetRelatesReq(typeReq *typedb_protocol.Type_Req, relationTypeSetRelatesReq *typedb_protocol.RelationType_SetRelates_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_RelationTypeSetRelatesReq{
		RelationTypeSetRelatesReq: relationTypeSetRelatesReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func SetRelatesReq(label common.Label, roleLabel string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeSetRelatesReq(newTypeRequest(label), &typedb_protocol.RelationType_SetRelates_Req{Label: roleLabel}))
}

func SetRelatesOverriddenReq(label common.Label, roleLabel string, overriddenLabel string) *typedb_protocol.Transaction_Req {
return newTransactionRequest(setRelationTypeSetRelatesReq(newTypeRequest(label),
	&typedb_protocol.RelationType_SetRelates_Req{
		Label: roleLabel,
		Overridden: &typedb_protocol.RelationType_SetRelates_Req_OverriddenLabel{
			OverriddenLabel: overriddenLabel,
		},
	}))
}

func setRelationTypeUnsetRelatesReq(typeReq *typedb_protocol.Type_Req, relationTypeUnsetRelatesReq *typedb_protocol.RelationType_UnsetRelates_Req) *typedb_protocol.Transaction_Req_TypeReq {
	typeReq.Req = &typedb_protocol.Type_Req_RelationTypeUnsetRelatesReq{
		RelationTypeUnsetRelatesReq: relationTypeUnsetRelatesReq,
	}

	return &typedb_protocol.Transaction_Req_TypeReq{
		TypeReq: typeReq,
	}
}

func UnsetRelatesReq(label common.Label, roleLabel string) *typedb_protocol.Transaction_Req {
	return newTransactionRequest(setRelationTypeUnsetRelatesReq(newTypeRequest(label),
		&typedb_protocol.RelationType_UnsetRelates_Req{
			Label: roleLabel,
		}))
}
