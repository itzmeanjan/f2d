package events

import (
	"context"
	"log"

	"github.com/gorilla/websocket"
	"github.com/itzmeanjan/f2d/app/config"
	"github.com/itzmeanjan/f2d/app/data"
)

// Subscribe - Subscribe to event emitted by `ette`, for listening to
// new block getting mined/ event logs getting emitted due to contract interaction
// and attempt to process those events
func Subscribe(ctx context.Context) bool {

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, config.GetEtteWSURL(), nil)
	if err != nil {

		log.Printf("[❗️] Failed to connect to `ette`: %s\n", err.Error())
		return false

	}

	if err := conn.WriteJSON(&data.EtteSubscriptionRequest{
		Name:   "block",
		Type:   "subscription",
		APIKey: config.GetEtteAPIKey(),
	}); err != nil {

		log.Printf("[❗️] Failed to send subscription request to `ette`: %s\n", err.Error())
		return false

	}

	go func() {

		for {
			select {

			case <-ctx.Done():

				log.Printf("[➕] Shutting down listener\n")
				break

			default:

				var block data.Block

				if err := conn.ReadJSON(&block); err != nil {

					log.Printf("[❗️] Failed to read event from `ette` : %s\n", err.Error())
					break

				}

				// @note Do process this block

			}
		}

	}()

	return true

}
