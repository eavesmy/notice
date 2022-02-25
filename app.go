// Usage:
// notice.Feishu(options...).Send(msg)
// or
// notice.Group{notice.Feishu(options...),notice.Dingding(options...)}.Send(msg)

package notice

import (
	"context"

	"github.com/eavesmy/notice/option"
)

func Feishu(opt option.Option, ctx context.Context) *Common {
	return (&Common{
		Channel: CHANNEL_FEISHU,
		Opt:     opt,
		Ctx:     ctx,
		Chan:    make(chan string, 10),
	}).Init()
}

type Group []*Common

func (g *Group) Send(msg interface{}) (chan int, chan error) {

	ack := make(chan int, len(*g))
	errs := make(chan error, len(*g))

	for _, comm := range *g {
		code, err := comm.Send(msg)

		ack <- code
		errs <- err
	}

	return ack, errs
}
