package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/hound/dd"
	"github.com/zorkian/go-datadog-api"
	"log"
	"os"
	"time"
)

func main() {

	dd.Init()

	dd := cli.NewApp()
	dd.Name = "hound"
	dd.Usage = "Command line for getting and creating Data Dog Events"
	dd.Email = "drewvanstone@gmail.com"
	dd.Author = "Drew Flower"
	dd.Commands = []cli.Command{
		{
			Name:  "create-event",
			Usage: "Create Data Dog Event",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "title",
					Usage: "The event title",
				},
				cli.StringFlag{
					Name:  "text",
					Usage: "The event text",
				},
				cli.BoolFlag{
					Name:  "time",
					Usage: "Boolean value to pass current time to event",
				},
				cli.StringFlag{
					Name:  "priority",
					Usage: "Set the priority. Valid options ['low', 'normal']",
				},
				cli.StringFlag{
					Name:  "alert-type",
					Usage: "Set event alert type. Valid options ['error', 'warning', 'info', 'success']",
				},
				cli.BoolFlag{
					Name:  "hostname",
					Usage: "Boolean value to pass hostname to event",
				},
			},
			Action: func(c *cli.Context) {
				createEvent(c)
			},
		},
	}
	dd.Run(os.Args)
}

func createEvent(c *cli.Context) {

	event := new(datadog.Event)

	// Set Event Title.
	// Support passing it in by flag or as an argument.
	if c.String("title") != "" {
		event.Title = c.String("title")
	} else if len(c.Args()) != 0 {
		event.Title = c.Args().First()
	} else {
		cli.ShowSubcommandHelp(c)
		os.Exit(1)
	}

	// Set Event Text
	if c.String("text") != "" {
		event.Text = c.String("text")
	}

	// Set Event Time
	if c.Bool("time") {
		event.Time = int(time.Now().Unix())
	}

	// Set Event Hostname
	if c.Bool("hostname") {
		host, err := os.Hostname()
		if err != nil {
			log.Fatalf("[hound] ERROR: %s\n", err)
		}
		event.Host = host
	}

	// Set Event Priority
	if c.String("priority") != "" {
		switch c.String("priority") {
		case "low", "normal":
			event.Priority = c.String("priority")
		default:
			log.Fatal("[hound] ERROR: Priority must be 'normal' or 'low'")
		}
	}

	// Set Event AlertType
	if c.String("alert-type") != "" {
		switch c.String("alert-type") {
		case "error", "warning", "info", "success":
			event.AlertType = c.String("alert-type")
		default:
			log.Fatal("[hound] ERROR: AlertType must be 'error', 'warning', 'info', 'success'")
		}
	}

	// Submit event
	returnedEvent, err := dd.Client.PostEvent(event)
	if err != nil {
		log.Fatalf("[hound] ERROR: %s\n", err)
	}
	printJson(returnedEvent)
}

func printJson(input ...interface{}) {
	b, _ := json.MarshalIndent(input, "", " ")
	fmt.Println(string(b))
}
