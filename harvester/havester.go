package harvester

import (
	"context"
	"github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"
	"league-havester/harvester/connector"
	"league-havester/harvester/finder"
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
	cancelFunc      context.CancelFunc
	harvesterOption Option
	finderOption    finder.Option
	pool            *ants.Pool
}

func (h *harvester) Start() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	h.infinity()
	sig := <-signals
	log.Info(sig)
	h.cancelFunc()
}

func (h *harvester) infinity() {
	go func() {
		ticker := time.NewTicker(h.harvesterOption.HarvesterDuration)
		for {
			select {
			case <-ticker.C:
				if connector.IsNotRunning() {
					if result, err := finder.New(h.ctx, h.finderOption).Execute(); err != nil {
						log.WithError(err).Debug("error")
					} else {
						log.WithField("result", result).Debug("find fail")
						c := connector.New(h.ctx, h.pool, connector.Config(result))
						c.Connect()
					}
				} else {
					log.Debug("connector is running")
				}
			case <-h.ctx.Done():
				log.Info("terminated")
				h.cancelFunc()
				break
			}
		}
	}()
}

func New(harvesterOption Option, finderOption finder.Option) Harvester {
	ctx, cancelFunc := context.WithCancel(context.Background())
	pool, err := ants.NewPool(-1)
	if err != nil {
		panic(err)
	}
	return &harvester{
		ctx:             ctx,
		cancelFunc:      cancelFunc,
		harvesterOption: harvesterOption,
		finderOption:    finderOption,
		pool:            pool,
	}
}
