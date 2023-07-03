package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
	"log"
	"reflect"
	_ "reflect"
	"strconv"
	"tg_bot_tuck/bd"
	"tg_bot_tuck/commands"
)

func main() {

	var NumUpTusk int
	var IdUser int
	var NewUpperTusk string

	var AllTusks []bd.Task

	s, err := bd.ConnectDB()
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI("6226586282:AAGArHkmCYMKBod6Rjy4ISeHlCgjLzzOShw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//defer bd.Testdb.Md.Close()

	for update := range updates {
		if update.Message != nil { // If we got a message

			ms := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			ms4 := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			ms5 := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			ms7 := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			ms8 := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			ms9 := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			switch update.Message.Text {
			case commands.Start:
				ms.Text = commands.StartMessage
				bot.Send(ms)
			case commands.Add:
				ms.Text = commands.AddMessage
				ms.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, InputFieldPlaceholder: commands.AddMessage, Selective: false}
				bot.Send(ms)
			case commands.ShowNote:
				AllTusks = bd.ShowNote(s, update.Message.From.ID)
				values := reflect.ValueOf(AllTusks)
				typesOf := values.Type()
				fmt.Println(typesOf)
				for _, v := range AllTusks {

					t := strconv.Itoa(v.Id)
					r := v.Task
					ms4 = tgbotapi.NewMessage(update.Message.Chat.ID, "Заметка №:"+t+"\n"+r)
					bot.Send(ms4)
				}

			case commands.Remove:
				ms5.Text = commands.RemoveMessage
				ms5.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, InputFieldPlaceholder: commands.RemoveMessage, Selective: false}
				bot.Send(ms5)

			case commands.UpdateNote:
				ms7.Text = commands.UpdateMessage
				ms7.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, InputFieldPlaceholder: commands.UpdateMessage, Selective: false}
				bot.Send(ms7)
			}
			if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text != "null" {
				var currReplMessage2 = update.Message.ReplyToMessage.Text
				if "Отправьте номер заметки,которую треубуеться обновить" == currReplMessage2 {
					NumUpTusk, _ = strconv.Atoi(update.Message.Text)
					fmt.Println(NumUpTusk)
					ms8.Text = commands.NewUpMessage
					ms8.ReplyMarkup = tgbotapi.ForceReply{ForceReply: true, InputFieldPlaceholder: commands.NewUpMessage, Selective: false}
					bot.Send(ms8)
				}

			}
			if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text != "null" {
				var curReplMessage3 = update.Message.ReplyToMessage.Text
				if "Отправьте ваши изменения" == curReplMessage3 {
					NewUpperTusk = update.Message.Text
					IdUser = int(update.Message.From.ID)
					ms9.Text = commands.SuccessfullyUpdated
					bd.UpdateNote(s, int64(IdUser), int64(NumUpTusk), NewUpperTusk)
					bot.Send(ms9)
				}

			}

			if update.Message.ReplyToMessage != nil && update.Message.ReplyToMessage.Text != "null" {
				var currentReplyText = update.Message.ReplyToMessage.Text
				if "Отправьте номер заметки, которую хотите удалить" == currentReplyText {
					var ms6 = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					if ms6.Text == "/remove" {
						continue
					}
					userId := update.Message.From.ID
					var taskNumber, _ = strconv.Atoi(update.Message.Text)
					bd.RemoveNote(s, int(userId), taskNumber)
					ms6.Text = "Задача " + update.Message.Text + " Успешно удалена"
					bot.Send(ms6)
				} else if "Отправьте вашу заметку" == currentReplyText {

					ms3 := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
					if ms3.Text == "/add" {
						continue
					}
					bd.InsertDB(s, update.Message.From.ID, update.Message.Text)

					ms3.Text = "Задача: " + update.Message.Text + " успешно добавленна в список задач на сегодня"

					bot.Send(ms3)
				}
			}

		}
	}

}
