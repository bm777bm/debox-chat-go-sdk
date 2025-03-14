package main

import (
	"fmt"
	"testing"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
	dbx_chat "github.com/debox-pro/debox-chat-go-sdk/boxbotapi/deboxapi"
)

//该函数用来发图片消息，发群消息。
//如果apikey绑定了debox用户则以debox用户的名义发送消息，具体某个用户
//如果没绑定debox账户，则报错，发送失败

// toUserId := "uvg2p6ho"                                                                                //接收者id
// groupId := "fxi3hqo5"                                                                                 //群组id
// title := "im title"                                                                                   //消息标题
// content := "im content"                                                                               //消息内容
// objectName := "RCD:Graphic"
// href :="https://debox.pro/"   图文消息，传入跳转链接

func TestSendRobotGroupMarkdownMsg1(t *testing.T) {

	xApiKey := "P55X0r5xfDpm5Yc5"
	xApiKey = "ggowK0QRl1UPkPA9" //测试chatbot ,370
	// xApiKey = "ggowK0QRl1UPkPA9" //测试 ,用户的

	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)
	// client = boxbotapi.CreateNormalInterface("http://t.open.debox.pro", xApiKey)
	client = dbx_chat.CreateNormalInterface("http://127.0.0.1:8042", xApiKey)

	//https://s.debox.pro/group?id=ii0jiad9
	var toUserId = "x1dei8zv1"
	groupId := "fxi3hqo5" //群组id
	groupId = "128907"    //test1 正式
	groupId = "ii0jiad9"  //112club//Test777
	// groupId = "3lifa7j6"  //test New7
	// groupId = "ayoe8lz6"       //test New7
	// groupId = "nhu775tk"       //test New7
	// groupId = "ymor0jin" //import

	contentHTML := `
	<span style="color:red">span123</span>
	<b>bold</b>,nobold <strong>bold</strong>
	<i>italic</i>, <em>italic</em>
	<u>underline</u>, <ins>underline</ins>
	<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
	<span class="box-spoiler">spoiler</span>, <box-spoiler>spoiler</box-spoiler>
	<b>bold <i>italic bold <s>italic bold strikethrough <span class="box-spoiler">italic bold strikethrough spoiler</span></s> <u>underline italic bold</u></i> bold</b>
	<a href="http://www.example.com/">inline URL</a>
	<a href="box://user?id=123456789">inline mention of a user</a>
	<a href="https://debox.pro">debox</a>
	<box-emoji emoji-id="5368324170671202286"></box-emoji>
	<code>inline fixed-width code</code>
	<pre>pre-formatted fixed-width code block</pre>
	<pre><code class="language-python">pre-formatted fixed-width code block written in the Python programming language</code></pre>
	<blockquote>Block quotation started\nBlock quotation continued\nThe last line of the block quotation</blockquote>
	<blockquote expandable>Expandable block quotation started\nExpandable block quotation continued\nExpandable block quotation continued\nHidden by default part of the block quotation started\nExpandable block quotation continued\nThe last line of the block quotation</blockquote>
	`
	contentMD := "*粗斜体*,\n" +
		"**粗斜体**,\n" +
		"~~strikethrough~~\n" +
		"# 一级标题。\n" +
		"[debox](https://debox.pro/)\n" +
		"## 22222222BTC\n" +
		"### 3333333BTC\n" +
		"#### 44444BTC\n" +
		"##### 55555555BTC\n" +
		"###### 6666666BTC\n" +
		"####### 7777777$BOX"
	contentNormal := "$box"
	objectName := "MarkdownV2" //消息类型（ "RCD:Graphic"  文本消息； "RCD:Graphic" 图文消息）
	// objectName = "richtext"    //消息类型（ "RCD:Graphic"  文本消息； "RCD:Graphic" 图文消息）
	objectName = "HTML" //消息类型（ "RCD:Graphic"  文本消息； "RCD:Graphic" 图文消息）
	content := contentMD
	content = contentNormal
	content = contentHTML
	// content = contentMD

	secondMenuMarkup.FontSize = "s"
	secondMenuMarkup.FontColor = "#0000ff"

	firstMenuMarkup.FontSize = "s"
	firstMenuMarkup.FontColor = "#ff0000"
	var message = boxbotapi.MarkdownV2Config{
		ToUserId:         toUserId,
		GroupId:          groupId,
		Content:          content,
		ObjectName:       objectName,
		UserActionMarkup: &firstMenuMarkup,
		ReplyMarkup:      &secondMenuMarkup,
	}
	_, err := client.Send(message)
	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")
}

func TestSendRobotGroupMarkdownMsg(t *testing.T) {

	xApiKey := "P55X0r5xfDpm5Yc5"
	xApiKey = "ggowK0QRl1UPkPA9" //测试chatbot ,370
	// xApiKey = "ggowK0QRl1UPkPA9" //测试 ,用户的

	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)
	// client = boxbotapi.CreateNormalInterface("http://t.open.debox.pro", xApiKey)
	client = dbx_chat.CreateNormalInterface("http://127.0.0.1:8042", xApiKey)

	// toUserId := "uvg2p6ho" //接收者id
	var toUserId = "x1dei8zv1"
	groupId := "fxi3hqo5" //群组id
	groupId = "128907"    //test1 正式
	groupId = "ii0jiad9"  //112club//Test777
	groupId = "ymor0jin"  //test New7
	title := "im title"   //消息标题
	title = ""
	objectName := "MarkdownV2" //消息类型（ "RCD:Graphic"  文本消息； "RCD:Graphic" 图文消息）
	objectName = "richtext"    //消息类型（ "RCD:Graphic"  文本消息； "RCD:Graphic" 图文消息）
	message := ""              //图文消息链接
	content := "![RUNOOB 图标](https://static.jyshare.com/images/runoob-logo.png)\n" +
		"*粗斜体*,\n" +
		"**粗斜体**,\n" +
		"~~strikethrough~~\n" +
		"# 这是一段红色的文字。\n" +
		"[debox](https://debox.pro/)\n" +
		"## 22222222\n" +
		"### 3333333\n" +
		"#### 444444\n" +
		"##### 55555\n" +
		"###### 666\n" +
		"Quote\n" +
		">Quote11111111\n" +
		">>Quote2222222\n" +
		">>Quote3333333\n" +
		">>Quote4444444\n" +
		"* List1\n" +
		"* List2\n" +
		"- List3\n" +
		"- List4\n" +
		"+ List5\n" +
		"+ List6\n" +
		"```function(){alert('1111');}```"
	_, err := client.SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, "send_robot_group_msg", "")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
