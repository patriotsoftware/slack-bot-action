package main

import (
	"fmt"
	"os"

	gjs "github.com/gopherjs/gopherjs/js"
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
		githubactions.Errorf("Could not create client. Error: %+v \n", err)
		gjs.Global.Call("ExitAndFail", 1)
	}

	bot := &Bot{client}
	_ = bot
	if err != nil {
		githubactions.Errorf("Error %+v: \n", err)
		gjs.Global.Call("ExitAndFail", 2)
	}

	_, err = bot.TestAuth()
	if err != nil {
		githubactions.Errorf("Unable to authenticate. Check your .slack_token file. Error: %+v\n", err)
		gjs.Global.Call("ExitAndFail", 3)
	}

	_ = bot.PostMessage(destination, message)

	if err != nil {
		githubactions.Errorf("Oh no! We can't post a message! %+v", err)
		gjs.Global.Call("ExitAndFail", 4)
	}

	message = fmt.Sprintf("Message sent to %s.\n", destination)
	githubactions.SetOutput("validate-output", message)
}
