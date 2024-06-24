package tgbot

import (
	"fmt"
	"os"
	"strings"

	"github.com/defloppka/brawlify_parsebot/scraper"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func RunBot() {
	botToken := "7244145047:AAErj2SuyUPFc3U2uMXI4XQOGLangG2LeMI"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)
	bh, _ := th.NewBotHandler(bot, updates)
	defer bh.Stop()
	defer bot.StopLongPolling()

	bh.Handle(func(bot *telego.Bot, update telego.Update) {
		_, _, args := tu.ParseCommand(update.Message.Text)
		mapName := strings.Join(args, " ")
		mapInfo := scraper.GetMapInfo(mapName)
		ok, message := mapInfo.Display()


		if ok {
			bot.SendPhoto(
				tu.Photo(
					update.Message.Chat.ChatID(),
					tu.FileFromURL(mapInfo.Image),
				),
			)
		}
		bot.SendMessage(
			tu.Message(
				update.Message.Chat.ChatID(),
				message,
			),
		)
	}, th.CommandEqual("get"))

	bh.Start()
}