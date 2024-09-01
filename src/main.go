package main

import (
	"flag"
	"fmt"
)

func main() {
	var url string
	var detect, exploit bool

	flag.StringVar(&url, "u", "", "URL to test")
	flag.BoolVar(&detect, "v", false, "Detect SQL injection")
	flag.BoolVar(&exploit, "e", false, "Exploit SQL injection")
	flag.Parse()

	if url == "" {
		fmt.Println("Usage: go run main.go -u <URL> [-v] [-e]")
		return
	}

	if detect {
		fmt.Printf("Scanning %s for SQL injection...\n", url)
		if ScanURLForSQLInjection(url) {
			fmt.Println("Possible SQL injection vulnerability found!")
		} else {
			fmt.Println("No SQL injection vulnerabilities detected.")
		}
	}

	if exploit {
		fmt.Printf("Trying to exploit %s...\n", url)
		if ExploitSQLInjection(url) {
			fmt.Println("SQL injection vulnerability exploited successfully!")
		} else {
			fmt.Println("Failed to exploit the SQL injection vulnerability.")
		}
	}
}
