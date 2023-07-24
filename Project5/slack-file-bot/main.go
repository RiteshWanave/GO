package main

import (
	"fmt"
	"os"
	"github.com/slack-go/slack"
)

func main () {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5607464929509-5607775660231-ZIU83LEEe6nqGsQHNOSRhqZz")
	os.Setenv("CHANNEL_ID", "C05JB0XP4LR")

	api := slack.New(os.Getenv(("SLACK_BOT_TOKEN")))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Aptitude.pdf", "test.cpp"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name: %s, URL: %s \n", file.Name, file.URL)
	}
}