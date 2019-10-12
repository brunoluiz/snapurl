package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	a := cli.NewApp()

	a.Name = "urlsnap server"
	a.Usage = "URLSnap server API"
	a.Description = "URLSnap server API"
	a.Action = start

	a.Flags = []cli.Flag{}

	if err := a.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context) error {
	log.Info("hello")
	return nil
}
