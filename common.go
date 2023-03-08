package notice

import (
	"bytes"
	"context"
	"sync"
	"time"

	"github.com/eavesmy/notice/option"
	"github.com/eavesmy/notice/service"
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

	lock sync.WaitGroup
}

func (c *Client) Init() *Client {

	c.Opt.Default()

	c.rateChan = make(chan bool, 1)

	c.ErrChan = make(chan error, 10)
	c.AckChan = make(chan int, 10)
	c.lock = sync.WaitGroup{}

	switch c.Channel {
	case channel_feishu:
		c.Cli = &service.Feishu{Ctx: c.Ctx, Opt: c.Opt}
	default:
	}

	go c.startLoop()
	go c.start()

	return c
}

func (c *Client) Send(msg string) (statusCode int, err error) {

	buffer := bytes.NewBufferString(msg)

	c.Chan <- string(buffer.Bytes())

	return <-c.AckChan, <-c.ErrChan
}

func (c *Client) SendMsgRun() {
	if len(c.Chan) > 0 {
		msg := <-c.Chan

		statusCode, err := c.Cli.Send(msg)

		c.AckChan <- statusCode
		c.ErrChan <- err
	}
}

func (c *Client) start() {
	for {
		<-c.rateChan

		go c.SendMsgRun()
	}
}

func (c *Client) startLoop() {
	time.AfterFunc(c.Opt.Rate, c.startLoop)

	c.rateChan <- true
}

// close all chan
func (c *Client) close() {
	<-c.Ctx.Done()

	c.Opt.Log.Println("Close notice service.")

	close(c.rateChan)
	close(c.Chan)
}

func (c *Client) Error() chan error {
	return c.ErrChan
}
