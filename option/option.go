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

type Option struct {
	Webhook   string
	Uuid      string
	Keyword   string
	Signature string

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
		o.Retry = 3
	}
	if o.Timeout == 0 {
		o.Timeout = time.Second * time.Duration(3)
	}
	if o.Rate == 0 {
		o.Rate = time.Millisecond * time.Duration(500)
	}

	if o.Log == nil {
		o.Log = log.New(os.Stdout, "", 0)
	}
}
