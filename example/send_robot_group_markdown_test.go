package main

import (
	"fmt"
	"testing"

	// boxbotapi "github.com/debox-pro/debox-chat-go-sdk"
	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
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
var (
	// Menu texts
	firstMenu  = "<b>Menu 1</b>\n\nA box button message."
	secondMenu = "<b>Menu 2</b>\n\nA box button message."

	// Button texts
	nextButton     = "Next"
	nextButton1    = "Next1"
	nextButton6    = "N"
	backButton     = "Back"
	tutorialButton = "Tutorial"
	tokenUrl       = "https://deswap.pro/?from_chain_id=-200&from_address=11111111111111111111111111111111&to_chain_id=-200&to_address=BpykKPT9DoPy2WoZspkd7MvUb9QAPtX86ojmrg48pump"
	// Store bot screaming status
	screaming = false
	bot       *boxbotapi.BotAPI

	// Keyboard layout for the first menu. One button, one row

	firstMenuMarkup = boxbotapi.NewInlineKeyboardMarkup(
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonURL("url1", tokenUrl),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonURL("url", tokenUrl),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonURL("url", tokenUrl),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonURL("url", tokenUrl),
		),
	)

	// Keyboard layout for the second menu. Two buttons, one per row
	secondMenuMarkup = boxbotapi.NewInlineKeyboardMarkup(
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonData(backButton, backButton),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonURL(tutorialButton, "https://core.telegram.org/bots/api"),
		),
	)

	thirdMenuMarkup = boxbotapi.NewInlineKeyboardMarkup(
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonURL(tutorialButton, tokenUrl),
			boxbotapi.NewInlineKeyboardButtonDataWithColor("👍🏻", "reaction", "61", "#00ff00"),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonDataWithColor("👍🏻", "reaction", "61", "#00ff00"),
			boxbotapi.NewInlineKeyboardButtonData("👎🏻", "reaction"),
			boxbotapi.NewInlineKeyboardButtonData("❤️", "reaction"),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
			boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
		),

		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BTC", "reaction1", "61", "#ff0000"),
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BNB", "reaction1", "27.5%", "#00ff00"),
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BNB", "reaction", "27.5%", "#0000ff"),
		),
		boxbotapi.NewInlineKeyboardRow(
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BTC", "reaction1", " ", "#ff0000"),
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BNB", "reaction1", " ", "#00ff00"),
			boxbotapi.NewInlineKeyboardButtonDataWithColor("BNB", "reaction", " ", "#0000ff"),
		),
	)
)

func TestSendRobotGroupMarkdownMsg1(t *testing.T) {

	xApiKey := "P55X0r5xfDpm5Yc5"
	xApiKey = "ggowK0QRl1UPkPA9" //测试chatbot ,370
	// xApiKey = "ggowK0QRl1UPkPA9" //测试 ,用户的

	client := boxbotapi.CreateNormalInterface("https://open.debox.pro", xApiKey)
	// client = boxbotapi.CreateNormalInterface("http://t.open.debox.pro", xApiKey)
	client = boxbotapi.CreateNormalInterface("http://127.0.0.1:8041", xApiKey)

	//https://s.debox.pro/group?id=ii0jiad9
	//https://s.debox.pro/group?id=ii0jiad9
	// toUserId := "uvg2p6ho" //接收者id
	//https://s.debox.pro/group?id=ayoe8lz6
	//https://s.debox.pro/group?id=nhu775tk
	//https://s.debox.pro/group?id=mao2vuey
	//https://s.debox.pro/group?id=w8cgtfof
	// https://s.debox.pro/group?id=ymor0jin
	var toUserId = "x1dei8zv1"
	groupId := "fxi3hqo5" //群组id
	groupId = "128907"    //test1 正式
	groupId = "ii0jiad9"  //112club//Test777
	groupId = "3lifa7j6"  //test New7
	// groupId = "ayoe8lz6"       //test New7
	// groupId = "nhu775tk"       //test New7
	groupId = "ymor0jin" //test New7

	contentHTML := `
	<span style="color:red">span123</span>
	<b>bold</b>, <strong>bold</strong>
	<i>italic</i>, <em>italic</em>
	<u>underline</u>, <ins>underline</ins>
	<s>strikethrough</s>, <strike>strikethrough</strike>, <del>strikethrough</del>
	<span class="tg-spoiler">spoiler</span>, <tg-spoiler>spoiler</tg-spoiler>
	<b>bold <i>italic bold <s>italic bold strikethrough <span class="tg-spoiler">italic bold strikethrough spoiler</span></s> <u>underline italic bold</u></i> bold</b>
	<a href="http://www.example.com/">inline URL</a>
	<a href="tg://user?id=123456789">inline mention of a user</a>
	<a href="https://debox.pro">debox</a>
	<tg-emoji emoji-id="5368324170671202286"></tg-emoji>
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
	// objectName = "HTML" //消息类型（ "RCD:Graphic"  文本消息； "RCD:Graphic" 图文消息）
	content := contentMD
	content = contentNormal
	content = contentHTML
	content = contentMD

	thirdMenuMarkup.FontSize = "s"
	thirdMenuMarkup.FontColor = "#0000ff"

	firstMenuMarkup.FontSize = "s"
	firstMenuMarkup.FontColor = "#ff0000"
	var message = boxbotapi.MarkdownV2Config{
		ToUserId:         toUserId,
		GroupId:          groupId,
		Content:          content,
		ObjectName:       objectName,
		UserActionMarkup: &firstMenuMarkup,
		ReplyMarkup:      &thirdMenuMarkup,
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

	client := boxbotapi.CreateNormalInterface("https://open.debox.pro", xApiKey)
	// client = boxbotapi.CreateNormalInterface("http://t.open.debox.pro", xApiKey)
	client = boxbotapi.CreateNormalInterface("http://127.0.0.1:8041", xApiKey)

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
