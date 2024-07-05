package main

import (
	"flag"
	"fmt"
	"os"
//	"log"
//	"time"

//	"net/http"
)

// Command line variables
var URLLink string
var OutputDir string

func init() {
	flag.StringVar(&URLLink, "i", "", "Input URL")
	flag.StringVar(&OutputDir, "o", "./GoRSSDefault/", "Output Directory")
	flag.Parse()
}

func main() {
	// Sanity checks of command line input.
	if len(URLLink) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify a link with '-i'.")
		return
	}

	// Display command line input.
	fmt.Printf("URLLink: %s\n", URLLink)
	fmt.Printf("OutputDir: %s\n", OutputDir)

	return
}


