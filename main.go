package main

import (
	"fmt"
	"os"
	"strings"

	gjs "github.com/gopherjs/gopherjs/js"
	githubactions "github.com/sethvargo/go-githubactions"
)

var (
	destinations        []string
	message             string
	jobResults          string
	gitRepo             string
	gitToken            string
	gitSha              string
	replaceRef          string
	fallbackDestination string
)

func init() {
	destinations = ParseDestinations(githubactions.GetInput("destination"))
	message = githubactions.GetInput("message")
	jobResults = githubactions.GetInput("results")
	gitRepo = githubactions.GetInput("github-repository")
	if gitRepo == "" {
		gitRepo = os.Getenv("GITHUB_REPOSITORY")
	}
	gitToken = githubactions.GetInput("github-token")
	gitSha = githubactions.GetInput("github-sha")
	if gitSha == "" {
		gitSha = os.Getenv("GITHUB_SHA")
	}
	replaceRef = githubactions.GetInput("remove-branch-prefix")
	fallbackDestination = githubactions.GetInput("fallback-destination")

}

func main() {
	if replaceRef == "true" {
		message = strings.ReplaceAll(message, "refs/heads/", "")
	}

	fmt.Printf("Hello, %s! \n%s \n\n%s\n", destinations, message, jobResults)

	_, present := os.LookupEnv("INPUT_SLACK-TOKEN")
	githubactions.Debugf("INPUT_SLACK-TOKEN env variable present: %t\n\n", present)

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

	parsedResults, err := ParseJobResults(jobResults)
	if err != nil {
		githubactions.Errorf("Unable to parse job results. Error: %+v\n", err)
		gjs.Global.Call("ExitAndFail", 4)
	}

	if parsedResults != "" {
		message = message + "\n\n" + parsedResults
	}

	useFallback := false
	for _, destination := range destinations {
		if destination == "" {
			continue
		}

		if destination == "committer" {
			email, err := GetCommitEmail(gitRepo, gitSha, gitToken)
			if err != nil {
				githubactions.Errorf("Error %+v: \n", err)
				useFallback = true
			}
			destination = email
			fmt.Println(email)
		}

		err = bot.PostMessage(strings.Trim(destination, " "), message)

		if err != nil {
			githubactions.Errorf("Oh no! We can't post a message to %s! %+v", destination, err)
		}
	}

	// When committing using the private email, we may need to fall back
	if useFallback {
		destinations = append(destinations, fallbackDestination)
		err = bot.PostMessage(strings.Trim(fallbackDestination, " "), message)

		if err != nil {
			githubactions.Errorf("Oh no! We can't post a message! %+v", err)
			gjs.Global.Call("ExitAndFail", 4)
		}
	}

	message = fmt.Sprintf("Message sent to %s.\n", destinations)
	githubactions.SetOutput("validate-output", message)
}
