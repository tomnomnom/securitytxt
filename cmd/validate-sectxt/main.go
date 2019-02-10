package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/tomnomnom/securitytxt"
)

func main() {
	flag.Parse()

	url := flag.Arg(0)
	if url == "" {
		fmt.Println("usage: validate-sectxt <url>")
		os.Exit(1)
	}

	r, err := getURL(url)
	if err != nil {
		fmt.Printf("failed to fetch url: %s\n", err)
		os.Exit(2)
	}

	f, err := securitytxt.FromReader(r)
	if err != nil {
		fmt.Printf("encountered %d errors during parsing\n\n", len(f.Errors()))
	}

	fmt.Println("errors:")
	for _, e := range f.Errors() {
		fmt.Printf("\t%s\n", e)
	}
	fmt.Println("")

	fmt.Println("comments:")
	for _, c := range f.Comments() {
		fmt.Printf("\t%s\n", c)
	}
	fmt.Println("")

	printFields(f, "contact")

	printFields(f, "encryption")

	printFields(f, "acknowledgments")

	if len(f.Errors()) > 0 {
		os.Exit(3)
	}

	os.Exit(0)
}

func printFields(file *securitytxt.File, option string) {
	fmt.Printf("%s:\n", option)
	for _, field := range file.Fields(option) {
		if len(field.Comments()) > 0 {
			fmt.Println("")
		}
		for _, com := range field.Comments() {
			fmt.Printf("\t%s\n", com)
		}
		fmt.Printf("\t%s\n", field)
	}
	fmt.Println("")
}

func getURL(url string) (io.Reader, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return bufio.NewReader(resp.Body), err
}
