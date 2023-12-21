package main

import (
	"fmt"
	"testing"

	dbx_chat "github.com/debox-pro/debox-chat-go-sdk"
)

// toUserId := "uvg2p6ho"            //接收者id
// groupId := "fxi3hqo5"             //群组id
// title := "im title"               //消息标题
// content := "im content"           //消息内容
// objectName := "RC:TxtMsg"         //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）
// message := "im SendRobotGroupMsg" //图文消息时传入图片链接，文字消息时传入文字

func TestSendRobotGroupMsg(t *testing.T) {

	xApiKey := "t2XJiou2Mu6AlEF6"

	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

	toUserId := "uvg2p6ho"            //接收者id
	groupId := "fxi3hqo5"             //群组id
	title := "im title"               //消息标题
	content := "im content"           //消息内容
	objectName := "RC:TxtMsg"         //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）
	message := "im SendRobotGroupMsg" //图文消息链接

	_, err := client.SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, "send_robot_group_msg")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}

// toUserId := "uvg2p6ho"                                                                                //接收者id
// groupId := "fxi3hqo5"                                                                                 //群组id
// title := "im title"                                                                                   //消息标题
// content := "im content"                                                                               //消息内容
// objectName := "RCD:Graphic"
func TestSendRobotGroupImg(t *testing.T) {

	xApiKey := "t2XJiou2Mu6AlEF6"

	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)

	toUserId := "uvg2p6ho"                                                                                //接收者id
	groupId := "fxi3hqo5"                                                                                 //群组id
	title := "im title"                                                                                   //消息标题
	content := "im content"                                                                               //消息内容
	objectName := "RCD:Graphic"                                                                           //消息类型（ "RCD:Graphic"  文本消息； "RCD:Graphic" 图文消息）
	message := "https://data.debox.space/static/2023/11/22/ii0k2v5n/85a86525054b20457949986850767e93.jpg" //图文消息链接

	_, err := client.SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, "send_robot_group_msg")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
