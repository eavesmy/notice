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
		Uuid:    "33cd14f7-d482-4464-8d51-20237962f62f",
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
