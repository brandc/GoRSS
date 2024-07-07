package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"net/http"
)

// Command line variables
var RSSLink string
var OutputDir string

func init() {
	flag.StringVar(&RSSLink,   "r", "", 		   "RSS Link")
	flag.StringVar(&OutputDir, "o", "./GoRSSDefault/", "Output Directory")
	flag.Parse()
}

func URLDownload(URL string) (data []byte, err error) {
	Response, err := http.Get(URL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "")
		return nil, fmt.Errorf("%s -> http.Get/URLDownloader/", err)
	}
	defer Response.Body.Close()

	if Response.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Status code: %s\n", Response.Status)
		return nil, errors.New("Response did not contain '200 OK'")
	}

	data, err = io.ReadAll(Response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "")
		return nil, fmt.Errorf("%s -> io.ReadAll/URLDownloader/", err)
	}

	return data, nil
}

func main() {
	// Sanity checks of command line input.
	if len(RSSLink) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a link with '-r'.")
		return
	}

	// Display command line input.
	fmt.Println("\n===================================================")
	fmt.Printf("RSS feed: %s\n", RSSLink)
	fmt.Printf("Output directory: \"%s\"\n", OutputDir)
	fmt.Println("===================================================")

	data, err := URLDownload(RSSLink)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	if len(data) > 0 {
		fmt.Printf("File size: %d\n", len(data))
	}

	return
}










































