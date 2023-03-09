/*
# File Name: ./option/option.go
# Author : eavesmy
# Email:eavesmy@gmail.com
# Created Time: 2022年02月25日 星期五 16时16分28秒
*/

package option

import (
	"log"
	"os"
	"time"
)

const DEFAULT_MAX_MSG_LIMIT = 200
const DEFAULT_RATE = time.Millisecond * 500
const DEFAULT_TIMEOUT = time.Second * 3
const DEFAULT_RETRY = 3

type Option struct {
	Webhook       string
	Uuid          string
	Keyword       string
	Signature     string
	MaxMsgLimit   int
	MaxBytesLimit int

	// Retry call times default: 3
	Retry int
	// Call timeout. default: 3s
	Timeout time.Duration
	// Send msg interval. default: 500ms.
	Rate time.Duration
	Log  *log.Logger
}

func (o *Option) Default() {

	if o.Retry == 0 {
		o.Retry = DEFAULT_RETRY
	}
	if o.Timeout == 0 {
		o.Timeout = DEFAULT_TIMEOUT
	}
	if o.Rate == 0 {
		o.Rate = DEFAULT_RATE
	}

	if o.Log == nil {
		o.Log = log.New(os.Stdout, "", 0)
	}

	if o.MaxMsgLimit == 0 {
		o.MaxMsgLimit = DEFAULT_MAX_MSG_LIMIT
	}
}
