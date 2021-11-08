package main

import (
	"flag"
	"fmt"
	"gopl/ch4/gitcmd/cli"
	"log"
	"os"
)

var updateFlag = flag.NewFlagSet("update", flag.ExitOnError)
var updateFlagNumber = updateFlag.Int("n", 0, "the issue number on github")
var updateFlagTitle = updateFlag.String("title", "", "title of issue")
var updateFlagBody = updateFlag.String("body", "", "body of issue")
var updateFlagState = updateFlag.String("state", "", "state of issue")

func updateIssue(number int, title, body, state, token string) {
	if number <= 0 {
		log.Fatal("invalid issue number")
	}
	// get repo
	repo := fetchRepo()
	
	editor := os.Getenv("EDITOR")

	var data map[string]string

	if editor != "" {
		result, err := cli.Retrieve(number, repo)
		if err != nil {
			log.Fatal("Error in retrieving: ", err)
		}

		fields := map[string]string{
			"title": result.Title,
			"body": result.Body,
			"state": result.State,
		}
		data = getFromEditor(editor, fields)
	} else {
		data = map[string]string{}
		if title != "" { data["title"] = title }
		if body != "" { data["body"] = body }
		if state != "" { data["state"] = state }
	}

	result, err := cli.Update(number, repo, data, token)
	if err != nil {
		log.Fatal("Error in updating: ", err)
	}
	fmt.Println("Issue updated successfully")
	displayIssue(result)
}