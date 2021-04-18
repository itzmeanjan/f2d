package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/itzmeanjan/f2d/app"
	"github.com/itzmeanjan/f2d/app/events"
)

func main() {
	log.Printf("Firebase for DApps 🔥\n")

	ctx, cancel := context.WithCancel(context.TODO())
	res, err := app.SetUp(ctx)
	if err != nil {

		log.Printf("[❗️] Shutting down `f2d` : %s\n", err.Error())
		os.Exit(1)

	}

	status, comm := events.Subscribe(ctx)
	if !status {

		log.Printf("[❗️] Shutting down `f2d`\n")
		os.Exit(1)

	}

	// Attempt to catch interrupt event(s)
	// so that graceful shutdown can be performed
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)

	go func() {

		// To be invoked when returning from this
		// go rountine's execution scope
		defer func() {

			if err := res.Release(); err != nil {
				log.Printf("[❗️] Graceful resource release failed : %s\n", err.Error())
				os.Exit(1)
			}

			// Stopping process, gracefully
			log.Printf("[✅] Gracefully shut down `f2d`\n")
			os.Exit(0)

		}()

	OUTER:
		for {

			select {

			case <-interruptChan:

				// When interruption is received, attempting to
				// let all other go routines know, master go routine
				// wants all to shut down, they must do a graceful stop
				// of what they're doing now
				cancel()

				// Giving workers 3 seconds, before forcing shutdown
				<-time.After(time.Second * time.Duration(3))
				break OUTER

			case <-comm:

				// As soon as it's received that `ette` event listener
				// has been stopped, `f2d` will attempt to do a graceful shutdown
				//
				// @note This can be handled better by spwaning new go routine
				// which will attempt to subscribe to `ette`
				//
				// But also need to take care of circuit breaking in that case
				log.Printf("[❗️] `ette` event listener stopped\n")

				// Asking all go routines to stop
				cancel()
				break OUTER

			}

		}

	}()

	c := make(chan struct{})
	<-c

}
