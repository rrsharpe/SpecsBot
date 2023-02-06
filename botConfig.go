package main

import (
	"os"

	"github.com/turnage/graw/reddit"
)

func BotConfig() reddit.BotConfig {
	agent, agentNotSet := os.LookupEnv("BOT_AGENT")
	id, idNotSet := os.LookupEnv("BOT_ID")
	secret, secretNotSet := os.LookupEnv("BOT_SECRET")
	username, usernameNotSet := os.LookupEnv("BOT_USERNAME")
	password, passwordNotSet := os.LookupEnv("BOT_PASSWORD")

	if agentNotSet && idNotSet && secretNotSet && usernameNotSet && passwordNotSet {
		app := reddit.App{
			ID:       id,
			Secret:   secret,
			Username: username,
			Password: password,
		}

		return reddit.BotConfig{
			Agent: agent,
			App:   app,
		}
	}
	panic("Some environment variables were not set.")
}
