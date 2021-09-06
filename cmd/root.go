package cmd

import (
	"github.com/UltiRequiem/quotabler/pkg"
	"github.com/fatih/color"
)

// Print a random quote and his/her Author
func Execute() {
	quote, author := quotabler.GetRandomQuoteAndAuthor()
	color.Cyan(" " + quote)
	color.Green(" - " + author)
}
