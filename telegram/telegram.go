package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot structure
type Bot struct {
	APIKey string
	bot    *tgbotapi.BotAPI
	chanID int64
}

// NewBot init
func NewBot(key string, id int64) *Bot {
	return &Bot{
		APIKey: key,
		chanID: id,
	}
}

// Start a bot
func (b *Bot) Start() {
	bot, err := tgbotapi.NewBotAPI(b.APIKey)
	if err != nil {
		log.Fatal(fmt.Sprintf("Telegram bot: %s", err))
	}

	b.bot = bot

	log.Printf("Authorized on account %s", bot.Self.UserName)
}

// Notify to channel
func (b *Bot) Notify(text string) {
	msg := tgbotapi.NewMessage(b.chanID, text)

	b.bot.Send(msg)
}
