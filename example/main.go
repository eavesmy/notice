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
		Keyword: "TEST",
		// Uuid:    "33cd14f7-d482-4464-8d51-20237962f62f",
		//   or
		Webhook: "https://open.feishu.cn/open-apis/bot/v2/hook/5bcf3df5-f6b6-4aac-a67e-8ad8c998af38",
		Rate:    time.Second * 2,
	}, ctx)

	for {
		fmt.Println(time.Now().Format(time.RFC3339), feishu.Send("test"))
	}
}
