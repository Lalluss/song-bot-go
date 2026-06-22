package main

import (
"fmt"
"log"
"net/http"
"os"
"strings"

tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)

func main() {
const CHANNEL_ID int64 = -1004306196694
go func() {
port := os.Getenv("PORT")
if port == "" {
port = "8000"
}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bot is running"))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}()

token := os.Getenv("BOT_TOKEN")

bot, err := tgbotapi.NewBotAPI(token)
if err != nil {
	log.Fatal(err)
}

log.Printf("Bot started: %s", bot.Self.UserName)

u := tgbotapi.NewUpdate(0)
u.Timeout = 60

updates := bot.GetUpdatesChan(u)

for update := range updates {

	if update.Message == nil {
		continue
	}
        if update.Message.Chat.ID == CHANNEL_ID &&
update.Message.Document != nil {

title := update.Message.Caption

if title == "" {
	continue
}

log.Printf("Movie uploaded: %s", title)

bot.Send(
	tgbotapi.NewMessage(
		update.Message.Chat.ID,
		"✅ Movie detected: "+title,
	),
)

continue

}
	// Welcome new members
	if len(update.Message.NewChatMembers) > 0 {
		for _, user := range update.Message.NewChatMembers {
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"👋 Welcome "+user.FirstName+" to the group!",
			)
			bot.Send(msg)
		}
		continue
	}

	// Anti-link filter
	if strings.Contains(update.Message.Text, "http://") ||
		strings.Contains(update.Message.Text, "https://") ||
		strings.Contains(update.Message.Text, "t.me/") {

		del := tgbotapi.NewDeleteMessage(
			update.Message.Chat.ID,
			update.Message.MessageID,
		)
		bot.Request(del)

		warn := tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"🚫 Links are not allowed.",
		)
		bot.Send(warn)

		continue
	}

	switch update.Message.Text {

	case "/start":
		bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Hello! I'm online ✅",
		))

	case "/ping":
		bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"Pong 🏓",
		))

	case "/rules":
		bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			"📜 Rules:\n1. No Spam\n2. No Links\n3. Respect Everyone",
		))

	case "/id":
		bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			fmt.Sprintf("Chat ID: %d", update.Message.Chat.ID),
		))
        case "/movie":
bot.Send(tgbotapi.NewMessage(
update.Message.Chat.ID,
"Movie search coming soon 🎬",
))

	case "/info":
		text := fmt.Sprintf(
			"👤 User: %s\n🆔 User ID: %d\n💬 Chat ID: %d",
			update.Message.From.FirstName,
			update.Message.From.ID,
			update.Message.Chat.ID,
		)

		bot.Send(tgbotapi.NewMessage(
			update.Message.Chat.ID,
			text,
		))
	}
}

}

