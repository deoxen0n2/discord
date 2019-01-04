package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	dc "github.com/bwmarrin/discordgo"
	resty "gopkg.in/resty.v1"
)

func main() {
	// Parse flags.
	webhookURL := flag.String("webhook-url", "", "a webhook URL")
	flag.Parse()

	if *webhookURL == "" {
		printUsage()
		os.Exit(1)
	}

	// Prepare message.
	reader := bufio.NewReader(os.Stdin)
	msgText, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	msg := dc.Message{
		Content: msgText,
	}

	// Send message.
	_, err = resty.R().SetBody(msg).Post(*webhookURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`
Usage: discord STDIN --webhook-url=WEBHOOK_URL

  Send a message to Discord channel.

Options:
  --webhook-url A Discord Webhook URL of the channel you want to send message to`)
}
