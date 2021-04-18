package data

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/itzmeanjan/f2d/app/config"
	"github.com/itzmeanjan/f2d/app/db"
	"gorm.io/gorm"
)

// RPCClients - Two clients connected to RPC node
// over both HTTP & WS transport
type RPCClients struct {
	HTTP, WS *ethclient.Client
}

// Resources - File, database, network resources which are to be accessed
// from several go routines, fulfilling different purposes, to be kept/ passed
// along using this struct
type Resources struct {
	DB  *gorm.DB
	RPC *RPCClients
}

// Acquire - During application start up, acquire all resources
func Acquire(ctx context.Context) (*Resources, error) {

	var resource *Resources = new(Resources)
	var clients *RPCClients = new(RPCClients)

	handle, err := db.Connect()
	if err != nil {
		return nil, err
	}

	resource.DB = handle
	httpClient, err := ethclient.DialContext(ctx, config.GetRPCHTTPURL())
	if err != nil {
		return nil, err
	}

	clients.HTTP = httpClient
	wsClient, err := ethclient.DialContext(ctx, config.GetRPCWSURL())
	if err != nil {
		return nil, err
	}

	clients.WS = wsClient
	resource.RPC = clients
	return resource, nil

}

// Release - Before shutting down application, release
// resources in a graceful manner
func (r *Resources) Release() error {

	sql, err := r.DB.DB()
	if err != nil {
		return err
	}

	if err := sql.Close(); err != nil {
		return err
	}

	r.RPC.HTTP.Close()
	r.RPC.WS.Close()

	return nil

}
