package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func upload(ip string, pokemon []uint8) {
	client := &http.Client{}
	// Build the request, but don't execute the post just yet
	req, err := http.NewRequest("POST", "http://"+ip+":9000", bytes.NewBuffer(pokemon))
	if err != nil {
		fmt.Println(err.Error())
	}
	// Do the post now.
	_, err = client.Do(req)
	if err != nil && !strings.Contains(err.Error(), "EOF") {
		fmt.Println(err.Error())
	}
	return
}
