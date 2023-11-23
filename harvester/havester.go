package harvester

import (
	"context"
	log "github.com/sirupsen/logrus"
	"league-havester/connector"
	"league-havester/finder"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Option struct {
	HarvesterDuration time.Duration
}

type Harvester interface {
	Start()
}

type harvester struct {
	ctx             context.Context
	harvesterOption Option
	finderOption    finder.Option
}

func (h *harvester) Start() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-signals
		log.Println(sig)
		done <- true
	}()

	// business logic
	ticker := time.NewTicker(h.harvesterOption.HarvesterDuration)
	go func() {
		for range ticker.C {
			if result, err := finder.New(h.ctx, h.finderOption).Execute(); err != nil {
				log.Println(err)
			} else {
				log.Printf("%+v", result)
				connector.New(h.ctx, connector.Config(result)).Connect()
			}
		}
	}()
	<-done
	log.Println("terminated")
}

func New(harvesterOption Option, finderOption finder.Option) Harvester {
	return &harvester{
		ctx:             context.Background(),
		harvesterOption: harvesterOption,
		finderOption:    finderOption,
	}
}
