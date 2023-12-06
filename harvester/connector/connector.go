package connector

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"
	"league-havester/api"
	"league-havester/harvester/model"
	"league-havester/harvester/model/eogStatsBlock"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	isRunning bool
	runMutex  sync.Mutex
)

func IsNotRunning() bool {
	runMutex.Lock()
	defer runMutex.Unlock()
	return !isRunning
}

func setRunning(running bool) {
	runMutex.Lock()
	defer runMutex.Unlock()
	isRunning = running
}

func init() {
	isRunning = false
	runMutex = sync.Mutex{}
}

type Option struct {
	ConnectTimeout time.Duration
	WriteTimeout   time.Duration
	PongWait       time.Duration
	PingPeriod     time.Duration
	Host           string
	Port           int
	Pid            int
	Token          string
}

type Connector interface {
	Connect()
}

type connector struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	pool       *ants.Pool
	Option
}

func (c *connector) Connect() {
	setRunning(true)
	defer setRunning(false)

	u := url.URL{Scheme: "wss", Host: "127.0.0.1:" + strconv.Itoa(c.Port)}
	timeout, _ := context.WithTimeout(c.ctx, c.ConnectTimeout)
	header := http.Header{}
	header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(("riot:"+c.Token))))
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	websocket.DefaultDialer.TLSClientConfig = tlsConfig
	conn, response, err := websocket.DefaultDialer.DialContext(timeout, u.String(), header)
	if err != nil {
		log.Error(err)
		return
	}
	log.WithField("response", response).Debug("websocket dial failed")

	defer conn.Close()

	conn.SetWriteDeadline(time.Now().Add(c.WriteTimeout))
	err = conn.WriteMessage(websocket.TextMessage, []byte("[5, \"OnJsonApiEvent_lol-end-of-game_v1_eog-stats-block\"]"))
	if err != nil {
		log.Error(err)
		return
	} else {
		log.Info("league client connected")
	}

	go func() {
		ticker := time.NewTicker(c.PingPeriod)
		for {
			select {
			case <-ticker.C:
				conn.SetWriteDeadline(time.Now().Add(c.WriteTimeout))
				if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					log.WithError(err).Error("ping failed")
					c.cancelFunc()
					if err := conn.Close(); err != nil {
						log.Error(err)
					}
					return
				} else {
					log.Debug("ping")
				}
			case <-c.ctx.Done():
				if err := conn.Close(); err != nil {
					log.Error(err)
				}
				log.Info("connector terminated")
				c.cancelFunc()
				return
			}
		}
	}()

	readPump := make(chan *eogStatsBlock.EogStatsBlock, 1)
	go func() {
		conn.SetReadDeadline(time.Now().Add(c.PongWait))
		conn.SetPongHandler(func(string) error {
			log.Debug("pong")
			conn.SetReadDeadline(time.Now().Add(c.PongWait))
			return nil
		})
		for {
			if _, message, err := conn.ReadMessage(); err != nil {
				log.WithError(err).Error("read failed")
				c.cancelFunc()
				break
			} else {
				rawMessage := string(message)
				if len(rawMessage) > 9 {
					secondCommaIndex := strings.Index(rawMessage, "\",")
					rawBody := rawMessage[secondCommaIndex+2 : len(rawMessage)-1]
					log.Debug("=========================================================================================================")
					log.WithField("rawBody", rawBody).Trace("read block")
					event := model.Event{}
					err := json.Unmarshal([]byte(rawBody), &event)
					if err != nil {
						log.WithError(err).Error("err unmarshal event")
						return
					}
					if event.Uri == "/lol-end-of-game/v1/eog-stats-block" {
						block := &eogStatsBlock.EogStatsBlock{}
						err := json.Unmarshal([]byte(rawBody), block)
						if err != nil {
							log.WithError(err).Error("err unmarshal block")
							return
						}
						readPump <- block
					}
				} else {
					log.WithFields(log.Fields{
						"raw": rawMessage,
						"len": len(rawMessage),
					}).Debug("short raw message")
				}
			}
		}
	}()

	for {
		select {
		case block := <-readPump:
			c.pool.Submit(func() {
				err := api.Request(block)
				if err != nil {
					log.WithError(err).Error("api request failed")
				}
			})
		case <-c.ctx.Done():
			if err := conn.Close(); err != nil {
				log.Error(err)
			}
			log.Info("connector terminated")
			c.cancelFunc()
			return
		}
	}
}

func New(ctx context.Context, pool *ants.Pool, option Option) Connector {
	innerCtx, cancelFunc := context.WithCancel(ctx)
	return &connector{
		ctx:        innerCtx,
		cancelFunc: cancelFunc,
		pool:       pool,
		Option:     option,
	}
}
