package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Bot ...
type Bot struct {
	APIKey string
	bot    *tgbotapi.BotAPI
	chanID int64
}

// NewBot ...
func NewBot(key string, id int64) *Bot {
	return &Bot{
		APIKey: key,
		chanID: id,
	}
}

// Start ...
func (b *Bot) Start() {
	bot, err := tgbotapi.NewBotAPI(b.APIKey)
	if err != nil {
		log.Panic(err)
	}

	b.bot = bot

	log.Printf("Authorized on account %s", bot.Self.UserName)
}

// Notify ...
func (b *Bot) Notify(text string) {
	msg := tgbotapi.NewMessage(b.chanID, text)

	b.bot.Send(msg)
}
