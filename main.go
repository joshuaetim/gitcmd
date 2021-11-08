package main

import (
	"fmt"
	// "gopl/ch4/gitcmd/cli"
	"log"
	"os"
)

func main() {
	token := CheckCreds()

	if len(os.Args) < 2 {
		log.Fatal("You must use a subcommand")
	}

	switch os.Args[1]{
	case "create":
		createFlag.Parse(os.Args[2:])
		createIssue(*createFlagTitle, *createFlagBody, token)
	case "get":
		getFlag.Parse(os.Args[2:])
		getIssue(*getFlagNumber)
	case "update":
		updateFlag.Parse(os.Args[2:])
		updateIssue(*updateFlagNumber, *updateFlagTitle, *updateFlagBody, *updateFlagState, token)
	}
}

func displayIssue(result *Issue) {
	fmt.Printf("Number:\t%d\nURL:\t%s\nTitle:\t%s\nBody:\t%s\nUser:\t%s\nState:\t%s\nCreated At:\t%s\nUpdated At:\t%s\n", result.Number, result.HTMLURL, result.Title, result.Body, result.User.Login, result.State, result.CreatedAt, result.UpdatedAt)
}