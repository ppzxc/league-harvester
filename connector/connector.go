package connector

import (
	"context"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
	"time"
)

type Option struct {
	ConnectTimeout time.Duration
	Host           string
	Port           int
	Pid            int
	Token          string
}

type Connector interface {
	Connect()
}

type client struct {
	ctx context.Context
	Option
}

func (c *client) Connect() {
	u := url.URL{Scheme: "wss", Host: "127.0.0.1:" + strconv.Itoa(c.Port), User: url.UserPassword("riot", c.Token)}
	log.Println(u.String())
	timeout, _ := context.WithTimeout(c.ctx, c.ConnectTimeout)
	conn, _, err := websocket.DefaultDialer.DialContext(timeout, u.String(), nil)
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	done := make(chan struct{})

	go func() {
		defer close(done)

		for {
			messageType, p, err2 := conn.ReadMessage()
			if err2 != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %d: %s", messageType, p)
		}
	}()
	<-done
}

func New(ctx context.Context, option Option) Connector {
	return &client{
		ctx:    ctx,
		Option: option,
	}
}
