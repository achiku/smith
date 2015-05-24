package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Message struct {
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
type Messages []Message

func main() {
	client := http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	req, err := http.NewRequest("GET", "http://localhost:8080/message/", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		fmt.Println(k, v)
	}

	bodyJson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var messages Messages
	err = json.Unmarshal([]byte(bodyJson), &messages)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(messages)
	for _, t := range messages {
		fmt.Println(t)
		fmt.Println(t.Name)
		fmt.Println(t.Completed)
	}
}
