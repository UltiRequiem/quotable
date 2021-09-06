package quotabler

import (
	"io/ioutil"
	"log"
	"net/http"
)

func fetch(url string) *http.Response {
	response, getError := http.Get(url)

	if getError != nil {
		log.Fatal("There was an error while doing the fetch, you may not have the internet.")
	}

	return response
}

func parseHttpResponse(response *http.Response) string {
	responseData, readError := ioutil.ReadAll(response.Body)

	if readError != nil {
		log.Fatal("Error reading response.")
	}

	return string(responseData)
}
