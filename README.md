# Slack Bot Action

A GitHub Action for sending alerts on github actions to slack.

## Example Usage

```yaml
- uses: synergydatasystems/slackbot@v1
  with:
    destination: committer
    message: "A new slackbot update has been triggered."
    slack-token: ${{ secrets.SLACK_TOKEN }}
    github-token: ${{ secrets.GITHUB_TOKEN }}
    fallback-destination: "#channel-if-any-destination-fails"
```

## Mapping GitHub users to Slack users

Located in `.github/user_mapping.yml` is a key value pair yaml file that needs to contain a `github_username: slack_id` for every user who commits to our repos. This ensures the `committer` destination always finds the correct slack user.

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
remove-branch-prefix:
  Removes branch /refs/head/ from string instances of /refs/heads. Values: true/false. Allows use of ${{ github.ref }} to print without /refs/heads
  Default: true
fallback-destination:
  The channel (in the format `#channel`) used for when any of the previous destinations fail.
  Set to "disabled" to skip a fallback
```

## Outputs

```yaml
validate-output:
  This is for verification that the code was run successfully.
```
