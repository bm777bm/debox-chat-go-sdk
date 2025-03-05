package main

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

//该函数用来发文字消息，发群消息。
//如果apikey绑定了debox用户则以debox用户的名义发送消息，具体某个用户
//如果没绑定debox账户，则报错，发送失败

// toUserId := "uvg2p6ho"            //接收者id
// groupId := "fxi3hqo5"             //群组id
// title := "im title"               //消息标题
// content := "im content"           //消息内容
// objectName := "RC:TxtMsg"         //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）
// message := "im SendRobotGroupMsg" //图文消息时传入图片链接，文字消息时传入文字
// href :="" 文字消息，此参数传空即可

func TestSendRobotGroupMsg1(t *testing.T) {
	mainwss()
}

func mainwss() {
	// WebSocket服务器地址
	u := url.URL{Scheme: "wss", Host: "io.dexscreener.com", Path: "/dex/screener/pairs/h24/1?rankBy[key]=trendingScoreH1&rankBy[order]=desc&filters[chainIds][0]=ethereum"}
	// 打印连接信息
	log.Printf("connecting to %s", u.String())

	// 建立WebSocket连接
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	// 启动协程用于接收来自WebSocket服务器的消息
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	// 发送GET请求
	err = c.WriteMessage(websocket.TextMessage, []byte("GET"))
	if err != nil {
		log.Println("write:", err)
		return
	}

	// 等待Ctrl+C退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-sig:
		log.Println("interrupt")

		// 发送关闭消息到WebSocket服务器
		err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			return
		}
		// 等待服务器关闭连接
		select {
		case <-done:
		case <-time.After(time.Second):
		}
		return
	}
}
