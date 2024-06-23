package main

import (
	"fmt"
	"os"
	"log"

	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading env file: %v", err)
	}
	log.Println("Env loaded successfully")
}

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"SlackFileBot.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}