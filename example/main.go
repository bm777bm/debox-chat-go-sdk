package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"strings"

	boxbotapi "github.com/debox-pro/debox-chat-go-sdk/boxbotapi"
)

// var (
// // Menu texts
// firstMenu  = "<b>Menu 1</b>\n\nA box."
// secondMenu = "<b>Menu 2</b>\n\nA box button message."

// // Button texts
// nextButton     = "Next"
// nextButton1    = "Next1"
// nextButton6    = "N777777777777777777777777777777"
// backButton     = "Back"
// tutorialButton = "Tutorial"

// Store bot screaming status
// screaming = false
// bot *boxbotapi.BotAPI

// Keyboard layout for the first menu. One button, one row

// firstMenuMarkup = boxbotapi.NewInlineKeyboardMarkup(
// 	boxbotapi.NewInlineKeyboardRow(
// 		boxbotapi.NewInlineKeyboardButtonData(nextButton, nextButton),
// 		boxbotapi.NewInlineKeyboardButtonSwitch("打开会话并切换", ""),
// 		boxbotapi.NewInlineKeyboardButtonSwitchCurrentChat("打开当前", ""),
// 	),
// 	boxbotapi.NewInlineKeyboardRow(
// 		boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
// 		boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
// 		boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
// 		boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
// 		boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
// 		boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
// 		boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
// 		boxbotapi.NewInlineKeyboardButtonData(nextButton1, nextButton1),
// 	),
// )

// Keyboard layout for the second menu. Two buttons, one per row
// secondMenuMarkup = boxbotapi.NewInlineKeyboardMarkup(
//
//	boxbotapi.NewInlineKeyboardRow(
//		boxbotapi.NewInlineKeyboardButtonData(backButton, backButton),
//	),
//	boxbotapi.NewInlineKeyboardRow(
//		boxbotapi.NewInlineKeyboardButtonURL(tutorialButton, "https://core.telegram.org/bots/api"),
//	),
//
// )
// // location button
// thirdMenuMarkup = boxbotapi.NewReplyKeyboard(
//
//	boxbotapi.NewKeyboardButtonRow(
//		boxbotapi.NewKeyboardButtonContact("分享联系人"),
//		boxbotapi.NewKeyboardButtonLocation("分享位置"),
//	),
//	boxbotapi.NewKeyboardButtonRow(
//		boxbotapi.NewKeyboardButtonContact("分享联系人2"),
//		boxbotapi.NewKeyboardButtonLocation("分享位置2"),
//	),
//
// )
// )

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
	screaming = true
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

func main() {
	var err error
	//bm set proxy begin
	// proxyURL, err := url.Parse("http://127.0.0.1:7890")
	// if err != nil {
	// 	log.Fatalf("Failed to parse proxy URL: %v", err)
	// }

	// client := &http.Client{
	// 	Transport: &http.Transport{
	// 		Proxy: http.ProxyURL(proxyURL),
	// 	},
	// 	Timeout: 10 * time.Second,
	// }
	// bot, err = tgbotapi.NewBotAPI("7748647347:AAGr9bPH1PcjtZ5h70FNlKEqh427Ww5SwFM")
	bot, err = boxbotapi.NewBotAPIWithClient("pPpHtOTtXsE6i5u6", boxbotapi.APIEndpoint, nil)
	//bm end
	// bot, err = tgbotapi.NewBotAPI("<YOUR_BOT_TOKEN_HERE>")
	// bot, err = tgbotapi.NewBotAPI("7748647347:AAGr9bPH1PcjtZ5h70FNlKEqh427Ww5SwFM")
	if err != nil {
		// Abort if something is wrong
		log.Panic(err)
	}

	// Set this to true to log all interactions with telegram servers
	bot.Debug = true

	u := boxbotapi.NewUpdate(0)
	u.Timeout = 60

	// Create a new cancellable background context. Calling `cancel()` leads to the cancellation of the context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// `updates` is a golang channel which receives telegram updates
	updates := bot.GetUpdatesChan(u)

	// Pass cancellable context to goroutine
	go receiveUpdates(ctx, updates)

	// Tell the user the bot is online
	log.Println("Start listening for updates. Press enter to stop")

	// Wait for a newline symbol, then cancel handling updates
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	cancel()

}

func receiveUpdates(ctx context.Context, updates boxbotapi.UpdatesChannel) {
	// `for {` means the loop is infinite until we manually stop it
	for {
		select {
		// stop looping if ctx is cancelled
		case <-ctx.Done():
			return
		// receive update from channel and then handle it
		case update := <-updates:
			handleUpdate(update)
		}
	}
}

func handleUpdate(update boxbotapi.Update) {
	switch {
	// Handle messages
	case update.Message != nil:
		handleMessage(update.Message)
		break

	// Handle button clicks
	case update.CallbackQuery != nil:
		handleButton(update.CallbackQuery)
		break
	}
}

func handleMessage(message *boxbotapi.Message) {
	user := message.From
	text := message.Text

	if user == nil {
		return
	}

	// Print to console
	log.Printf("%s wrote %s", user.Name, text)

	var err error
	if strings.HasPrefix(text, "/") {
		err = handleCommand(message.Chat.ID, text)
	} else if screaming && len(text) > 0 {
		// msg := boxbotapi.NewMessage(message.Chat.ID, strings.ToUpper(text))
		msg := boxbotapi.NewMessageResponse(message)
		// To preserve markdown, we attach entities (bold, italic..)
		// msg.Entities = message.Entities
		_, err = bot.Send(msg)
	} else {
		// This is equivalent to forwarding, without the sender's name
		copyMsg := boxbotapi.NewCopyMessage(message.Chat.ID, message.Chat.ID, message.MessageID)
		_, err = bot.CopyMessage(copyMsg)
	}

	if err != nil {
		log.Printf("An error occured: %s", err.Error())
	}
}

// When we get a command, we react accordingly
func handleCommand(chatId string, command string) error {
	var err error

	switch command {
	case "/scream":
		screaming = true
		break

	case "/whisper":
		screaming = false
		break

	case "/menu":
		err = sendMenu(chatId)
		break

	case "/menu2":
		err = sendMenu2(chatId)
		break
	}

	return err
}

func handleButton(query *boxbotapi.CallbackQuery) {
	var text string

	markup := boxbotapi.NewInlineKeyboardMarkup()
	message := query.Message

	if query.Data == nextButton {
		text = secondMenu
		markup = secondMenuMarkup
	} else if query.Data == backButton {
		text = firstMenu
		markup = firstMenuMarkup
	}

	callbackCfg := boxbotapi.NewCallback(query.ID, "")
	bot.Send(callbackCfg)

	// Replace menu text and keyboard
	msg := boxbotapi.NewEditMessageTextAndMarkup(message.Chat.ID, message.MessageID, text, markup)
	msg.ParseMode = boxbotapi.ModeHTML
	bot.Send(msg)
}

func sendMenu(chatId string) error {
	msg := boxbotapi.NewMessage(chatId, firstMenu)
	msg.ParseMode = boxbotapi.ModeHTML
	msg.ReplyMarkup = firstMenuMarkup
	_, err := bot.Send(msg)
	return err
}

func sendMenu2(chatId string) error {
	msg := boxbotapi.NewMessage(chatId, firstMenu)
	msg.ParseMode = boxbotapi.ModeHTML
	// msg.ReplyMarkup = firstMenuMarkup
	msg.ReplyMarkup = thirdMenuMarkup
	_, err := bot.Send(msg)
	return err
}
