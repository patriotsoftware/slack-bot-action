# PatriotSoftware.GHA.SlackBot

A GitHub Action for sending alerts on github actions to slack. We recommend using patriotsoftware/slackbot@v1 to get the latest changes. If new features require breaking changes, we will release them to @v2. You can also use a full semantic version tag.

## Example Usage

```yaml
- uses: patriotsoftware/slackbot@v1
  with:
    destination: committer
    message: "A new slackbot update has been triggered."
    slack-token: ${{ secrets.SLACK_TOKEN }}
    github-token: ${{ secrets.GITHUB_TOKEN }}
    fallback-destination: "#channel-if-any-destination-fails"
```

## Inputs

```yaml
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
results:
  List of job results to append to the message. Example:
    job-results: |
          Job 1:${{needs.job-one.result}}
          Job 2:${{needs.other-job.result}}
github-token:
  GitHub Repository to for getting commit email.
  Use this most times: ${{ secrets.GITHUB_TOKEN }}"
github-repository:
  GitHub Repository to for getting commit email.
  Use this most times: ${{ github.repository }}"
github-sha:
  GitHub SHA to for getting commit email.
  Use this most times: ${{ github.sha }}
remove-branch-prefix:
    Removes branch /refs/head/ from string instances of /refs/heads. Values: true/false. Allows use of ${{ github.ref }} to print without /refs/heads
    Default: true
fallback-destination:
    The channel (in the format `#channel`) used for when any of the previous destinations fail.
```

## Outputs

```yaml
validate-output:
  This is for verification that the code was run successfully.
```

To test locally either run the test-action workflow with unique inputs, or run the main.go locally.

## Testing the action locally

- Create a .slack_token file with your slack token
- `export INPUT_MESSAGE="Your message"`
- `export INPUT_DESTINATION="#yourChannel"` or `"@your_user"`
- `go run main.go`

Nate was X
