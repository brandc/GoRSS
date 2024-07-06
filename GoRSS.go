package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
//	"log/slog"
//	"time"

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
		return nil, errors.Join(err, " -> http.Get/RSSFeedDownloader/")
	}
	defer Response.Body.Close()

	if Response.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "Status code: %s\n", Response.Status)
		return nil, errors.Join("Response did not contain '200 OK'")
	}

	data = io.ReadAll(Response.Body)
	return data, nil
}

func main() {
	// Sanity checks of command line input.
	if len(RSSLink) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a link with '-r'.")
		return
	}

	// Display command line input.
	fmt.Printf("RSS feed: %s\n", RSSLink)
	fmt.Printf("Output directory: \"%s\"\n", OutputDir)

	return
}


