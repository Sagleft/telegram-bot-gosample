package main

import (
	"log"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	settingsFilePath = "settings.json"
)

type botSettings struct {
	Token string `json:"token"`
}

func main() {
	settings, err := getBotSettings()
	if err != nil {
		log.Fatal(err)
		return
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  settings.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Hello World!")
	})

	b.Start()
}
