package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/botfaces"
	"github.com/turnage/graw/reddit"

	"github.com/rrsharpe/Go-SSD-Bot/dispatcher"
	"github.com/rrsharpe/Go-SSD-Bot/ssd"
)

func main() {
	fmt.Println("Starting bot...")

	bot, err := reddit.NewBot(BotConfig())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	subreddit, subredditDefined := os.LookupEnv("SUBREDDIT")
	if !subredditDefined {
		panic("Subreddit not configured")
	}

	cfg := graw.Config{Subreddits: []string{subreddit}}
	ssdHandler := ssd.InitSSDPostHandler(bot)

	combinedHandler := dispatcher.GetUnionType(ssdHandler)
	for {
		// Fetch last few posts if something was missed during a restart
		readPrevious(bot, subreddit, combinedHandler)

		// Listen for new posts
		if _, wait, err := graw.Run(combinedHandler, bot, cfg); err != nil {
			fmt.Println("Failed to start graw run: ", err)
		} else {
			fmt.Println("graw run failed: ", wait())
		}
		fmt.Println("Sleeping for 5 minutes...")
		time.Sleep(5 * time.Minute)
	}
}

// Reads the previous 10 posts and applies the PostHandler method
func readPrevious(bot reddit.Bot, subreddit string, handler botfaces.PostHandler) {
	r_sub := fmt.Sprintf("/r/%s", subreddit)
	harvest, err := bot.Listing(r_sub, "")
	if err != nil {
		fmt.Println("Failed to fetch", r_sub, err)
		return
	}
	for _, post := range harvest.Posts[:int(math.Min(10, float64(len(harvest.Posts))))] {
		handler.Post(post)
	}
}
