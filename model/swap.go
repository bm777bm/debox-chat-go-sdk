package model

type TokenData struct {
	ChainId         int64  `json:"chain_id"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	ContractAddress string `json:"contract_address"`
	Volm            string
	Unit            string
	PriceDecrease   float64
}
type DexTokenInfo struct {
	Code int32       `json:"code"`
	Msg  string      `json:"message"`
	Data []TokenData `json:"data"`
}

// {"instType":"SWAP","instId":"BTC-USDC-SWAP","last":"67447.9","lastSz":"439","askPx":"67436.8","askSz":"439","bidPx":"67436.2","bidSz":"1564","open24h":"66817.7","high24h":"68121.3","low24h":"66556.6","volCcy24h":"1769.657","vol24h":"17696570","ts":"1709887584704","sodUtc0":"66936.4","sodUtc8":"67338.1"}
type TokenSwapInfoFromOk struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Open24h string `json:"open24h"`
		High24h string `json:"high24h"`
		Low24h  string `json:"low24h"`
		Vol24h  string `json:"vol24h"`
		Last    string `json:"last"`
		InstId  string `json:"instId"`
	} `json:"data"`
}

type UITagA struct {
	Uitag string `json:"uitag"`
	Text  string `json:"text"`
	Href  string `json:"href"`
}

type UITagFont struct {
	Uitag  string `json:"uitag,omitempty"`
	Text   string `json:"text,omitempty"`
	Bold   string   `json:"bold,omitempty"`
	Italic string `json:"italic,omitempty"`
	Color  string `json:"color,omitempty"`
}

type UITagImg struct {
	Uitag    string `json:"uitag,omitempty"`    //img
	Src      string `json:"src,omitempty"`      //img src
	Position string `json:"position,omitempty"` // head foot
	Width    string `json:"width,omitempty"`    //img height，大于0生效，否则表示没设置将用默认值
	Height   string `json:"height,omitempty"`   //img height，大于0生效，否则表示没设置将用默认值
	Href     string `json:"href,omitempty"`     // img href
}
type UITagImgInt struct {
	Uitag    string `json:"uitag"`            //img
	Src      string `json:"src"`              //img src
	Position string `json:"position"`         // head foot
	Width    int    `json:"width,omitempty"`  //img height，大于0生效，否则表示没设置将用默认值
	Height   int    `json:"height,omitempty"` //img height，大于0生效，否则表示没设置将用默认值
	Href     string `json:"href"`             // img href
}
