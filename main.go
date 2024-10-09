package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

var (
	destinations        []string
	message             string
	jobResults          string
	fallbackDestination string
	replaceRef          bool
	gha                 githubactions.Action
)

func init() {
	gha = *githubactions.New()

	destinations = ParseDestinations(gha.GetInput("destination"))
	message = gha.GetInput("message")
	jobResults = gha.GetInput("results")

	replaceRef = ParseBool(gha.GetInput("remove-branch-prefix"))
	fallbackDestination = gha.GetInput("fallback-destination")
}

func main() {
	if replaceRef {
		message = strings.TrimPrefix(message, "refs/heads/")
	}

	fmt.Printf("Hello, %s! \n%s \n\n%s\n", destinations, message, jobResults)

	_, present := os.LookupEnv("INPUT_SLACK-TOKEN")
	gha.Infof("INPUT_SLACK-TOKEN env variable present: %t\n\n", present)

	client, err := NewClient()
	if err != nil {
		gha.Fatalf("Could not create client. Error: %+v \n", err)
	}

	bot := &Bot{client}
	_ = bot
	if err != nil {
		gha.Fatalf("Error %+v: \n", err)
	}

	_, err = bot.TestAuth()
	if err != nil {
		gha.Fatalf("Unable to authenticate. Check your .slack_token file. Error: %+v\n", err)
	}

	parsedResults, err := ParseJobResults(jobResults)
	if err != nil {
		gha.Fatalf("Unable to parse job results. Error: %+v\n", err)
	}

	if parsedResults != "" {
		message = message + "\n\n" + parsedResults
	}

	var useFallback bool = false
	for _, destination := range destinations {
		if destination == "" {
			continue
		}

		if destination == "committer" {
			destination = gha.GetInput("committer-email")

			gha.Infof(destination)
		}

		destination = strings.TrimSpace(destination)

		err = bot.PostMessage(destination, message)
		if err != nil {
			gha.Warningf("Oh no! We can't post a message to %s! %+v", destination, err)
			useFallback = true
		}
	}

	// When committing using the private email, we may need to fall back
	if useFallback {
		fallbackDestination = strings.TrimSpace(fallbackDestination)
		if fallbackDestination == "" {
			gha.Fatalf("No fallback destination given (`fallback-destination`) and one or more original destinations failed. Exiting.")
		}
		gha.Warningf("Using fallback destination: %s", fallbackDestination)

		err = bot.PostMessage(fallbackDestination, message)

		if err != nil {
			gha.Fatalf("Oh no! We can't post a message to %s! %+v", fallbackDestination, err)
		}
		gha.Infof("Fallback destination succeeded.")
	}

	message = fmt.Sprintf("Message sent to %s.\n", destinations)
	gha.SetOutput("validate-output", message)
}
