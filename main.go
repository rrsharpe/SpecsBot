package main

import (
	"fmt"
	"math"
	"time"

	"github.com/turnage/graw"
	"github.com/turnage/graw/botfaces"
	"github.com/turnage/graw/reddit"

	"github.com/rrsharpe/Go-SSD-Bot/ssd"
)

func main() {
	fmt.Println("Starting bot...")

	bot, err := reddit.NewBot(BotConfig())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	subreddits := []string{"SpecsBotTesting", "bapcsalescanada"}
	// subreddits := []string{"SpecsBotTesting"}

	cfg := graw.Config{Subreddits: subreddits}
	handler := ssd.InitSSDPostHandler(bot)
	for {
		// Fetch last few posts if something was missed during a restart
		readPrevious(bot, subreddits, handler)

		// Listen for new posts
		if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
			fmt.Println("Failed to start graw run: ", err)
		} else {
			fmt.Println("graw run failed: ", wait())
		}
		fmt.Println("Sleeping for 5 minutes...")
		time.Sleep(5 * time.Minute)
	}
}

// Reads the previous 10 posts and applies the PostHandler method
func readPrevious(bot reddit.Bot, subreddits []string, handler botfaces.PostHandler) {
	for _, subreddit := range subreddits {
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
}
