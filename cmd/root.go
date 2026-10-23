package cmd

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"sync"

	chigo "github.com/UltiRequiem/chigo/pkg"
	quotable "github.com/UltiRequiem/quotable/pkg"
)

func Execute() {
	apiURL := flag.String(
		"api-url",
		"",
		fmt.Sprintf("base URL of a Quotable server (default $%s, then %s)", quotable.APIURLEnv, quotable.DefaultAPIURL),
	)

	flag.Parse()

	times := flag.Arg(0)

	if times == "" {
		times = "1"
	}

	parsedTimes, errorParsingNumber := strconv.Atoi(times)

	if errorParsingNumber != nil || parsedTimes < 1 {
		fmt.Fprintln(os.Stderr, "The input was not a valid number.")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	var failed sync.Once

	wg.Add(parsedTimes)

	for i := 0; i < parsedTimes; i++ {
		go func() {
			defer wg.Done()

			data, err := quotable.GetQuotableFrom(*apiURL)

			if err != nil {
				failed.Do(func() {
					fmt.Fprintf(os.Stderr, "Could not get a quote from %s.\n", quotable.BaseURL(*apiURL))
					fmt.Fprintf(os.Stderr, "Pass -api-url or set %s to point at a running Quotable server.\n", quotable.APIURLEnv)
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				})

				return
			}

			message := fmt.Sprintf("%s\n -%s", data.Content, data.Author)

			fmt.Println(chigo.Colorize(message))
		}()
	}

	wg.Wait()
}
