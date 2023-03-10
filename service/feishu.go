/*
# File Name: feishu.go
# Author : eavesmy
# Email:eavesmy@gmail.com
# Created Time: 2022年02月25日 星期五 15时07分22秒
*/

package service

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/eavesmy/notice/option"
)

const BASE_URL = "https://open.feishu.cn/open-apis/bot/v2/hook/"

type Feishu struct {
	client *http.Client
	Ctx    context.Context
	Opt    option.Option
}

type FeishuPayload struct {
	MsgType string `json:"msg_type"`
	Content struct {
		Text string `json:"text"`
	} `json:"content"`
}

func (s *Feishu) Send(msg *bytes.Buffer) (statusCode int, err error) {

	if s.client == nil {
		s.client = &http.Client{}
	}

	body := &FeishuPayload{
		MsgType: "text",
	}

	body.Content.Text = msg.String()

	if s.Opt.Webhook == "" {
		s.Opt.Webhook = BASE_URL + s.Opt.Uuid
	}

	if s.Opt.Keyword != "" {
		body.Content.Text = s.Opt.Keyword + "\n" + msg.String()
	}

	b, err := json.Marshal(body)
	if err != nil {
		return
	}

	buffer := bytes.NewBuffer(b)

	req, err := http.NewRequestWithContext(s.Ctx, http.MethodPost, s.Opt.Webhook, buffer)

	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	res, err := s.client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()

	// 一次性读完
	b, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	statusCode = res.StatusCode

	return
}
