package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}

// Message msg
type Message struct {
    Name, Text string
}
func mainT() {
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("err: %v", err)
		}
		// fmt.Printf("%s: %s \n", m.Name, m.Text)
		fmt.Println(json.NewEncoder(os.Stdout).Encode(&m))
	}
}