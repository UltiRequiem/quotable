package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Quotable struct {
	Content string `json:"content"`
	Author  string `json:"author"`
}

func fetch(url string) *http.Response {
	response, getError := http.Get(url)

	if getError != nil {
		fmt.Println("There was an error while doing the fetch, you may not have the internet.")
		os.Exit(1)
	}

	return response
}

func parseHttpResponse(response *http.Response) string {
	responseData, readError := ioutil.ReadAll(response.Body)

	if readError != nil {
		fmt.Println("Error reading response.")
		os.Exit(1)
	}

	return string(responseData)
}

func main() {
	var data string = parseHttpResponse(fetch("https://api.quotable.io/random"))

	var parsedData Quotable
	json.Unmarshal([]byte(data), &parsedData)

	fmt.Println(" " + parsedData.Content)
	fmt.Println(" - " + parsedData.Author)
}
