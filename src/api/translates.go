package main

var Translates = map[string]map[string]string{
	"Tapped" : {
		"en" : "Tapped",
		"ru" : "Готово",
	},
	"Trying to remove notice by wrong ID.\nTry again later" : {
		"en" : "Trying to remove notice by wrong ID.\nTry again later",
		"ru" : "Попытка удаления заметки по неверному ID.\nПопробуйте позже",
	},
	"Cannot remove notice with ID: " : {
		"en" : "Cannot remove notice with ID: ",
		"ru" : "Не могу удалить заметку с ID: ",
	},
	".\nTry again later" : {
		"en" : ".\nTry again later",
		"ru" : ".\nПовторите еще раз позже",
	},
	"Removed" : {
		"en" : "Removed",
		"ru" : "Удалено",
	},
	"Remove" : {
		"en" : "Remove",
		"ru" : "Удалить",
	},
	"Author" : {
		"en" : "Author",
		"ru" : "Автор",
	},
	"Tap me" : {
		"en" : "Tap me",
		"ru" : "Нажми меня",
	},
	"Welcome to shopmaker bot! Menu:\n" : {
		"en" : "Welcome to shopmaker bot! Menu:\n",
		"ru" : "Добро пожаловать в shopmaker bot! Меню:\n",
	},
	"about Shopmaker bot\n" : {
		"en" : "about Shopmaker bot\n",
		"ru" : "о боте Shopmaker\n",
	},
	"save a notice\n" : {
		"en" : "save a notice\n",
		"ru" : "сохранить заметку\n",
	},
	"read a saved notice" : {
		"en" : "read a saved notice",
		"ru" : "прочитать сохраненную заметку",
	},
	"Shopmake Bot (c) April, 2016" : {
		"en" : "Shopmake Bot (c) April, 2016",
		"ru" : "Shopmake Bot (c) Апрель, 2016",
	},
	"Cannot save the notice. Try again later" : {
		"en" : "Cannot save the notice. Try again later",
		"ru" : "Не могу сохранить заметку, попробуйте позже",
	},
	"Notice saved! Notice ID: " : {
		"en" : "Notice saved! Notice ID: ",
		"ru" : "Заметка сохранена! ID заметки: ",
	},
	"Cannot save empty notice.\nYou can use command \"/saveNotice some text\" to save notice \"some text\"" : {
		"en" : "Cannot save empty notice.\nYou can use command \"/saveNotice some text\" to save notice \"some text\"",
		"ru" : "Не могу сохранить пустую заметку.\nВы можете использовать команду \"/saveNotice какой-то текст\" для сохранения заметки \"какой-то текст\"",
	},
	"Cannot find notice with ID: " : {
		"en" : "Cannot find notice with ID: ",
		"ru" : "Не могу найти заметку с ID: ",
	},
	".\nYou can use command \"/readNotice 1\" to read notice with ID = 1" : {
		"en" : ".\nYou can use command \"/readNotice 1\" to read notice with ID = 1",
		"ru" : ".\nВы можете использовать команду \"/readNotice 1\" для чтения заметки с ID = 1",
	},
	"Your notice with ID " : {
		"en" : "Your notice with ID ",
		"ru" : "Ваша заметка с ID ",
	},
	"Trying to read notice by wrong ID.\nYou can use command \"/readNotice 1\" to read notice with ID = 1" : {
		"en" : "Trying to read notice by wrong ID.\nYou can use command \"/readNotice 1\" to read notice with ID = 1",
		"ru" : "Попытка найти заметку по неверному ID.\nВы можете использовать команду \"/readNotice 1\" для чтения заметки с ID = 1",
	},
	"Language" : {
		"en" : "Language",
		"ru" : "Язык",
	},
	"Installed" : {
		"en" : "Installed",
		"ru" : "Установлен",
	},
	"english" : {
		"en" : "english",
		"ru" : "английский",
	},
	"russian" : {
		"en" : "russian",
		"ru" : "русский",
	},
	"English" : {
		"en" : "English",
		"ru" : "Английский",
	},
	"Russian" : {
		"en" : "Russian",
		"ru" : "Русский",
	},
	"language" : {
		"en" : "language",
		"ru" : "язык",
	},
	"Help" : {
		"en" : "Help",
		"ru" : "Помощь",
	},
}

func GetTranslate(key string, lang string) string {
	return Translates[key][lang]
}