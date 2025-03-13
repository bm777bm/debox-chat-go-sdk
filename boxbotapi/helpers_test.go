package boxbotapi

import (
	"testing"
)

func TestNewWebhook(t *testing.T) {
	result, err := NewWebhook("https://example.com/token")

	if err != nil ||
		result.URL.String() != "https://example.com/token" ||
		result.Certificate != interface{}(nil) ||
		result.MaxConnections != 0 ||
		len(result.AllowedUpdates) != 0 {
		t.Fail()
	}
}

func TestNewInlineKeyboardButtonLoginURL(t *testing.T) {
	result := NewInlineKeyboardButtonLoginURL("text", LoginURL{
		URL:                "url",
		ForwardText:        "ForwardText",
		BotUsername:        "username",
		RequestWriteAccess: false,
	})

	if result.Text != "text" ||
		result.LoginURL.URL != "url" ||
		result.LoginURL.ForwardText != "ForwardText" ||
		result.LoginURL.BotUsername != "username" ||
		result.LoginURL.RequestWriteAccess != false {
		t.Fail()
	}
}
