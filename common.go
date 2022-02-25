package notice

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/eavesmy/notice/option"
	"github.com/eavesmy/notice/service"
)

const (
	CHANNEL_FEISHU = iota
)

type Common struct {
	Channel  int
	Opt      option.Option
	Ctx      context.Context
	Cli      service.Service
	Chan     chan string
	rateChan chan bool

	ErrChan chan error
	AckChan chan int

	lock sync.WaitGroup
}

func (c *Common) Init() *Common {

	c.Opt.Default()

	c.rateChan = make(chan bool, 1)

	c.ErrChan = make(chan error, 10)
	c.AckChan = make(chan int, 10)
	c.lock = sync.WaitGroup{}

	switch c.Channel {
	case CHANNEL_FEISHU:
		c.Cli = &service.Feishu{Ctx: c.Ctx, Opt: c.Opt}
	default:
	}

	go c.startLoop()
	go c.start()

	return c
}

func (c *Common) Send(msg interface{}) (statusCode int, err error) {

	b, err := json.Marshal(msg)
	if err != nil {
		c.ErrChan <- err
		return
	}

	buffer := bytes.NewBuffer(b)

	fmt.Println(len(c.Chan), len(c.AckChan), len(c.ErrChan))

	c.Chan <- string(buffer.Bytes())

	fmt.Println(len(c.Chan), len(c.AckChan), len(c.ErrChan))

	return <-c.AckChan, <-c.ErrChan
}

func (c *Common) SendMsgRun() {
	if len(c.Chan) > 0 {
		msg := <-c.Chan

		statusCode, err := c.Cli.Send(msg)

		c.AckChan <- statusCode
		c.ErrChan <- err
	}
}

func (c *Common) start() {
	for {
		<-c.rateChan

		fmt.Println("rate done")

		go c.SendMsgRun()
	}
}

func (c *Common) startLoop() {
	time.AfterFunc(c.Opt.Rate, c.startLoop)

	c.rateChan <- true
}

// close all chan
func (c *Common) close() {
	<-c.Ctx.Done()

	c.Opt.Log.Println("Close notice service.")

	close(c.rateChan)
	close(c.Chan)
}

func (c *Common) Error() chan error {
	return c.ErrChan
}
