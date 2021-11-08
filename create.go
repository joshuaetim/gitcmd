package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func Create(repo string, data map[string]string, token string) (*Issue, error) {
	if repo == "" {
		return nil, errors.New("the repository field is required")
	}

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	requestBody, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	url := strings.Join([]string{Url, repo, "issues"}, "/")
	// url := fmt.Sprintf("%s/%s/issues", Url, repo)
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	request.Header.Add("Authorization", token)
	request.Header.Add("accept", "application/vnd.github.v3+json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("response error: %v", resp.StatusCode)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}