package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"testing"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk"

	common "github.com/debox-pro/debox-chat-go-sdk/common"
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

func TestSendRobotSwapMsg(t *testing.T) {

	xApiKey := "t2XJiou2Mu6AlEF6"     //配置齐全,正式
	xApiKeyTest := "ggowK0QRl1UPkPA9" //测试
	client := boxbotapi.CreateNormalInterface("https://open.debox.pro", xApiKey)
	//7d6089qb
	groupId := "fxi3hqo5" //test
	groupId = "l3ixp32y"  //test1
	// groupId = "9rib1g30" //对抗群
	// groupId = "uu08d676"  //正式kobe群
	// groupId = "xfinpfrf"  //测试喵喵群
	title := "im title"         //消息标题
	content := "im content"     //无用
	objectName := "RCD:Graphic" //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）
	message := "6"
	href := ""

	url := "http://127.0.0.1:8041/openapi/swap/statistics"
	var header = map[string]string{
		"X-API-KEY": xApiKeyTest,
	}

	tokenInfo := model.DexTokenInfo{}
	err := HttpGet2Obj(url, header, &tokenInfo)
	if err != nil {
		return
	}

	tokenSwapInfo := model.TokenSwapInfoFromOk{}
	err = HttpGet2Obj("https://www.okx.com/api/v5/market/tickers?instType=SWAP", nil, &tokenSwapInfo)
	if err != nil {
		return
	}
	//补偿成交量、涨跌幅数据
	for i, token := range tokenInfo.Data {
		InstId := fmt.Sprintf("%s-USDT-SWAP", token.Symbol)
		for _, tokenSwap := range tokenSwapInfo.Data {
			if tokenSwap.InstId == InstId {
				volmStr, _ := strconv.ParseFloat(tokenSwap.Vol24h, 64)
				tokenInfo.Data[i].Volm, tokenInfo.Data[i].Unit = common.CalsVolumUnit(float64(volmStr))
				open, _ := strconv.ParseFloat(tokenSwap.Open24h, 64)
				last, _ := strconv.ParseFloat(tokenSwap.Last, 64)
				tokenInfo.Data[i].PriceDecrease = last/open - 1
				break
			}
		}
	}
	tokenInfo.Data = common.Filter(tokenInfo.Data, func(td model.TokenData) bool {
		return td.Volm != ""
	})
	//过滤ok没有的币种
	//按涨跌幅排序
	sort.Slice(tokenInfo.Data, func(i, j int) bool {
		return tokenInfo.Data[i].PriceDecrease > tokenInfo.Data[j].PriceDecrease
	})

	title = "热门交易币种"
	var str = "项目      成交额(U)    涨跌幅"
	for i, token := range tokenInfo.Data {
		if i > 10 {
			break
		}
		var from_address = "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE"
		var swapUrlForamt = "https://debox.pro/deswap/?from_chain_id=%d&from_address=%s&to_chain_id=%d&to_address=%s"
		var href = fmt.Sprintf(swapUrlForamt, token.ChainId, from_address, token.ChainId, token.ContractAddress)
		var uiA = model.UITagA{
			Uitag: "a",
			Text:  token.Symbol,
			Href:  href,
		}
		jsonUIA, err := json.Marshal(uiA)
		if err != nil {
			return
		}
		priceDecrease := strconv.FormatFloat(token.PriceDecrease*100, 'f', 2, 64) + "%"
		str += "\n" + string(jsonUIA) + "       " + token.Volm + token.Unit + "     " + priceDecrease
	}
	content = str
	_, err = client.SendRobotGroupMsg("", groupId, title, content, message, objectName, "send_robot_group_msg", href)

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}
