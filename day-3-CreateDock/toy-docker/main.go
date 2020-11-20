package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const usage = `toy-docker is a simple container runtime implementation.
				The purpose of this project is to learn how docker works and how
				to write a docker by ourselves 
							Enjoy it, just for fun.`

func main() {
	app := cli.NewApp()
	app.Name = "toy-docker"
	app.Usage = usage

	app.Commands = []cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		// Log as Json instead of the default Ascii formatter
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		return nil
	}

	if err:=app.Run(os.Args) ; err!= nil{

	}
}
