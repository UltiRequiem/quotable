package main

import (
	"github.com/UltiRequiem/quotabler/pkg"
	"github.com/fatih/color"
)

func main() {
	quote, author := quotabler.GetRandomQuoteAndAuthor()
	color.Cyan(" " + quote)
	color.Green(" - " + author)
}
