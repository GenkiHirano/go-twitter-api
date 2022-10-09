package twitter

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func getTwitterApi() *anaconda.TwitterApi {
	loadEnv()
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_KEY_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	return api
}

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Failed to load environment variables: %v", err)
	}
}