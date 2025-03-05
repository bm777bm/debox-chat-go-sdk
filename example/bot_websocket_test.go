package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

func getSignature(secret string) (nonce, timestamp, signature string) {
	nonceInt := rand.Int()
	nonce = strconv.Itoa(nonceInt)
	timeInt64 := time.Now().Unix()
	timestamp = strconv.FormatInt(timeInt64, 10)
	h := sha1.New()
	_, _ = io.WriteString(h, secret+nonce+timestamp)
	signature = fmt.Sprintf("%x", h.Sum(nil))
	return
}
func TestBotClientTest(t *testing.T) {
	var appid = "LlLFPjkZTSBoG771"
	var apiKey = "svRfurVpnBcxPh11"
	var secret = "gzzwrxxMvUHlfkXu9ZrhpRoiJlfzEKa3"
	nonce, timestamp, signature := getSignature(secret)

	// WebSocket 服务的 URL
	//http://t.open.debox.pro/openapi/robot_msg/notify
	url := "ws://localhost:8041/openapi/robot_msg/listen?nonce=%s&timestamp=%s&signature=%s&app_id=%s&api_key=%s"
	// url = "ws://t.open.debox.pro/openapi/robot_msg/listen?nonce=%s&timestamp=%s&signature=%s&app_id=%s&api_key=%s"
	url = fmt.Sprintf(url, nonce, timestamp, signature, appid, apiKey)

	// 创建一个新的 WebSocket 连接
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatalf("Dial error: %v", err)
	}
	defer conn.Close()

	// 读取 WebSocket 消息
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			// 读取消息
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Printf("Read error: %v", err)
				return
			}
			// 检查消息类型是否为文本消息
			if messageType == websocket.TextMessage {
				fmt.Printf("Received message: %s\n", p)
			} else {
				log.Printf("Received non-text message: %d bytes", len(p))
			}
		}
	}()

	// 运行一段时间以接收消息，然后关闭连接
	// time.Sleep(30 * time.Second) // 例如，运行30秒
	// conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(time.Second))
	<-done

	log.Println("Client exiting")
}
