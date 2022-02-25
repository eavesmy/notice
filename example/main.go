package main

import (
	"context"
	"fmt"
	"time"

	"github.com/eavesmy/notice"
	"github.com/eavesmy/notice/option"
)

func main() {

	ctx := context.Background()
	feishu := notice.Feishu(option.Option{
		Keyword: "PCOIN",
		Webhook: "",
		Rate:    time.Second * 2,
	}, ctx)

	code, err := feishu.Send("test")
	fmt.Println(code, err)

	code, err = feishu.Send("test")
	fmt.Println(code, err)

	code, err = feishu.Send("test")
	fmt.Println(code, err)

	code, err = feishu.Send("test")
	fmt.Println(code, err)
}
