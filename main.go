package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := flag.String("u", "", "url")
	list := flag.String("l", "", "list")
	responseCode := flag.Int("rc", 200, "response code")
	flag.Parse()

	if *url != "" {
		processURL(*url, *responseCode)
	} else if *list != "" {
		file, err := os.Open(*list)
		if err != nil {
			fmt.Printf("open file error: %v\n", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			processURL(scanner.Text(), *responseCode)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("read file error: %v\n", err)
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			processURL(scanner.Text(), *responseCode)
		}

		if err := scanner.Err(); err != nil {
			fmt.Printf("input read error: %v\n", err)
		}
	}
}

func processURL(url string, expectedCode int) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == expectedCode {
		fmt.Printf("%s\n", url)
}
}
