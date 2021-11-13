package user

import (
	"context"
	"github.com/taliesins/typedb-client-go/v2/client/common/rpc/request_builder/cluster/user"
	"github.com/taliesins/typedb-client-go/v2/typedb_protocol"
)

type ClusterUserManager interface {
	//Get(username string, metadata map[string]string) (*typedb_protocol.ClusterUser, error)
	Contains(username string, metadata map[string]string) (bool, error)
	Create(username string, password string) error
	Delete(username string) error
	ChangePassword(username string, password string) error
	All() ([]*string, error)
}

type clusterUserManagerImpl struct {
	TransactionClient typedb_protocol.TypeDB_TransactionClient
	Context           context.Context
}

func NewClusterUserManager(transactionClient typedb_protocol.TypeDB_TransactionClient, ctx context.Context) *clusterUserManagerImpl {
	return &clusterUserManagerImpl{
		TransactionClient: transactionClient,
		Context: ctx,
	}
}

func (c *clusterUserManagerImpl) Contains(username string, metadata map[string]string)(bool, error) {

}

func (c *clusterUserManagerImpl) Create(username string, password string)(error) {

}

func (c *clusterUserManagerImpl) Delete(username string)(error) {
	request := user.DeleteReq(username)
}

func (c *clusterUserManagerImpl) ChangePassword(username string, password string)(error) {
	request := user.PasswordReq(username, password)
}

func (c *clusterUserManagerImpl) All()([]*string, error) {

}
