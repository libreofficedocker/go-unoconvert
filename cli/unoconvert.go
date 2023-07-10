package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/libreofficedocker/go-unoconvert/unoconvert"
	"github.com/urfave/cli"
)

var Version = "0.0.0"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Usage = "A Go wrapper implementation for unoconvert"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: "127.0.0.1:2002",
			Usage: "The addr used by the unoserver api server",
		},
		cli.StringFlag{
			Name:  "infile",
			Usage: "The input file to convert",
		},
		cli.StringFlag{
			Name:  "outfile",
			Usage: "The output file to convert",
		},
		cli.StringFlag{
			Name:  "executable",
			Value: "unoconvert",
			Usage: "The executable to use for unoconvert",
		},
	}
	app.Authors = []cli.Author{
		{
			Name:  "libreofficedocker",
			Email: "https://github.com/libreofficedocker",
		},
	}
	app.Action = action

	// Set log prefix
	log.SetPrefix(fmt.Sprintf("[%s]: ", app.Name))

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := unoconvert.Default()

	infile := c.String("infile")
	outfile := c.String("outfile")

	cmd := client.CommandContext(ctx, infile, outfile)
	log.Printf("Command: %s %s", client.Executable, cmd.Args)

	return cmd.Run()
}
