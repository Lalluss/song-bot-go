package main

import (
	"log"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
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

		switch update.Message.Text {

		case "/start":
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Hello! I'm online ✅",
			)
			bot.Send(msg)

		case "/ping":
			msg := tgbotapi.NewMessage(
				update.Message.Chat.ID,
				"Pong 🏓",
			)
			bot.Send(msg)
		}
	}
}
