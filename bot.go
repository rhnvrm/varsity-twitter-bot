package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

const (
	maxTweetLen = 280 - 25 // 25 char reserved for the url
)

func getTweet() ([]string, error) {
	// Read the file
	file, err := ioutil.ReadFile("data.txt")
	if err != nil {
		return nil, err
	}

	// get a random line from the file
	lines := strings.Split(string(file), "\n")
	rLine := ""
	lNo := 0

	for rLine == "" || strings.Contains(rLine, "http") {
		lNo = rand.Intn(len(lines))
		rLine = lines[lNo]
	}

	var outTweet []string
	for len(rLine) > 0 {
		if len(rLine) > maxTweetLen {
			outTweet = append(outTweet, rLine[:maxTweetLen-3]+"...")
			rLine = rLine[maxTweetLen-3:]
		} else {
			outTweet = append(outTweet, rLine)
			rLine = ""
		}
	}

	// from lno, find the nearest line with http
	for i := lNo; i < len(lines) && i >= 0; i-- {
		if strings.Contains(lines[i], "http") {
			outTweet[len(outTweet)-1] = outTweet[len(outTweet)-1] + "\n" + lines[i]
			break
		}
	}

	return outTweet, nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rand.Seed(time.Now().UnixNano())
	config := oauth1.NewConfig(os.Getenv("CONSUMER_KEY"), os.Getenv("CONSUMER_SECRET"))
	token := oauth1.NewToken(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_SECRET"))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	tweets, err := getTweet()
	if err != nil {
		log.Fatalf("Error getting tweet: %s", err)
	}

	lastID := int64(0)
	for _, txt := range tweets {
		var params *twitter.StatusUpdateParams
		if lastID != 0 {
			params = &twitter.StatusUpdateParams{InReplyToStatusID: lastID}
		}

		log.Println("sending tweet: ", txt)
		tweet, _, err := client.Statuses.Update(txt, params)
		if err != nil {
			log.Fatalf("Error tweeting: %s", err)
		}

		lastID = tweet.ID
	}

}
