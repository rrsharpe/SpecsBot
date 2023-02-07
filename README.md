# SpecsBot

This project is the source code for [/u/SpecsBot](https://www.reddit.com/user/SpecsBot/) that currently runs in [/r/bapcsalescanada](https://www.reddit.com/r/bapcsalescanada/). It was inspired by a [similar project](https://github.com/ocmarin/ssd-bot) but plans to add more features.

---

## Running the bot
You must have `go` version 1.19 or greater installed to build and run. You can find the latest installer for your system [here](https://go.dev/dl/).

To build the application, run the following command to create a binary for your system.

```sh
go build ./...
```

Alternatively you can use the Dockerfile to build this on any system.
```sh
docker build .
```

To run the application you need to set the following environment variables
 * BOT_AGENT - A user agent for your bot as defined [here](https://github.com/reddit-archive/reddit/wiki/API#rules)
 * BOT_ID - The [Reddit OAuth2](https://github.com/reddit-archive/reddit/wiki/OAuth2) `client_id`
 * BOT_SECRET - The [Reddit OAuth2](https://github.com/reddit-archive/reddit/wiki/OAuth2) `secret`
 * BOT_USERNAME - The username of your bot
 * BOT_PASSWORD - The password of the bot account

 ---


## Project Structure
This project is separated into packages. The main package simply fetches and calls the various post handlers from each child module. Whenever a post from a followed subreddit is detected, each module's post handler will be called. It is up to each post handler to only respond to their type of post.

---

## License
This project's code is licensed under the MIT license. However, some of the resources used by this bot such as spreadsheets are not. Therefore, while the source code of this project is licensed under MIT, the comments created by this bot and the resources it dynamically fetches at run time may not be MIT licensed.