# Quotable

[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/quotable)](https://goreportcard.com/report/github.com/UltiRequiem/dadjokes)
[![Go Reference](https://pkg.go.dev/badge/github.com/UltiRequiem/quotable/pkg.svg)](https://pkg.go.dev/github.com/UltiRequiem/dadjokes/pkg)

![Screenshot](./.github/assets/screenshot.png)

A wrapper around the [Quotable.io API](https://api.quotable.io).

## Usage

```go
import (
	"fmt"
	"github.com/UltiRequiem/quotable/pkg"
)


func main(){
        data, error := quotable.GetQuotable()

        fmt.Printf("%s\n -%s\n", data.Content, data.Author)
}
```

## Docs

Is hosted on
[pkg.go.dev](https://pkg.go.dev/github.com/UltiRequiem/quotable/pkg) 📖

## CLI

### Installation

```bash
go install github.com/UltiRequiem/quotable@latest
```

### Usage

```bash
quotable
```

Or an `X` quantity of times :👇

```bash
quotable 38
```

[Video Showcase](https://youtu.be/rfBqpPCBrDE) 📷

## Other Art

I also made this on [JavaScript](https://github.com/UltiRequiem/ranmess),
[Python](https://github.com/UltiRequiem/quoteran),
[a website](https://github.com/UltiRequiem/ulti-random-quotes), and
[Rust](https://github.com/UltiRequiem/ruquotes).

## Licence

Licensed under the MIT License 📄
