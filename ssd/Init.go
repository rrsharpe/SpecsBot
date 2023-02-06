package ssd

import (
	"github.com/turnage/graw/botfaces"
	"github.com/turnage/graw/reddit"
)

type SSDPostHandler struct {
	botfaces.PostHandler
	bot         reddit.Bot
	ssdModelMap map[modelKey][]string
}

func InitSSDPostHandler(bot reddit.Bot) *SSDPostHandler {
	return &SSDPostHandler{bot: bot, ssdModelMap: prepareProcessedData()}
}
