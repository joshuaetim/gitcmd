package main

import (
	"log"
	"os"
	"encoding/json"
	"os/exec"
	
)


func getFromEditor(editor string, fields interface{}) map[string]string {
	tmpfile := "tmp"

	f, err := os.Create(tmpfile)
	if err != nil {
		log.Fatal("Create file error: ", err)
	}
	defer f.Close()
	defer os.Remove(tmpfile)

	jsonObject, err := json.MarshalIndent(fields, "", "    ")
	if err != nil {
		log.Fatal("Marshal error: ", err)
	}

	_, err = f.Write(jsonObject)
	if err != nil {
		log.Fatal("Write file error: ", err)
	}
	
	cmd := exec.Command(editor, tmpfile, "-w", "-n")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal("Error in running command: ", err)
	}

	contents, err := os.ReadFile(tmpfile)
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	// convert json to map
	var data map[string]string
	if err := json.Unmarshal(contents, &data); err != nil {
		log.Fatal("Unmarshal error: ", err)
	}

	return data
}