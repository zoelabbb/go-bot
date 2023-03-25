package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	boToken := "5886144265:AAHZ-JqJgdAC6aQSTQg4vxdMP9n1WwNrnuc"
	bot, err := tgbotapi.NewBotAPI(boToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Success connected on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			if strings.Contains(update.Message.Text, "Halo") || strings.Contains(update.Message.Text, "halo") || strings.Contains(update.Message.Text, "hai") || strings.Contains(update.Message.Text, "hi") {
				welcomeChat := fmt.Sprintf("Haii kak @%s 😉, ada yang bisa Go-bot bantu ?? \n", update.Message.From.UserName)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(welcomeChat))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)

			} else if strings.Contains(update.Message.Text, "kamu siapa") || strings.Contains(update.Message.Text, "km siapa") || strings.Contains(update.Message.Text, "km sp") {
				message, err := os.ReadFile("template/ikam.txt") // Membaca file ikam.txt dari folder template
				if err != nil {
					log.Fatal(err)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "kata kunci") {
				message, err := os.ReadFile("template/kunci.txt") // Membaca file ikam.txt dari folder template
				if err != nil {
					log.Fatal(err)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(message))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "bisa apa") {
				welcomeChat := fmt.Sprintf("Aku bisa menghitung kak @%s \n", update.Message.From.UserName)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(welcomeChat))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "calon suami aku") || strings.Contains(update.Message.Text, "suami aku") {
				welcomeChat := fmt.Sprintf("Iyaa.. iyaa.. kak @%s, dia punya kakak 😩. Galak banget deh \n", update.Message.From.UserName)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(welcomeChat))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else if strings.Contains(update.Message.Text, "kiss") || strings.Contains(update.Message.Text, "ayo kiss") {
				welcomeChat := fmt.Sprintf("😘 \n")

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(welcomeChat))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			} else {
				welcomeChat := fmt.Sprintf("Maaf kak @%s, Go-bot ga paham 😞 \n", update.Message.From.UserName)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(welcomeChat))
				msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
		}
	}
}
