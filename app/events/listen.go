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
func Subscribe(ctx context.Context) (bool, chan struct{}) {

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, config.GetEtteWSURL(), nil)
	if err != nil {

		log.Printf("[❗️] Failed to connect to `ette`: %s\n", err.Error())
		return false, nil

	}

	if err := conn.WriteJSON(&data.EtteSubscriptionRequest{
		Name:   "block",
		Type:   "subscription",
		APIKey: config.GetEtteAPIKey(),
	}); err != nil {

		log.Printf("[❗️] Failed to send subscription request to `ette`: %s\n", err.Error())
		return false, nil

	}

	comm := make(chan struct{})

	go func() {

		var first bool = true

		for {
			select {

			case <-ctx.Done():

				log.Printf("[➕] Shutting down listener\n")
				break

			default:

				// After subscription request is sent to `ette`,
				// we're expecting to receive subscription confirmation message
				//
				// Once subscription is confirmed, it'll now expect block(s), as soon
				// as they get mined
				if first {

					var confirmation data.EtteSubscriptionResponse

					if err := conn.ReadJSON(&confirmation); err != nil {

						log.Printf("[❗️] Failed to receive confirmation from `ette` : %s\n", err.Error())

						// Closing communication channel to let
						// supervisor know, this worker has failed
						close(comm)

						break

					}

					if confirmation.Code == 0 {

						log.Printf("[❗️] Failed to subscribe to `ette` events\n")

						// Closing communication channel to let
						// supervisor know, this worker has failed
						close(comm)

						break

					}

					first = !first

				}

				var block data.EtteBlock

				if err := conn.ReadJSON(&block); err != nil {

					log.Printf("[❗️] Failed to read event from `ette` : %s\n", err.Error())

					// Closing communication channel to let
					// supervisor know, this worker has failed
					close(comm)

					break

				}

				// @note Do process this block

			}
		}

	}()

	return true, comm

}
