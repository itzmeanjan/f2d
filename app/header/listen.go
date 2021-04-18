package header

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Subscribe - Subscribe for listening to block headers & process block
func Subscribe(ctx context.Context, rpc *ethclient.Client, closeChan chan struct{}) {

	headerChan := make(chan *types.Header, 1)
	sub, err := rpc.SubscribeNewHead(ctx, headerChan)
	if err != nil {
		log.Printf("[❗️] Failed to subscribe to block headers : %s\n", err.Error())

		close(closeChan)
		return
	}

	for {

		select {

		case <-ctx.Done():
			sub.Unsubscribe()
			return

		case err := <-sub.Err():
			log.Printf("[❗️] Header subscription canceled : %s\n", err.Error())

			close(closeChan)
			return

		case header := <-headerChan:
			log.Printf("[+] Received block header : %d\n", header.Number.Uint64())

		}

	}

}
