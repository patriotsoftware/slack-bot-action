# PatriotSoftware.GHA.SlackBot

A GitHub Action for sending alerts on github actions to slack. We recommend using patriotsoftware/slackbot@v1 to get the latest changes. If new features require breaking changes, we will release them to @v2. You can also use a full semantic version tag.

# Example Usage

```
- uses: patriotsoftware/slackbot@v1
```


# Inputs

```
destination:
  This is either the channel or email address tied to the user who will receive the direct message.
  If this is a channel it must begin with the '#'.
message:
  Message to send in string format.
slack-token:
  This is the slack token made in your slack account.
  For use in an action, add it as a secret within your repo.
  If you are running main.go directly, you can include a file named .slack_token and include the token there.
  See https://api.slack.com/authentication/token-types for more info.
```

# Outputs

```
validate-output:
  This is for verification that the code was run successfully.
```

To test locally either run the test-action workflow with unique inputs, or run the main.go locally.

# Testing the action locally
Running locally does not require gopherjs (it's just used in the action to speed up the startup time). 

- Create a .slack_token file with your slack token
- `export INPUT_MESSAGE="Your message"`
- `export INPUT_DESTINATION="#yourChannel"` or `"@your_user"`
- `go run main.go`