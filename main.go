package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	r "github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/sashabaranov/go-openai"
)

// 微信相关密钥配置
var (
	wxAppId         = kingpin.Flag("wx_app_id", "please access https://mp.weixin.qq.com/").Required().String()
	wxOriId         = kingpin.Flag("wx_ori_id", "please access https://mp.weixin.qq.com/").Required().String()
	wxToken         = kingpin.Flag("wx_token", "please access https://mp.weixin.qq.com/").Required().String()
	wxEncodedAESKey = kingpin.Flag("wx_aes_key", "please access https://mp.weixin.qq.com/").Required().String()
)

// openapi相关配置
var (
	baseURL = kingpin.Flag("base_url", "access open api prefix url").
		Default("https://agent-openai.ccrui.dev/v1").String()
	openAPIToken = kingpin.Flag("token", "please access https://platform.openai.com/account/api-keys").Required().String()
)

// http相关
var (
	httpPort = kingpin.Flag("port", "http port").Default("80").Int()
	httpPath = kingpin.Flag("path", "http path").Default("/").String()
)

// 下面两个变量不一定非要作为全局变量, 根据自己的场景来选择.
var (
	msgHandler core.Handler
	msgServer  *core.Server
)

func textMsgHandler(ctx *core.Context) {
	log.Printf("收到文本消息:\n%s\n", ctx.MsgPlaintext)
	msg := request.GetText(ctx.MixedMsg)
	sendMsg(ctx, msg)
}

// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func wxCallbackHandler(w http.ResponseWriter, r *http.Request) {
	msgServer.ServeHTTP(w, r, nil)
}

func main() {
	kingpin.Parse()

	mux := core.NewServeMux()
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)

	msgHandler = mux
	msgServer = core.NewServer(*wxOriId, *wxAppId, *wxToken, *wxEncodedAESKey, msgHandler, nil)
	http.HandleFunc(*httpPath, wxCallbackHandler)
	log.Printf("Starting server on :%v", *httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func sendMsg(ctx *core.Context, msg *request.Text) {
	config := openai.DefaultConfig(*openAPIToken)
	config.BaseURL = *baseURL
	client := openai.NewClientWithConfig(config)
	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: msg.Content,
				},
			},
			MaxTokens: 2048,
		},
	)
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return
	}
	resp := r.NewText(msg.FromUserName, msg.ToUserName, time.Now().Unix(),
		response.Choices[0].Message.Content)
	ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}
