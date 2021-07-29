package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type EmailResponse struct {
	Commit struct {
		Author struct {
			Email string
		}
	}
	Message string
}

func GetCommitEmail(repo string, sha string, token string) (string, error) {
	if repo == "" || sha == "" || token == "" {
		return "", fmt.Errorf("the GitHub token, Repository, and SHA are required")
	}

	client := &http.Client{}
	endpoint := fmt.Sprintf("https://api.github.com/repos/%s/commits/%s", repo, sha)
	token = "token " + token
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", token)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var email EmailResponse
	err = json.Unmarshal(body, &email)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return email.Message, err
	}
	return email.Commit.Author.Email, nil
}
