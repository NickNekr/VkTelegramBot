package main

import (
	"VkTelegramBot/helper"
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
)

type Bot struct {
	botToken string
	baseUrl  string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

}

func main() {

	bot := Bot{os.Getenv("TOKEN"), os.Getenv("BASE_URL")}
	if bot.botToken == "" {
		log.Fatal("TOKEN is not set")
	}
	bot.infinityPolling()
}

func (bot *Bot) infinityPolling() {
	lastMessageId := 0
	for {
		get, err := http.Get(bot.baseUrl + bot.botToken + "/getUpdates?offset=-1")
		if err != nil || get == nil {
			return
		}

		var u helper.Response
		err = json.NewDecoder(get.Body).Decode(&u)
		if !u.Ok || err != nil {
			log.Fatal("Bad get Result!")
		}

		if len(u.Result) == 0 || lastMessageId == u.Result[0].UpdateID {
			continue
		}

		lastMessageId = u.Result[0].UpdateID
		var chatId int
		var messageText string
		if u.Result[0].CallbackQuery != nil {
			chatId = u.Result[0].CallbackQuery.From.ID
			messageText = u.Result[0].CallbackQuery.Data
		} else {
			chatId = u.Result[0].Message.From.ID
			messageText = u.Result[0].Message.Text
		}

		switch {
		case messageText == "/start":
			go bot.sayHello(helper.BodyMessage{ChatId: chatId, Text: helper.WELCOME_MESSAGE,
				ReplyMarkup: helper.Markup,
			})
		case strings.Contains(messageText, "Button"):
			go bot.sayHello(helper.BodyMessage{ChatId: chatId, Text: helper.PRESSED_MESSAGE + messageText, ReplyMarkup: helper.InlineKeyboardMarkup{[][]helper.InlineKeyboardButton{}}})
		}
	}
}

func (bot *Bot) sayHello(bodyMessage helper.BodyMessage) {
	url := bot.baseUrl + bot.botToken + "/sendMessage"
	body, err := json.Marshal(bodyMessage)
	request, err := http.Post(url, "application/json", bytes.NewBuffer(body))

	defer request.Body.Close()
	var mR helper.MessageResponse

	err = json.NewDecoder(request.Body).Decode(&mR)
	if !mR.Ok || err != nil {
		log.Fatal("Bad get Result!")
	}
	log.Println("Sended message!")
}
