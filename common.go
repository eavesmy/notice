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
	Chan     chan string
	rateChan chan bool
	ErrChan  chan error
	AckChan  chan int
}

func (c *Client) Init() *Client {

	c.Opt.Default()

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

	if c.Opt.MaxBytesLimit != 0 && len(buffer.Bytes()) > c.Opt.MaxBytesLimit {
		return errors.New("msg oversize: %d/%d")
	}

	if len(c.Chan) > c.Opt.MaxMsgLimit {
		return errors.New(fmt.Sprintf("msg channel overflow: %d/%d", len(c.Chan), c.Opt.MaxMsgLimit))
	}

	c.Chan <- string(buffer.Bytes())

	return nil
}

func (c *Client) start() {
	for {
		msg := <-c.Chan

		statuc, err := c.Cli.Send(msg)
		fmt.Println(statuc, err)

		time.Sleep(c.Opt.Rate)
	}
}

// close all chan
func (c *Client) close() {
	<-c.Ctx.Done()

	c.Opt.Log.Println("Close notice service.")

	close(c.Chan)
}
