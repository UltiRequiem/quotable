package cmd

import (
	"github.com/UltiRequiem/quotable/pkg"
	"github.com/fatih/color"
)

// Print a random quote and his/her Author
func Execute() {
	quote, author := quotable.GetRandomQuoteAndAuthor()
	color.Cyan(" " + quote)
	color.Green(" - " + author)
}
