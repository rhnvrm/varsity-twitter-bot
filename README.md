# varsity-twitter-bot

I am a [bot](https://twitter.com/VarsityBot) that tweets key takeaways from Zerodha Varsity modules every 15 mins.

## Development Docs

The bot resides in `bot.go` and can be built with `go build -o bot.bin .`

### Deployment

The bot can be deployed by setting a cron job in your server.

`*/15 * * * * cd $HOME/apps/varsity-bot && ./bot.bin`

It depends on the following:

### .env File

A .env file or environment variables set in the shell containing the following:

```
CONSUMER_KEY=xxx
CONSUMER_SECRET=xxx
ACCESS_TOKEN=xxx
ACCESS_SECRET=xxx
```

You can use `login.py` with your consumer key to obtain the access token for the account you wish to tweet from.

### data.txt

The format of the data.txt file is 
```
<url of module>
<text>
.
.
<text>
```

Where each `<text>` is considered a single tweet and <url of module> is the link that will be attached to the tweet.

In case the `<text>` is longer that max tweet length, its split into multiple tweets as a thread.