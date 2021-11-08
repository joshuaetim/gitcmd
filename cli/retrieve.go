package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"errors"
)

func Retrieve(number int, repo string) (*Issue, error) {
	if repo == "" {
		return nil, errors.New("the repository field is required")
	}
	
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	url := fmt.Sprintf("%s/%s/issues/%d", Url, repo, number)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	request.Header.Add("accept", "application/vnd.github.v3+json")

	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response error: %d", resp.StatusCode)
	}

	var result Issue
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	return &result, nil
}