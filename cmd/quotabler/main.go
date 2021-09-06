package main

import (
	"fmt"
	"github.com/UltiRequiem/quotabler/pkg"
)

func main() {
	parsedData := quotabler.FetchQuotabler()
	fmt.Println(" " + parsedData.Content)
	fmt.Println(" - " + parsedData.Author)
}
