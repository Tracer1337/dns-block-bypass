// Command remote is a chromedp example demonstrating how to connect to an
// existing Chrome DevTools instance using a remote WebSocket URL.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

const chromeURL = "wss://chrome.browserless.io"

func main() {
	token := flag.String("token", "", "Your Browserless API Token")
	url := flag.String("url", "", "The URL you want to screenshot")
	out := flag.String("out", "screenshot.png", "The output file")
	flag.Parse()

	if *token == "" {
		log.Fatal("Must specify -token")
	}

	if *url == "" {
		log.Fatal("Must specify -url")
	}

	remoteURL := fmt.Sprintf("%s?token=%s", chromeURL, *token)
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), remoteURL)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	var buf []byte
	if err := chromedp.Run(ctx,
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Navigate(*url),
		chromedp.FullScreenshot(&buf, 100),
	); err != nil {
		log.Fatalf("Failed to capture a screenshot of %s: %v", *url, err)
	}

	if err := os.WriteFile(*out, buf, 0777); err != nil {
		log.Fatal(err)
	}

	log.Printf("Wrote to %s", *out)
}
