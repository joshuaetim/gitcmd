package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"fmt"
)

func saveRepo(repo string) {
	f, err := os.Create("repo.txt")
	if err != nil {
		log.Fatal("Error creating file: ", err)
	}
	_, err = f.WriteString(repo)
	if err != nil {
		log.Fatal("Can't save repo: ", err)
	}
}

func fetchRepo() string {
	var in io.Reader
	f, err := os.Open("repo.txt")
	in = f
	saved := true
	if err != nil {
		in = os.Stdin
		saved = false
		fmt.Println("Please enter the repo:")
	}
	input := bufio.NewScanner(in)
	input.Scan()
	repo := input.Text()

	if !saved {
		saveRepo(repo)
	}

	if saved {
		fmt.Printf("Using %q...\nEnter new repository to override\n(empty for no changes): ", repo)
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		response := input.Text()
		if response != "" {
			repo = response
			saveRepo(repo)
		}
	}

	return repo
}