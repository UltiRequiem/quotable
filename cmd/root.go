package cmd

import (
	"flag"
	"fmt"
	"strconv"
	"sync"

	chigo "github.com/UltiRequiem/chigo/pkg"
	quotable "github.com/UltiRequiem/quotable/pkg"
)

func Execute() {
	flag.Parse()

	times := flag.Arg(0)

	if times == "" {
		times = "1"
	}

	parsedTimes, errorParsingNumber := strconv.Atoi(times)

	if errorParsingNumber != nil {
		fmt.Println("The input was not a valid number.")
		panic(errorParsingNumber)
	}

	var wg sync.WaitGroup

	wg.Add(parsedTimes)

	for i := 0; i < parsedTimes; i++ {
		go func() {
			data, error := quotable.GetQuotable()

			if error != nil {
				panic(error)
			}

			message := fmt.Sprintf("%s\n -%s", data.Content, data.Author)

			fmt.Println(chigo.Colorize(message))

			wg.Done()
		}()
	}

	wg.Wait()
}
