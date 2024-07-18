package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"time"

//	"RSSParse"

	"net/http"
)

var RSSTimeFormat = "Mon, 02 Jan 2006 15:04:05 -0700"

// Command line variables
var RSSLinks []string
var OutputDir string

func init() {
	//flag.StringVar(&RSSLink,   "r", "", 		   "RSS Link")
	flag.StringVar(&OutputDir, "o", "./GoRSSDefault/", "Output Directory")
	flag.Parse()
	RSSLinks = flag.Args()
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
		return nil, errors.New("response did not contain '200 OK'")
	}

	data, err = io.ReadAll(Response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "")
		return nil, fmt.Errorf("%s -> io.ReadAll/URLDownloader/", err)
	}

	return data, nil
}

func dumpRSS(feedLink, OutputDir string, ItemsToDisplay int) {
	data, err := URLDownload(feedLink)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	} else if len(data) < 1 {
		fmt.Printf("Zero Length Response.\n")
		return
	}

	// Display command line input.
	fmt.Println("\n===================================================")
	fmt.Printf("RSS feed: %s\n", feedLink)
	fmt.Printf("Output directory: \"%s\"\n", OutputDir)
	fmt.Printf("File size: %d\n", len(data))	
	fmt.Println("===================================================")

	feed, err := RSSFeedParse(string(data[:]))
	if err != nil {
		fmt.Printf("main -> %s\n", err)
	}

	if len(feed.Channels) < 1 {
		fmt.Printf("No Channels\n")
		return
	}

	for _, ChannelInstance := range(feed.Channels) {
		fmt.Printf("Channel: %s\n", ChannelInstance.Title)
		if len(ChannelInstance.Links) > 0 {
			for _, Link := range(ChannelInstance.Links) {
				if len(Link.URL) > 0 {
					fmt.Printf("\tLink: %s\n", Link.URL)
				} else if len(Link.Href) > 0 {
					fmt.Printf("\tLink: %s\n", Link.Href)
				}
			}
		}
		if len(ChannelInstance.Images) > 0 {
			for _, ImageInstance := range(ChannelInstance.Images) {
				if len(ImageInstance.Href) > 0 {
					fmt.Printf("\tImage URL: %s\n", ImageInstance.Href)
				} else {
					continue
				}
				if ImageInstance.Height > 0 {
					fmt.Printf("\t\tHeight: %d\n", ImageInstance.Height)
				}
				if ImageInstance.Width > 0 {
					fmt.Printf("\t\tWidth: %d\n", ImageInstance.Width)
				}
			}
		}
		fmt.Printf("\tDescription: %s\n", ChannelInstance.Description)

		fmt.Println()

		if len(ChannelInstance.Items) < 1 {
			continue
		}

		for Index, ItemInstance := range(ChannelInstance.Items) {
			if Index >= ItemsToDisplay {
				break
			}

			fmt.Printf("--------------------------------------------------------------\n")
			fmt.Printf("Title: %s\n", ItemInstance.Title)
			fmt.Printf("Link: %s\n", ItemInstance.Link)
			//fmt.Printf("Description: %s\n", ItemInstance.Description)
			fmt.Printf("Publication Date: %s\n", ItemInstance.PublicationDate)

			PublicationTime, err := time.Parse(RSSTimeFormat, ItemInstance.PublicationDate)
			if err != nil {
				fmt.Printf("Failed to parse time string with error: %s\n", err)
			}
			fmt.Printf("Unix: %d\n", PublicationTime.Unix())
			for _, enclosureInstance := range(ItemInstance.Enclosures) {
				fmt.Printf("URL: %s\n", enclosureInstance.URL)
				fmt.Printf("Length: %d\n", enclosureInstance.Length)
			}
			fmt.Printf("--------------------------------------------------------------\n\n")
		}
	}
}


func main() {
	// Sanity checks of command line input.
	if len(RSSLinks) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a link with '-r'.")
		return
	}

	for _, link := range(RSSLinks) {
		dumpRSS(link, OutputDir, 1)
	}
}










































