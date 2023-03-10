// Usage:
// notice.Feishu(options...).Send(msg)
// or
// notice.Group{notice.Feishu(options...),notice.Dingding(options...)}.Send(msg)

package notice

import (
	"bytes"
	"context"
	"sync"

	"github.com/eavesmy/notice/option"
)

func Feishu(opt option.Option, ctx context.Context) *Client {

	opt.Default()

	return (&Client{
		Channel: channel_feishu,
		Opt:     opt,
		Ctx:     ctx,
		Chan:    make(chan *bytes.Buffer, opt.MaxMsgLimit),
	}).Init()
}

type Group []*Client

func (g *Group) Send(msg string) (errors []error) {

	wg := sync.WaitGroup{}
	wg.Add(len(*g))

	for _, comm := range *g {
		go func(c *Client) {
			errors = append(errors, c.Send(msg))
		}(comm)
	}

	return errors
}
