package network

import (
	"calcobot/internal/database"
	"calcobot/internal/model"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type bot struct {
	api     *tgbotapi.BotAPI
	isDebug bool
}

// Create new telegram bot with this token and with mode
// Returning bot and error
// Can produce TG API error
func NewBot(token string, isDebug bool) (bot, error) {
	botApi, err := tgbotapi.NewBotAPI(token)

	botApi.Debug = isDebug

	if err != nil {
		return bot{}, err
	}

	return bot{botApi, isDebug}, nil
}

func (bot bot) getUpdatesChan() tgbotapi.UpdatesChannel {

	config := tgbotapi.NewUpdate(0) // offset
	config.Timeout = 60

	return bot.api.GetUpdatesChan(config) //todo сделать свою обёртку
}

func (bot bot) sendMessage(chatId int64, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatId, text)

	return bot.api.Send(msg)

}

func (bot bot) replyToMessage(chatId int64, messageId int, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyToMessageID = messageId

	return bot.api.Send(msg)
}

// Start bot polling with this database
func (bot bot) StartWorking(db database.Database) {
	updates := bot.getUpdatesChan()

	fmt.Println("Telegram bot is working")
	for update := range updates {
		if update.Message == nil {
			continue
		}

		request := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				_, err := bot.replyToMessage(request.ChatID, update.Message.MessageID,
					"Этот бот предназначен для вычисления арифметических выражений.\n\n"+
						"Вы можете написать любой пример я его посчитаю!\n"+
						"Я знаю следующие операции: +, -, *, /, ^ (возведение в степень), ! (факториал), ~ (унарный минус).\n\n"+
						"Просто напишите выражение с числами и я скажу вам ответ!\n")

				fmt.Println(err)
			case "meow":
				_, err := bot.replyToMessage(request.ChatID, update.Message.MessageID,
					"Мяу")
				if err != nil {
					fmt.Println(err)
				}
			default:
				_, err := bot.replyToMessage(request.ChatID, update.Message.MessageID,
					"Неизвестная команда")

				if err != nil {
					fmt.Println(err)
				}
			}
		} else {
			if !isExpression(request.Text) {
				bot.replyToMessage(request.ChatID, update.Message.MessageID,
					"Я не смог распознать ваш пример. Возможно, что вы заполнили его неверно."+
						"Используйте /start, чтобы узнать правила заполнения")
				continue
			}

			postfixExpression, err := model.ToPostfix(request.Text)

			if err != nil {
				bot.replyToMessage(request.ChatID, update.Message.MessageID,
					"Я не смог распознать ваш пример. Возможно, что вы заполнили его неверно."+
						"Используйте /start, чтобы узнать правила заполнения")

				if err != nil {
					fmt.Println(err)
				}
				continue
			}

			answer, err := model.CalculatePostfix(postfixExpression)

			if err != nil {
				bot.replyToMessage(request.ChatID, update.Message.MessageID, "Ошибка: "+err.Error())

				if err != nil {
					fmt.Println(err)
				}
				continue
			}

			log := database.Log{Time: time.Now(), Username: update.Message.From.UserName,
				Request: update.Message.Text, Answer: answer}
			db.AddLog(log)

			_, err = bot.replyToMessage(request.ChatID, update.Message.MessageID,
				"Ответ на ваш пример: "+strconv.FormatFloat(answer, 'f', 4, 64))
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

// Check that this user input is expression
func isExpression(input string) bool {
	regex, _ := regexp.Compile(`[a-zA-Zа-яА-Я;\\]+`)

	results := regex.FindAllString(input, 1)

	return len(results) == 0
}
