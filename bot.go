package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/slack-go/slack"
)

type Bot struct {
	Client *slack.Client
}

func NewClient() (*slack.Client, error) {

	token, ok := os.LookupEnv("INPUT_SLACK-TOKEN")
	if ok {
		fmt.Printf("Using slack token from env. \n")
		return slack.New(token), nil
	}

	if fileExists(".slack_token") {
		fileToken, err := os.ReadFile(".slack_token")
		if err != nil {
			return nil, err
		}
		fmt.Printf("Using slack token from .slack_token file. \n")
		return slack.New(string(fileToken)), nil
	}

	return nil, fmt.Errorf("no token found")

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func (b *Bot) TestAuth() (string, error) {
	authTestResult, err := b.Client.AuthTest()
	if err != nil {
		return "", err
	}

	return authTestResult.BotID, nil
}

func (b Bot) PostMessage(destination string, message string) error {

	if !strings.Contains(destination, "#") {
		userID, err := b.Client.GetUserByEmail(destination)

		if err != nil {
			return fmt.Errorf("unable to get user by email. Error: %+v", err)
		}
		destination = userID.ID
	}

	_, _, err := b.Client.PostMessage(
		destination,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	)

	if err != nil {
		return err
	}

	return nil
}
