package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"

	dbx_chat "github.com/debox-pro/debox-chat-go-sdk/boxbotapi/deboxapi"

	"github.com/debox-pro/debox-chat-go-sdk/common"
	"github.com/debox-pro/debox-chat-go-sdk/model"
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

func TestSendRobotSwapMsgHot(t *testing.T) {
	// sendSwapHot()
}

func sendSwapHot1() error {

	xApiKey := "t2XJiou2Mu6AlEF6" //配置齐全,正式
	// xApiKeyTest := "ggowK0QRl1UPkPA9" //测试
	client := dbx_chat.CreateNormalInterface("https://open.debox.pro", xApiKey)
	//7d6089qb
	groupId := "fxi3hqo5" //test
	groupId = "l3ixp32y"  //test1
	// groupId = "xfinpgfj"  //白金瀚
	// groupId = "9rib1g30" //对抗群
	// groupId = "uu08d676"  //正式kobe群
	// groupId = "xfinpfrf"  //测试喵喵群
	// groupId = "ztuvxylb" //星辰群
	title := "1h热门交易币种"         //消息标题
	content := "im content"     //无用
	objectName := "RCD:Graphic" //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）

	url := "http://127.0.0.1:8042/openapi/swap/hot?chain_id=1"
	// url := "https://open.debox.pro/openapi/swap/hot?chain_id=1"
	var from_address = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"

	var header = map[string]string{
		"X-API-KEY": xApiKey,
	}
	tokenInfo := model.DexTokenInfo{}
	err := HttpGet2Obj(url, header, &tokenInfo)
	if err != nil || len(tokenInfo.Data) == 0 {
		return err
	}
	tokenInfo.Data = common.Filter(tokenInfo.Data, func(td model.TokenData) bool {
		return td.ContractAddress != from_address
	})
	tokenInfo.Data = common.RemoveDuplication(tokenInfo.Data, func(td model.TokenData) string {
		return td.Symbol
	})

	var str = "项目      成交额(U)    涨跌(%)"
	// for i, token := range tokenInfo.Data {
	for i, token := range tokenInfo.Data {
		if i > 9 {
			break
		}
		var swapUrlForamt = "https://debox.pro/deswap/?from_chain_id=%d&from_address=%s&to_chain_id=%d&to_address=%s"
		var href = fmt.Sprintf(swapUrlForamt, token.ChainId, from_address, token.ChainId, token.ContractAddress)
		var uiA = model.UITagA{
			Uitag: "a",
			Text:  token.Symbol,
			Href:  href,
		}
		jsonUIA, err := json.Marshal(uiA)
		if err != nil {
			return err
		}
		priceDecrease := strconv.FormatFloat(token.PriceDecrease*100, 'f', 2, 64) //+ "%"
		//成交额从第8个开始、涨跌幅从19个开始
		var strCoin = string(jsonUIA)

		var lenCoinSpace = 7 - len(token.Symbol)
		var strCoinSpace = strings.Repeat(" ", lenCoinSpace*2)

		var strVolum = token.Volm + token.Unit

		var lenVolumSpace = 17 - (8 + len(strVolum))
		var strVolumSpace = strings.Repeat(" ", lenVolumSpace*2)

		str += "\n" + strCoin + strCoinSpace + strVolum + strVolumSpace + priceDecrease
	}
	content = str
	_, err = client.SendRobotGroupMsg("", groupId, title, content, "1", objectName, "send_robot_group_msg", "")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return err
	}
	fmt.Println("send chat message success.")
	return nil

}
