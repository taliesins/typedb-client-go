package concept

import (
	"context"
	"github.com/taliesins/typedb-client-go/v2/client/common"
	"github.com/taliesins/typedb-client-go/v2/client/common/rpc/request_builder/concept_manager"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

type ConceptManager interface {
	GetRootThingType(metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<ThingType>;
	GetRootEntityType(metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<EntityType>;
	GetRootRelationType(metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<RelationType>;
	GetRootAttributeType(metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<AttributeType>;
	GetThingType(label string, metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<ThingType>;
	GetThing(iid string, metadata map[string]string) (*typedb_protocol.Thing, error)//: Promise<Thing>;
	GetEntityType(label string, metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<EntityType>;
	PutEntityType(label string, metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<EntityType>;
	GetRelationType(label string, metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<RelationType>;
	PutRelationType(label string, metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<RelationType>;
	GetAttributeType(label string, metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<AttributeType>;
	PutAttributeType(label string, valueType typedb_protocol.AttributeType_ValueType, metadata map[string]string) (*typedb_protocol.Type, error)//: Promise<AttributeType>;
}

type conceptManagerImpl struct {
	TransactionClient typedb_protocol.TypeDB_TransactionClient
	Context context.Context
}

func NewConceptManager(transactionClient typedb_protocol.TypeDB_TransactionClient, ctx context.Context) *conceptManagerImpl {
	return &conceptManagerImpl{
		TransactionClient: transactionClient,
		Context: ctx,
	}
}

func (c *conceptManagerImpl) GetRootThingType(metadata map[string]string) (*typedb_protocol.Type, error) {
	return c.GetThingType("thing", metadata)
}

func (c *conceptManagerImpl) GetRootEntityType(metadata map[string]string) (*typedb_protocol.Type, error) {
	return c.GetEntityType("entity", metadata)
}

func (c *conceptManagerImpl) GetRootRelationType(metadata map[string]string) (*typedb_protocol.Type, error) {
	return c.GetRelationType("relation", metadata)
}

func (c *conceptManagerImpl) GetRootAttributeType(metadata map[string]string) (*typedb_protocol.Type, error) {
	return c.GetAttributeType("attribute", metadata)
}

func (c *conceptManagerImpl) GetThingType(label string, metadata map[string]string) (*typedb_protocol.Type, error) {
	request := concept_manager.GetThingTypeReq(label)
	request.Metadata = metadata

	response, err := common.NewPromise(c.TransactionClient, c.Context, request)
	if err != nil {
		return nil, err
	}

	var thingType *typedb_protocol.Type = nil

	for {
		select {
		case <-response.Ctx.Done():
			return thingType, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}

			thingType = data.GetConceptManagerRes().GetGetThingTypeRes().GetThingType()
		}
	}
}

func (c *conceptManagerImpl) GetThing(iid []byte, metadata map[string]string) (*typedb_protocol.Thing, error) {
	request := concept_manager.GetThingReq(iid)
	request.Metadata = metadata

	response, err := common.NewPromise(c.TransactionClient, c.Context, request)
	if err != nil {
		return nil, err
	}

	var thing *typedb_protocol.Thing = nil

	for {
		select {
		case <-response.Ctx.Done():
			return thing, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}

			thing = data.GetConceptManagerRes().GetGetThingRes().GetThing()
		}
	}
}

func (c *conceptManagerImpl) GetEntityType(label string, metadata map[string]string) (*typedb_protocol.Type, error) {
	thingType, err := c.GetThingType(label, metadata)
	if err != nil {
		return nil, err
	}

	if thingType.GetEncoding() != typedb_protocol.Type_ENTITY_TYPE {
		return nil, nil
	}

	entityType := thingType

	return entityType, nil
}

func (c *conceptManagerImpl) PutEntityType(label string, metadata map[string]string) (*typedb_protocol.Type, error) {
	request := concept_manager.PutEntityTypeReq(label)
	request.Metadata = metadata

	response, err := common.NewPromise(c.TransactionClient, c.Context, request)
	if err != nil {
		return nil, err
	}

	var entityType *typedb_protocol.Type = nil

	for {
		select {
		case <-response.Ctx.Done():
			return entityType, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}

			entityType = data.GetConceptManagerRes().GetPutEntityTypeRes().GetEntityType()
		}
	}
}

func (c *conceptManagerImpl) GetRelationType(label string, metadata map[string]string) (*typedb_protocol.Type, error) {
	thingType, err := c.GetThingType(label, metadata)
	if err != nil {
		return nil, err
	}

	if thingType.GetEncoding() != typedb_protocol.Type_RELATION_TYPE {
		return nil, nil
	}

	relationType := thingType

	return relationType, nil
}

func (c *conceptManagerImpl) PutRelationType(label string, metadata map[string]string) (*typedb_protocol.Type, error) {
	request := concept_manager.PutRelationTypeReq(label)
	request.Metadata = metadata

	response, err := common.NewPromise(c.TransactionClient, c.Context, request)
	if err != nil {
		return nil, err
	}

	var relationType *typedb_protocol.Type = nil

	for {
		select {
		case <-response.Ctx.Done():
			return relationType, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}

			relationType = data.GetConceptManagerRes().GetPutRelationTypeRes().GetRelationType()
		}
	}
}

func (c *conceptManagerImpl) GetAttributeType(label string, metadata map[string]string) (*typedb_protocol.Type, error) {
	thingType, err := c.GetThingType(label, metadata)
	if err != nil {
		return nil, err
	}

	if thingType.GetEncoding() != typedb_protocol.Type_ATTRIBUTE_TYPE {
		return nil, nil
	}

	attributeType := thingType

	return attributeType, nil
}

func (c *conceptManagerImpl) PutAttributeType(label string, valueType typedb_protocol.AttributeType_ValueType, metadata map[string]string) (*typedb_protocol.Type, error) {
	request := concept_manager.PutAttributeTypeReq(label, valueType)
	request.Metadata = metadata

	response, err := common.NewPromise(c.TransactionClient, c.Context, request)
	if err != nil {
		return nil, err
	}

	var attributeType *typedb_protocol.Type = nil

	for {
		select {
		case <-response.Ctx.Done():
			return attributeType, response.Ctx.Err()
		case data, ok := <-response.Data:
			if !ok {
				return nil, response.Err
			}

			attributeType = data.GetConceptManagerRes().GetPutAttributeTypeRes().GetAttributeType()
		}
	}
}
