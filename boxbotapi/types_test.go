package boxbotapi

import (
	"testing"
	"time"
)

func TestUserStringWith(t *testing.T) {
	user := User{
		UserId:       "dfdfd",
		FirstName:    "Test",
		LastName:     "Test",
		Name:         "",
		LanguageCode: "en",
		IsBot:        false,
	}

	if user.String() != "Test Test" {
		t.Fail()
	}
}

func TestUserStringWithUserName(t *testing.T) {
	user := User{
		UserId:       "fdfd",
		FirstName:    "Test",
		LastName:     "Test",
		Name:         "@test",
		LanguageCode: "en",
	}

	if user.String() != "@test" {
		t.Fail()
	}
}

func TestMessageTime(t *testing.T) {
	message := Message{Date: 0}

	date := time.Unix(0, 0)
	if message.Time() != date {
		t.Fail()
	}
}

func TestMessageIsCommandWithCommand(t *testing.T) {
	message := Message{Text: "/command"}
	message.Entities = []MessageEntity{{Type: "bot_command", Offset: 0, Length: 8}}

	if !message.IsCommand() {
		t.Fail()
	}
}

func TestIsCommandWithText(t *testing.T) {
	message := Message{Text: "some text"}

	if message.IsCommand() {
		t.Fail()
	}
}

func TestIsCommandWithEmptyText(t *testing.T) {
	message := Message{Text: ""}

	if message.IsCommand() {
		t.Fail()
	}
}

func TestCommandWithCommand(t *testing.T) {
	message := Message{Text: "/command"}
	message.Entities = []MessageEntity{{Type: "bot_command", Offset: 0, Length: 8}}

	if message.Command() != "command" {
		t.Fail()
	}
}

func TestCommandWithEmptyText(t *testing.T) {
	message := Message{Text: ""}

	if message.Command() != "" {
		t.Fail()
	}
}

func TestCommandWithNonCommand(t *testing.T) {
	message := Message{Text: "test text"}

	if message.Command() != "" {
		t.Fail()
	}
}

func TestCommandWithBotName(t *testing.T) {
	message := Message{Text: "/command@testbot"}
	message.Entities = []MessageEntity{{Type: "bot_command", Offset: 0, Length: 16}}

	if message.Command() != "command" {
		t.Fail()
	}
}

func TestCommandWithAtWithBotName(t *testing.T) {
	message := Message{Text: "/command@testbot"}
	message.Entities = []MessageEntity{{Type: "bot_command", Offset: 0, Length: 16}}

	if message.CommandWithAt() != "command@testbot" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsWithArguments(t *testing.T) {
	message := Message{Text: "/command with arguments"}
	message.Entities = []MessageEntity{{Type: "bot_command", Offset: 0, Length: 8}}
	if message.CommandArguments() != "with arguments" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsWithMalformedArguments(t *testing.T) {
	message := Message{Text: "/command-without argument space"}
	message.Entities = []MessageEntity{{Type: "bot_command", Offset: 0, Length: 8}}
	if message.CommandArguments() != "without argument space" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsWithoutArguments(t *testing.T) {
	message := Message{Text: "/command"}
	if message.CommandArguments() != "" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsForNonCommand(t *testing.T) {
	message := Message{Text: "test text"}
	if message.CommandArguments() != "" {
		t.Fail()
	}
}

func TestMessageEntityParseURLGood(t *testing.T) {
	entity := MessageEntity{URL: "https://www.google.com"}

	if _, err := entity.ParseURL(); err != nil {
		t.Fail()
	}
}

func TestMessageEntityParseURLBad(t *testing.T) {
	entity := MessageEntity{URL: ""}

	if _, err := entity.ParseURL(); err == nil {
		t.Fail()
	}
}

func TestChatIsPrivate(t *testing.T) {
	chat := Chat{ID: "fdfd", Type: "private"}

	if !chat.IsPrivate() {
		t.Fail()
	}
}

func TestChatIsGroup(t *testing.T) {
	chat := Chat{ID: "fdfd", Type: "group"}

	if !chat.IsGroup() {
		t.Fail()
	}
}

func TestChatIsChannel(t *testing.T) {
	chat := Chat{ID: "fdfd", Type: "channel"}

	if !chat.IsChannel() {
		t.Fail()
	}
}

func TestChatIsSuperGroup(t *testing.T) {
	chat := Chat{ID: "fdfd", Type: "supergroup"}

	if !chat.IsSuperGroup() {
		t.Fail()
	}
}

func TestMessageEntityIsMention(t *testing.T) {
	entity := MessageEntity{Type: "mention"}

	if !entity.IsMention() {
		t.Fail()
	}
}

func TestMessageEntityIsHashtag(t *testing.T) {
	entity := MessageEntity{Type: "hashtag"}

	if !entity.IsHashtag() {
		t.Fail()
	}
}

func TestMessageEntityIsBotCommand(t *testing.T) {
	entity := MessageEntity{Type: "bot_command"}

	if !entity.IsCommand() {
		t.Fail()
	}
}

func TestMessageEntityIsUrl(t *testing.T) {
	entity := MessageEntity{Type: "url"}

	if !entity.IsURL() {
		t.Fail()
	}
}

func TestMessageEntityIsEmail(t *testing.T) {
	entity := MessageEntity{Type: "email"}

	if !entity.IsEmail() {
		t.Fail()
	}
}

func TestMessageEntityIsBold(t *testing.T) {
	entity := MessageEntity{Type: "bold"}

	if !entity.IsBold() {
		t.Fail()
	}
}

func TestMessageEntityIsItalic(t *testing.T) {
	entity := MessageEntity{Type: "italic"}

	if !entity.IsItalic() {
		t.Fail()
	}
}

func TestMessageEntityIsCode(t *testing.T) {
	entity := MessageEntity{Type: "code"}

	if !entity.IsCode() {
		t.Fail()
	}
}

func TestMessageEntityIsPre(t *testing.T) {
	entity := MessageEntity{Type: "pre"}

	if !entity.IsPre() {
		t.Fail()
	}
}

func TestMessageEntityIsTextLink(t *testing.T) {
	entity := MessageEntity{Type: "text_link"}

	if !entity.IsTextLink() {
		t.Fail()
	}
}

func TestFileLink(t *testing.T) {
	file := File{FilePath: "test/test.txt"}

	if file.Link("token") != "https://api.telegram.org/file/bottoken/test/test.txt" {
		t.Fail()
	}
}

// Ensure all configs are sendable
var (
	_ Chattable = CloseConfig{}
	_ Chattable = CopyMessageConfig{}
	_ Chattable = DeleteWebhookConfig{}
	_ Chattable = ForwardConfig{}
	_ Chattable = LogOutConfig{}
	_ Chattable = MessageConfig{}
	_ Chattable = PhotoConfig{}
	_ Chattable = WebhookConfig{}
)

// Ensure all Fileable types are correct.
var (
	_ Fileable = (*PhotoConfig)(nil)
	_ Fileable = (*WebhookConfig)(nil)
)

// Ensure all RequestFileData types are correct.
var (
	_ RequestFileData = (*FilePath)(nil)
	_ RequestFileData = (*FileBytes)(nil)
	_ RequestFileData = (*FileReader)(nil)
	_ RequestFileData = (*FileURL)(nil)
	_ RequestFileData = (*FileID)(nil)
	_ RequestFileData = (*fileAttach)(nil)
)
