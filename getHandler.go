package main

import (
	"flag"
	// "gopl/ch4/gitcmd/cli"
	"log"
)

var getFlag = flag.NewFlagSet("get", flag.ExitOnError)
var getFlagNumber = getFlag.Int("n", 0, "the issue number on github")

func getIssue(number int) {
	if number <= 0 {
		log.Fatal("Invalid issue number")
	}

	// get repo
	repo := fetchRepo()
	// end repo fetch

	result, err := Retrieve(number, repo)
	if err != nil {
		log.Fatal("Error in retrieving: ", err)
	}
	displayIssue(result)
}

