package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "getouts"
	app.Usage = "Get the conversation history for a Hangouts conversation."
	app.Version = "0.0.1"
	app.Author = "vaneck"
	app.Email = "https://github.com/vaneck/getouts"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "message_count",
			Value: 10,
			Usage: "Number of messages to pull.",
		},
		cli.StringFlag{
			Name:  "conversation_id",
			Value: "",
			Usage: "Full hangouts conversation id to fetch.",
		},
	}
	app.Action = func(c *cli.Context) error {
		msgcount := c.Int("message_count")
		convid := c.String("conversation_id")
		if convid == "" {
			cli.ShowAppHelp(c)
			return cli.NewExitError("Conversation ID should be specified.", 11)
		}
		Getouts(convid, msgcount)
		return nil
	}

	app.Run(os.Args)

}
