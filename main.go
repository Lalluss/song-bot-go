package main

import (
"log"
"os"

tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
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
