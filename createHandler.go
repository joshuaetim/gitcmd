package main

import (
	"flag"
	"fmt"
	// "gopl/ch4/gitcmd/cli"
	"log"
	"os"
)

var createFlag = flag.NewFlagSet("create", flag.ExitOnError)
var createFlagTitle = createFlag.String("title", "", "title of issue")
var createFlagBody = createFlag.String("body", "", "body of issue")

func createIssue(title, body, token string) {
	editor := os.Getenv("EDITOR")

	// get repo
	repo := fetchRepo()

	var data map[string]string

	if editor != "" {
		fields := map[string]string{
			"title": "",
			"body": "",
		}
		data = getFromEditor(editor, fields)
	} else {
		data = map[string]string{
			"title": title,
		}
		if body != "" { data["body"] = body }
	}

	if data["title"] == "" {
		log.Fatal("error: title field is required")
	}
	
	result, err := Create(repo, data, token)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Issue Created successfully")
	displayIssue(result)
}