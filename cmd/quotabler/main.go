package main

import (
	"github.com/UltiRequiem/quotabler/pkg"
	"github.com/fatih/color"
)

func main() {
	var parsedData quotabler.Quotable = quotabler.FetchQuotabler()
	color.Cyan(" " + parsedData.Content)
	color.Red(" - " + parsedData.Author)
}
