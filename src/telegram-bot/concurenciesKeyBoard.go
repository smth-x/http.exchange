package telegram_bot

import (
	tu "github.com/mymmrac/telego/telegoutil"
)

var KeyBoard = tu.InlineKeyboard(
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("USD : 🇺🇸 United States Dollar").WithCallbackData("USD"), tu.InlineKeyboardButton("EUR : 🇪🇺 Euro").WithCallbackData("EUR")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("UAH : 🇺🇦 Ukrainian Hryvnia").WithCallbackData("UAH"), tu.InlineKeyboardButton("PLN : 🇵🇱 Polish Zloty").WithCallbackData("PLN")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("GBP : 🇬🇧 British Pound Sterling").WithCallbackData("GBP"), tu.InlineKeyboardButton("CZK : 🇨🇿 Czech Koruna").WithCallbackData("CZK")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("HUF : 🇭🇺 Hungarian Forint").WithCallbackData("HUF"), tu.InlineKeyboardButton("JPY : 🇯🇵 Japanese Yen").WithCallbackData("JPY")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("NOK : 🇳🇴 Norwegian Krone").WithCallbackData("NOK"), tu.InlineKeyboardButton("KRW : 🇰🇷 South Korean Won").WithCallbackData("KRW")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("CHF : 🇨🇭 Swiss Franc").WithCallbackData("CHF"), tu.InlineKeyboardButton("AUD : 🇦🇺 Australian Dollar").WithCallbackData("AUD")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("SEK : 🇸🇪 Swedish Krona").WithCallbackData("SEK"), tu.InlineKeyboardButton("DKK : 🇩🇰 Danish Krone").WithCallbackData("DKK")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("CAD : 🇨🇦 Canadian Dollar").WithCallbackData("CAD"), tu.InlineKeyboardButton("CNY : 🇨🇳 Chinese Yuan").WithCallbackData("CNY")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("NZD : 🇳🇿 New Zealand Dollar").WithCallbackData("NZD"), tu.InlineKeyboardButton("MXN : 🇲🇽 Mexican Peso").WithCallbackData("MXN")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("SGD : 🇸🇬 Singapore Dollar").WithCallbackData("SGD"), tu.InlineKeyboardButton("KZT : 🇰🇿 Kazakh tenge").WithCallbackData("KZT")),
	tu.InlineKeyboardRow(tu.InlineKeyboardButton("TRY : 🇹🇷 Turkish Lira").WithCallbackData("TRY"), tu.InlineKeyboardButton("ILS : 🇮🇱 Israeli Shekel").WithCallbackData("ILS")),
)
