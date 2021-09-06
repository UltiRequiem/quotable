# Quotabler

[![GitMoji](https://img.shields.io/badge/Gitmoji-%F0%9F%8E%A8%20-FFDD67.svg)](https://gitmoji.dev)
[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/quotabler)](https://goreportcard.com/report/github.com/UltiRequiem/quotabler)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Lines Of Code](https://img.shields.io/tokei/lines/github.com/UltiRequiem/quotabler?color=blue&label=Total%20Lines)
![CodeQL](https://github.com/UltiRequiem/quotabler/workflows/CodeQL/badge.svg)

![Screenshot](./.github/assets/screenshot.png)

Get random quotes in terminal.

This project fetch the [Quotable.io API](https://api.quotable.io/random).

## Installation

If you just want to use the CLI app:

```bash
go install github.com/UltiRequiem/quotabler@latest
```

To be able to execute the command anywhere you need to have correctly setup your [Gopath](https://golang.org/doc/gopath_code).

Example: [Dotfiles](https://github.com/UltiRequiem/dotfiles/blob/53fece48cc95521e67a7a9277d6146aa14fe32f3/.zshrc#L32)

If you want to use it in code, install it in the normal way.

## Usage

CLI:

```bash
quotabler
```

Code:

```go
package main

import (
	"github.com/UltiRequiem/quotabler/pkg"
	"fmt"
)

func main() {
	quote, _ := quotabler.GetRandomQuoteAndAuthor()
	fmt.Println(quote)
}
```

This package exports 3 things:

- [`Quotable`](https://github.com/UltiRequiem/quotabler/blob/main/pkg/root.go#L7)
- [`GetQuotableObject`](https://github.com/UltiRequiem/quotabler/blob/main/pkg/root.go#L18)
- [`GetRandomQuoteAndAuthor`](https://github.com/UltiRequiem/quotabler/blob/main/pkg/root.go#L27)

### License

This project is Licensed under the [MIT](./LICENSE.md) License.

### Alternative

I also developed this in [Node.js](https://github.com/UltiRequiem/ranmess) and [Python](https://github.com/UltiRequiem/quoteran).
