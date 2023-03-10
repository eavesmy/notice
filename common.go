package notice

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/eavesmy/notice/option"
	"github.com/eavesmy/notice/service"
	"time"
)

const (
	channel_feishu = iota
)

type Client struct {
	Channel  int
	Opt      option.Option
	Ctx      context.Context
	Cli      service.Service
	Chan     chan *bytes.Buffer
	rateChan chan bool
	ErrChan  chan error
	AckChan  chan int
}

func (c *Client) Init() *Client {

	switch c.Channel {
	case channel_feishu:
		c.Cli = &service.Feishu{Ctx: c.Ctx, Opt: c.Opt}
	default:
	}

	go c.start()

	return c
}

func (c *Client) Send(msg string) (err error) {

	buffer := bytes.NewBufferString(msg)

	if c.Opt.MaxBytesLimit != 0 && buffer.Len() > c.Opt.MaxBytesLimit {
		return errors.New("msg oversize: %d/%d")
	}

	if len(c.Chan) > c.Opt.MaxMsgLimit {
		return errors.New(fmt.Sprintf("msg channel overflow: %d/%d", len(c.Chan), c.Opt.MaxMsgLimit))
	}

	go func() {
		c.Chan <- buffer
	}()

	return nil
}

func (c *Client) start() {

	timer := time.NewTimer(c.Opt.Rate)

	for {
		<-timer.C

		fmt.Printf("发送 Feishu. 当前剩余 %d/%d 条\n ", len(c.Chan), c.Opt.MaxMsgLimit)
		c.Cli.Send(<-c.Chan)

		timer.Reset(c.Opt.Rate)
	}
}

// close all chan
func (c *Client) close() {
	<-c.Ctx.Done()

	c.Opt.Log.Println("Close notice service.")

	close(c.Chan)
}
