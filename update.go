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

func Update(number int, repo string, data map[string]string, token string) (*Issue, error) {
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

	url := strings.Join([]string{Url, repo, "issues", fmt.Sprint(number)}, "/")
	// url := fmt.Sprintf("%s/%s/issues/%d", Url, repo, number)
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	request.Header.Add("Authorization", token)
	request.Header.Add("accept", "application/vnd.github.v3+json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error\t%d", resp.StatusCode)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &result, nil
}