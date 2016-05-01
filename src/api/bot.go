package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"github.com/go-gorp/gorp"
	"strconv"
	"strings"
	"time"
)

func botRouter(bot *tgbotapi.BotAPI, postgres *gorp.DbMap)  {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, _ := bot.GetUpdatesChan(u)
	for update := range updates {
		log.Println("")
		switch true {
		case (update.CallbackQuery != nil):handleCallbackQuery(bot, update, postgres)
		case update.InlineQuery != nil: handleInlineQuery(bot, update)
		case update.Message != nil: handleMessage(bot, update, postgres)
		}
	}
}

func handleCallbackQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update, postgres *gorp.DbMap)  {
	lang := checkUserGetLang(update.CallbackQuery.From.ID, postgres)
	switch update.CallbackQuery.Data  {
	case "set_russian" :
		err := setUserLang(update.CallbackQuery.From.ID, "ru", postgres)
		if err != nil {
			log.Printf("+%v\n", err)
		} else {
			lang = checkUserGetLang(update.CallbackQuery.From.ID, postgres)
			userAlert := GetTranslate("Installed", lang) + " " + GetTranslate("russian", lang) + " " + GetTranslate("language", lang)
			c := tgbotapi.NewCallback(update.CallbackQuery.ID, userAlert)
			bot.AnswerCallbackQuery(c)
		}
	case "set_english" :
		err := setUserLang(update.CallbackQuery.From.ID, "en", postgres)
		if err != nil {
			log.Printf("+%v\n", err)
		} else {
			lang = checkUserGetLang(update.CallbackQuery.From.ID, postgres)
			userAlert := GetTranslate("Installed", lang) + " " + GetTranslate("english", lang) + " " + GetTranslate("language", lang)
			c := tgbotapi.NewCallback(update.CallbackQuery.ID, userAlert)
			bot.AnswerCallbackQuery(c)
		}
	case "help" :
		msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Help page")
		bot.Send(msg)
		c := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
		c.ShowAlert = false
		bot.AnswerCallbackQuery(c)
	default:
		id, err := strconv.Atoi(update.CallbackQuery.Data)
		if err != nil {
			log.Printf("+%v\n", err)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Trying to remove notice by wrong ID.\nTry again later", lang))
			bot.Send(msg)
		} else {
			_, err = postgres.Exec("DELETE FROM notices WHERE id=$1", id)
			if err != nil {
				log.Printf("+%v\n", err)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Cannot remove notice with ID: ", lang) + update.CallbackQuery.Data + GetTranslate(".\nTry again later", lang))
				bot.Send(msg)
			} else {
				c := tgbotapi.NewCallback(update.CallbackQuery.ID, GetTranslate("Removed", lang))
				bot.AnswerCallbackQuery(c)
			}
		}
	}
}

func handleInlineQuery(bot *tgbotapi.BotAPI, update tgbotapi.Update)  {
	log.Println("INLINE REQUEST")
	log.Printf("%+v\n", update.InlineQuery)
	inlineConfig := new(tgbotapi.InlineConfig)
	inlineResult := tgbotapi.NewInlineQueryResultArticle(update.InlineQuery.ID, "Hello query", "some text")
	inlineConfig.InlineQueryID = update.InlineQuery.ID
	inlineConfig.Results = append(inlineConfig.Results, inlineResult)
	bot.AnswerInlineQuery(*inlineConfig)
}

func handleMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, postgres *gorp.DbMap) {
	log.Printf("\n[%s] %s \n", update.Message.From.UserName, update.Message.Text)
	lang := checkUserGetLang(update.Message.From.ID, postgres)
	command := update.Message.Command()
	switch command {
	case "start":
		var m tgbotapi.InlineKeyboardMarkup
		if update.Message.Chat.IsPrivate() {
			m = getStartKeyboardMarkup(lang, true)
		} else {
			m = getStartKeyboardMarkup(lang, false)
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Welcome to shopmaker bot! Menu:\n", lang) + "/about - " + GetTranslate("about Shopmaker bot\n", lang) + "/saveNotice - " + GetTranslate("save a notice\n", lang) + "/readNotice - " + GetTranslate("read a saved notice", lang))
		msg.ReplyMarkup = m
		bot.Send(msg)
	case "about":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Shopmake Bot (c) April, 2016", lang))
		bot.Send(msg)
	case "saveNotice":
		notice := new(TextNotice)
		notice.Message = strings.Replace(update.Message.Text, "/" + command + " ", "", -1)
		if (len(notice.Message) > 0) && (notice.Message != "/saveNotice") {
			notice.UserId = update.Message.From.ID
			err := postgres.Insert(notice)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Cannot save the notice. Try again later", lang))
				bot.Send(msg)
			} else {
				strId := strconv.Itoa(notice.Id)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Notice saved! Notice ID: ", lang) + strId)
				bot.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Cannot save empty notice.\nYou can use command \"/saveNotice some text\" to save notice \"some text\"", lang))
			bot.Send(msg)
		}
	case "readNotice":
		strId := strings.Replace(update.Message.Text, "/" + command, "", -1)
		strId = strings.Replace(strId, " ", "", -1)
		id, err := strconv.Atoi(strId)
		if (id != 0) && (err == nil) {
			notice := new(TextNotice)
			err = postgres.SelectOne(&notice, "SELECT * FROM public.notices WHERE id=$1", id)
			if (err != nil) || (notice.UserId != update.Message.From.ID){
				log.Printf("+%v\n", err)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Cannot find notice with ID: ", lang) + strId + GetTranslate(".\nYou can use command \"/readNotice 1\" to read notice with ID = 1", lang))
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Your notice with ID ", lang) + strId + ":\n" + notice.Message)
				buttonRemove := tgbotapi.NewInlineKeyboardButtonData(GetTranslate("Remove", lang), strId)
				var buttonsSet []tgbotapi.InlineKeyboardButton
				buttonsSet = append(buttonsSet, buttonRemove)
				if update.Message.Chat.IsPrivate() {
					m := tgbotapi.NewInlineKeyboardMarkup(buttonsSet)
					msg.ReplyMarkup = m
				}
				bot.Send(msg)
			}
		} else {
			log.Printf("+%v\n", err)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, GetTranslate("Trying to read notice by wrong ID.\nYou can use command \"/readNotice 1\" to read notice with ID = 1", lang))
			bot.Send(msg)
		}
	}
}

func checkUserGetLang(userId int, postgres *gorp.DbMap) (lang string) {
	user := new(User)
	user.Lang = "en"
	err := postgres.SelectOne(&user, "SELECT * FROM users WHERE id=$1", userId)
	if err != nil {
		log.Printf("+%v\n", err)
		user.Id = userId
		user.Lang = "en"
		user.CreatedAt = time.Now()
		err = postgres.Insert(user)
		if err != nil {
			log.Printf("+%v\n", err)
		}
	}
	lang = user.Lang
	return
}

func setUserLang(userId int, lang string, postgres *gorp.DbMap) error {
	user := new(User)
	err := postgres.SelectOne(&user, "SELECT * FROM users WHERE id=$1", userId)
	if err != nil {
		return err
	}
	if user.Lang != lang {
		user.Lang = lang
		_, err = postgres.Update(user)
	}
	return err
}

func getStartKeyboardMarkup(lang string, isPrivate bool) tgbotapi.InlineKeyboardMarkup {
	buttonAuthor := tgbotapi.NewInlineKeyboardButtonURL(GetTranslate("Author", lang), "https://denniselite.me");
	buttonHelp := tgbotapi.NewInlineKeyboardButtonData(GetTranslate("Help", lang), "help");
	var publicRow []tgbotapi.InlineKeyboardButton
	var privateRow []tgbotapi.InlineKeyboardButton
	publicRow = append(publicRow, buttonAuthor)
	publicRow = append(publicRow, buttonHelp)
	if isPrivate {
		buttonEnglishLanguage := tgbotapi.NewInlineKeyboardButtonData(GetTranslate("English", lang), "set_english")
		buttonRussianLanguage := tgbotapi.NewInlineKeyboardButtonData(GetTranslate("Russian", lang), "set_russian")
		var buttonsSetPrivate []tgbotapi.InlineKeyboardButton
		buttonsSetPrivate = append(buttonsSetPrivate, buttonEnglishLanguage)
		buttonsSetPrivate = append(buttonsSetPrivate, buttonRussianLanguage)
		privateRow = buttonsSetPrivate
	}
	if len(privateRow) > 0 {
		return tgbotapi.NewInlineKeyboardMarkup(privateRow, publicRow)
	} else {
		return tgbotapi.NewInlineKeyboardMarkup(publicRow)
	}
}