package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

func botRouter(bot *tgbotapi.BotAPI)  {
	buttonAuthor := tgbotapi.NewInlineKeyboardButtonURL("autor", "https://denniselite.me");
	buttonTapMe := tgbotapi.NewInlineKeyboardButtonData("tap me", "tapped")
	var buttonsSet []tgbotapi.InlineKeyboardButton
	buttonsSet = append(buttonsSet, buttonAuthor)
	buttonsSetPrivate := make([]tgbotapi.InlineKeyboardButton, len(buttonsSet))
	buttonsSetPrivate = append([]tgbotapi.InlineKeyboardButton(nil), buttonsSet...)
	buttonsSetPrivate = append(buttonsSet, buttonTapMe)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		log.Println("")
		if ((update.CallbackQuery != nil) && (update.CallbackQuery.Data == "tapped")) {
			c := tgbotapi.NewCallback(update.CallbackQuery.ID, "tapped")
			bot.AnswerCallbackQuery(c)
		} else if ((update.InlineQuery != nil)) {
			log.Println("INLINE REQUEST")
			log.Printf("%+v\n", update.InlineQuery)
			inlineConfig := new(tgbotapi.InlineConfig)
			inlineResult := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Hello query", "some text")
			inlineConfig.InlineQueryID = update.InlineQuery.ID
			inlineConfig.Results = append(inlineConfig.Results, inlineResult)
			bot.AnswerInlineQuery(*inlineConfig)
		} else {
			m := tgbotapi.NewInlineKeyboardMarkup(buttonsSet)
			log.Printf("\n[%s] %s \n", update.Message.From.UserName, update.Message.Text)
			switch update.Message.Command() {
			case "start":
				if update.Message.Chat.IsPrivate() {
					m = tgbotapi.NewInlineKeyboardMarkup(buttonsSetPrivate)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to shopmaker bot! Menu:\n /about")
				msg.ReplyMarkup = m
				bot.Send(msg)
			case "about":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Shopmake Bot (c) April, 2016")
				bot.Send(msg)
			case "saveText":
			}
		}
	}
}