package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"testing"
	"time"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/deboxapiold"

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
type DexTokenInfo struct {
	Id              int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT;comment:主键" json:"id"`
	ChainId         int64  `gorm:"column:chain_id;type:int(11);default:0;comment:链id" json:"chain_id"`
	Name            string `gorm:"column:name;type:varchar(255);comment:名字" json:"name"`
	Symbol          string `gorm:"column:symbol;type:varchar(255);comment:符号" json:"symbol"`
	Decimals        int    `gorm:"column:decimals;type:int(11);default:0;comment:精度" json:"decimals"`
	ContractAddress string `gorm:"column:contract_address;type:varchar(255);comment:合约地址" json:"contract_address"`
	LogoUrl         string `gorm:"column:logo_url;type:varchar(255);comment:token图标" json:"logo_url"`
	AllowanceCount  string `gorm:"column:allowance_count;type:int(11);comment:token图标" json:"allowance_count"`
	CreateTime      int    `gorm:"column:create_time;type:int(11);comment:创建时间;NOT NULL" json:"create_time"`
	ModifyTime      int    `gorm:"column:modify_time;type:int(11);comment:更新时间" json:"modify_time"`
	IsDelete        int    `gorm:"column:is_delete;type:int(11);comment:是否删除" json:"is_delete"`
	// 排序权重，值越大权重越高，JSON中不显示
	Weight int `gorm:"column:weight;type:int(11);default:0;comment:权重;NOT NULL" json:"-"`
}

// {"instType":"SWAP","instId":"BTC-USDC-SWAP","last":"67447.9","lastSz":"439","askPx":"67436.8","askSz":"439","bidPx":"67436.2","bidSz":"1564","open24h":"66817.7","high24h":"68121.3","low24h":"66556.6","volCcy24h":"1769.657","vol24h":"17696570","ts":"1709887584704","sodUtc0":"66936.4","sodUtc8":"67338.1"}
type TokenSwapInfo struct {
	Open24h string `json:"open24h"`
	High24h string `json:"high24h"`
	Low24h  int    `json:"low24h"`
	Vol24h  string `json:"vol24h"`
	Last    string `json:"last"`
	InstId  string `json:"instId"`
}

func TestSendRobotSwapMsg1(t *testing.T) {

	xApiKey := "t2XJiou2Mu6AlEF6" //配置齐全,正式
	// xApiKeyTest := "ggowK0QRl1UPkPA9" //测试
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

	// url := "http://127.0.0.1:8041/openapi/swap/statistics"
	url := "https://open.debox.pro/openapi/swap/hot?chain_id=1"
	var header = map[string]string{
		"X-API-KEY": xApiKey,
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
		var uiA = UITagA{
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
	var uiImg1 = UITagImgString{
		Uitag:    "img",
		Src:      "https://data.debox.space/dao/newpic/one.png",
		Position: "head",
		Href:     "https://www.sina.com.cn",
		Height:   "200",
	}
	jsonUIImg1, err := json.Marshal(uiImg1)
	var jsonUIImg11 = string(jsonUIImg1)
	if err != nil {
		println(jsonUIImg1)
		println(jsonUIImg11)
		return
	}
	var uiImg2 = UITagImgInt{
		Uitag:    "img",
		Src:      "https://data.debox.space/dao/newpic/two.png",
		Position: "foot",
		// Href:     "https://www.163.com",
		Height: 300,
	}
	jsonUIImg2, err := json.Marshal(uiImg2)
	if err != nil {
		print(jsonUIImg2)
		return
	}

	var href3 = "https://app.debox.pro/"
	var uiA = UITagA{
		Uitag: "a",
		Text:  "DeBox",
		Href:  href3,
	}
	jsonUIA3, err := json.Marshal(uiA)
	if err != nil {
		print(jsonUIA3)
		return
	}
	// content = str + "你好：" + string(jsonUIA3) + "，欢迎加入" + string(jsonUIImg2)

	// content = jsonUIImg11
	// content = jsonUIImg11 + string(jsonUIImg2)
	content = str + jsonUIImg11 + string(jsonUIImg2)

	// content = jsonUIImg11 + string(jsonUIImg2)
	// title = ""

	// _, err := client.SendRobotGroupMsg(toUserId, groupId, title, content, message, objectName, "send_robot_group_msg", href)
	// title = ""
	_, err = client.SendRobotGroupMsg("", groupId, title, content, message, objectName, "send_robot_group_msg", href)

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}

type UITagA struct {
	Uitag string `json:"uitag"`
	Text  string `json:"text omitempty"`
	Href  string `json:"href omitempty"`
}
type UITagImgInt struct {
	Uitag    string `json:"uitag1"`             //img
	Src      string `json:"src omitempty"`      //img src
	Position string `json:"position omitempty"` // head foot
	Height   int32  `json:"height omitempty"`   //img height，大于0生效，否则表示没设置将用默认值
	Href     string `json:"href"`               // img href
}
type UITagImgString struct {
	Uitag    string `json:"uitag"`              //img
	Src      string `json:"src omitempty"`      //img src
	Position string `json:"position omitempty"` // head foot
	Height   string `json:"height omitempty"`   //img height，大于0生效，否则表示没设置将用默认值
	Href     string `json:"href omitempty"`     // img href
}

// 给私人发送消息
// 只能发送text消息
// objectName 必须为 "RCD:Command"
func TestSendRobotSwapPrivateMsg(t *testing.T) {

	xApiKey := "t2XJiou2Mu6AlEF6"

	client := boxbotapi.CreateNormalInterface("https://open.debox.pro", xApiKey)
	// https://m.debox.pro/card?id=ii0k2v5n
	toUserId := "ii0k2v5n"           //接收者id xul
	toUserId = "oo0epqhd"            //小战
	message := "im  message content" //消息内容
	objectName := "RCD:Command"      //消息类型（ "RC:TxtMsg"  文本消息； "RCD:Graphic" 图文消息）

	var url = "http://127.0.0.1:8041/openapi/swap/statistics"
	var header = map[string]string{
		"X-API-KEY": "ggowK0QRl1UPkPA9",
	}
	resp, err1 := HttpGet(url, header)
	if err1 != nil {
		return
	}
	println(resp["data"])
	var str = "10分钟热门交易币种\n"
	str += "项目名称       成交量\n"
	for _, obj := range resp["data"].([]interface{}) {
		var obj1 = obj.(map[string]interface{})
		var coinName = obj1["symbol"].(string)
		var volm = obj1["vol"].(float64)
		var unit = ""
		if volm/1000 > 1 {
			unit = "k"
			volm = volm / 1000
		}
		if volm/1000 > 1 {
			unit = "m"
			volm = volm / 1000
		}
		if volm/1000 > 1 {
			unit = "b"
			volm = volm / 1000
		}
		var vol = strconv.FormatFloat(volm, 'f', 2, 64)
		str += coinName + "           " + vol + unit + "\n"
	}
	str += "戳我交易 https://debox.pro/deswap/"
	message = str

	_, err := client.SendRobotMsg(toUserId, message, objectName, "send_robot_msg")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

	if err != nil {
		fmt.Println("send chat message fail:", err)
		return
	}

	fmt.Println("send chat message success.")

}

//发送群组、文字消息
//test https://m.debox.pro/card?id=ii0k2v5n
// curl -X POST -H "Content-Type: application/json" -H "X-API-KEY: t2XJiou2Mu6AlEF6" -d '{"to_user_id":"ii0k2v5n1","group_id":"fxi3hqo5","object_name":"RC:TxtMsg","message":"I am a message sent by bot. 三四五点钟分别发一条"}'  "https://open.debox.pro/openapi/send_robot_group_message"

//对抗群
// curl -X POST -H "Content-Type: application/json" -H "X-API-KEY: t2XJiou2Mu6AlEF6" -d '{"to_user_id":"uvg2p6ho","group_id":"fxi3hqo5","object_name":"RC:TxtMsg","message":"i am a message sent by bot"}'  "https://open.debox.pro/openapi/send_robot_group_message"

// //发送群组、图片文字  消息
// // curl -X POST -H "Content-Type: application/json" -H "X-API-KEY: t2XJiou2Mu6AlEF6" -d '{"to_user_id":"uvg2p6ho","group_id":"fxi3hqo5","object_name":"RCD:Graphic","title":"i am title","content":"i am content","message":"https://data.debox.space/dao/newpic/one.png"}'  "https://open.debox.pro/openapi/send_robot_group_message"

// //发送个人、文字消息
// // curl -X POST -H "Content-Type: application/json" -H "X-API-KEY: t2XJiou2Mu6AlEF6" -d '{"to_user_id":"uvg2p6ho","object_name":"RCD:Command","message":"i am a message to user from bot"}'  "https://open.debox.pro/openapi/send_robot_message"

func HttpGet1(url string, header map[string]string) (map[string]interface{}, error) {
	var ret map[string]interface{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get new failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "en_US")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Timeout: 600 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get send failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// LogOut("error", "http get resp failed, url:"+url+", Body :"+lib1.Json_Package(resp)+"result:"+lib1.Json_Package(ret))
		return ret, errors.New("wrong http code" + strconv.Itoa(resp.StatusCode))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get unmarshal failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}

	//LogOut("info", fmt.Sprintf("http get success, url: %s, resp: %v", url, header, ret))
	return ret, nil
}

func HttpPost1(url string, header map[string]string) (map[string]interface{}, error) {
	var ret map[string]interface{}

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get new failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Content-Type", "en_US")
	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{
		Timeout: 600 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get send failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		// LogOut("error", "http get resp failed, url:"+url+", Body :"+lib1.Json_Package(resp)+"result:"+lib1.Json_Package(ret))
		return ret, errors.New("wrong http code" + strconv.Itoa(resp.StatusCode))
	}

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ret)
	if err != nil {
		// LogOut("error", fmt.Sprintf("http get unmarshal failed, url: %s, err: %s", url, err.Error()))
		return ret, err
	}

	//LogOut("info", fmt.Sprintf("http get success, url: %s, resp: %v", url, header, ret))
	return ret, nil
}
