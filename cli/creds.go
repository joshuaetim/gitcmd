package cli

import (
	"bufio"
	"fmt"
	"log"

	// "log"
	"os"
)

func CheckCreds() (token string) {
	f, err := os.Open("creds.txt")
	in := f
	saved := true
	if err != nil {
		in = os.Stdin
		saved = false
		fmt.Println("Please enter your access token: ")
	}
	defer f.Close()

	input := bufio.NewScanner(in)
	// get line 1; token
	input.Scan()
	token = input.Text()

	if !saved {
		f, err = os.Create("creds.txt")
		if err != nil {
			log.Fatal("Can't create file: ", err)
		}
		_, err = f.WriteString(token)
		if err != nil {
			log.Fatal("Can't save token: ", err)
		}
	}

	return token
}