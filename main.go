package main

import (
	"fmt"
	"os"

	githubactions "github.com/sethvargo/go-githubactions"
)

func main() {
	destination := githubactions.GetInput("destination")
	message := githubactions.GetInput("message")

	fmt.Printf("Hello, %s! %s \n", destination, message)

	_, present := os.LookupEnv("INPUT_SLACK-TOKEN")
	fmt.Printf("INPUT_SLACK-TOKEN env variable present: %t\n", present)

	client, err := NewClient()
	if err != nil {
		panic("Could not create client")
	}

	bot := &Bot{client}
	_ = bot
	if err != nil {
		fmt.Printf("Error %+v: ", err)
		os.Exit(1)
	}

	_, err = bot.TestAuth()
	if err != nil {
		fmt.Printf("Unable to authenticate. Check your .slack_token file. Error: %+v\n", err)
		os.Exit(1)
	}

	_ = bot.PostMessage(destination, message)

	if err != nil {
		fmt.Printf("Oh no! We can't post a message! %+v", err)
	}

	message = fmt.Sprintf("Message sent to %s.\n", destination)
	fmt.Printf("::set-output name=validate-output::%s \n", message)
	githubactions.SetOutput("validate-output", message)
}
