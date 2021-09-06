package main

import (
	"encoding/json"
	"fmt"
	"github.com/UltiRequiem/quotabler/pkg"
)

func main() {
	var data string = quotabler.ParseHttpResponse(quotabler.Fetch("https://api.quotable.io/random"))

	var parsedData quotabler.Quotable
	json.Unmarshal([]byte(data), &parsedData)

	fmt.Println(" " + parsedData.Content)
	fmt.Println(" - " + parsedData.Author)
}
