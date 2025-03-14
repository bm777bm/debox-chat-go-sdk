package boxbotapi

import (
	"log"
	"net/http"
	"testing"
	"time"
)

const (
	TestToken               = "153667468:AAHlSHlMqSt1f_uFmVRJbm5gntu2HI4WW8I"
	ChatID                  = "76918703"
	Channel                 = "@tgbotapitest"
	SupergroupChatID        = "-1001120141283"
	ReplyToMessageID        = 35
	ExistingPhotoFileID     = "AgACAgIAAxkDAAEBFUZhIALQ9pZN4BUe8ZSzUU_2foSo1AACnrMxG0BucEhezsBWOgcikQEAAwIAA20AAyAE"
	ExistingDocumentFileID  = "BQADAgADOQADjMcoCcioX1GrDvp3Ag"
	ExistingAudioFileID     = "BQADAgADRgADjMcoCdXg3lSIN49lAg"
	ExistingVoiceFileID     = "AwADAgADWQADjMcoCeul6r_q52IyAg"
	ExistingVideoFileID     = "BAADAgADZgADjMcoCav432kYe0FRAg"
	ExistingVideoNoteFileID = "DQADAgADdQAD70cQSUK41dLsRMqfAg"
	ExistingStickerFileID   = "BQADAgADcwADjMcoCbdl-6eB--YPAg"
)

type testLogger struct {
	t *testing.T
}

func (t testLogger) Println(v ...interface{}) {
	t.t.Log(v...)
}

func (t testLogger) Printf(format string, v ...interface{}) {
	t.t.Logf(format, v...)
}

func getBot(t *testing.T) (*BotAPI, error) {
	bot, err := NewBotAPI(TestToken)
	bot.Debug = true

	logger := testLogger{t}
	SetLogger(logger)

	if err != nil {
		t.Error(err)
	}

	return bot, err
}

func TestNewBotAPI_notoken(t *testing.T) {
	_, err := NewBotAPI("")

	if err == nil {
		t.Error(err)
	}
}

func TestGetUpdates(t *testing.T) {
	bot, _ := getBot(t)

	u := NewUpdate(0)

	_, err := bot.GetUpdates(u)

	if err != nil {
		t.Error(err)
	}
}

func TestSendWithMessage(t *testing.T) {
	bot, _ := getBot(t)

	msg := NewMessage(ChatID, "group", "A test message from the test library in telegram-bot-api")
	msg.ParseMode = ModeMarkdown
	_, err := bot.Send(msg)

	if err != nil {
		t.Error(err)
	}
}

func ExampleNewBotAPI() {
	bot, err := NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.Name)

	u := NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// Optional: wait for updates and clear them if you don't want to handle
	// a large backlog of old messages
	time.Sleep(time.Millisecond * 500)
	updates.Clear()

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.Name, update.Message.Text)

		msg := NewMessage(update.Message.Chat.ID, update.Message.Chat.Type, update.Message.Text)

		bot.Send(msg)
	}
}

func ExampleNewWebhook() {
	bot, err := NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.Name)

	wh, err := NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, FilePath("cert.pem"))

	if err != nil {
		panic(err)
	}

	_, err = bot.Request(wh)

	if err != nil {
		panic(err)
	}

	info, err := bot.GetWebhookInfo()

	if err != nil {
		panic(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("failed to set webhook: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}

func ExampleWebhookHandler() {
	bot, err := NewBotAPI("MyAwesomeBotToken")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.Name)

	wh, err := NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, FilePath("cert.pem"))

	if err != nil {
		panic(err)
	}

	_, err = bot.Request(wh)
	if err != nil {
		panic(err)
	}
	info, err := bot.GetWebhookInfo()
	if err != nil {
		panic(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("[DeBox callback failed]%s", info.LastErrorMessage)
	}

	http.HandleFunc("/"+bot.Token, func(w http.ResponseWriter, r *http.Request) {
		update, err := bot.HandleUpdate(r)
		if err != nil {
			log.Printf("%+v\n", err.Error())
		} else {
			log.Printf("%+v\n", *update)
		}
	})

	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
}
