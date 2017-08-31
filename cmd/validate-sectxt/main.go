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
		fmt.Printf("Encountered %d errors during parsing\n\n", len(f.Errors()))
	}

	fmt.Println("Errors:")
	for _, e := range f.Errors() {
		fmt.Printf("\t%s\n", e)
	}
	fmt.Println("")

	fmt.Println("Comments:")
	for _, c := range f.Comments() {
		fmt.Printf("\t%s\n", c)
	}
	fmt.Println("")

	fmt.Println("Contact:")
	for _, c := range f.Contact() {
		fmt.Printf("\t%s\n", c)
	}
	fmt.Println("")

	fmt.Println("Encryption:")
	for _, e := range f.Encryption() {
		fmt.Printf("\t%s\n", e)
	}
	fmt.Println("")

	fmt.Println("Disclosure:")
	for _, d := range f.Disclosure() {
		fmt.Printf("\t%s\n", d)
	}
	fmt.Println("")

	fmt.Println("Acknowledgement:")
	for _, a := range f.Acknowledgement() {
		fmt.Printf("\t%s\n", a)
	}

	if len(f.Errors()) > 0 {
		os.Exit(3)
	}

	os.Exit(0)
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
