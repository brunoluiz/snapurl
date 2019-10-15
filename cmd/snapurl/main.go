package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/brunoluiz/snapurl"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	a := cli.NewApp()

	a.Name = "snapurl command line tool"
	a.Usage = "easy snapshots for websites"
	a.ArgsUsage = "[url Website URL]"
	a.Action = start

	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "out-dir",
			Usage: "output directory",
			Value: ".",
		},
		cli.StringFlag{
			Name:  "out",
			Usage: "output path (folder + filename)",
		},
		cli.Int64Flag{
			Name:  "wait",
			Usage: "wait period in seconds to render the page",
			Value: 5,
		},
		cli.Int64Flag{
			Name:  "width",
			Usage: "viewport width",
			Value: 1440,
		},
		cli.Int64Flag{
			Name:  "height",
			Usage: "viewport width",
			Value: 900,
		},
		cli.Float64Flag{
			Name:  "scale",
			Usage: "viewport scale",
			Value: 1,
		},
	}

	if err := a.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context) error {
	url := c.Args().Get(0)
	if url == "" {
		return errors.New("no URL was set as argument")
	}

	buf, err := snapurl.Snap(context.Background(), url, snapurl.Params{
		WaitPeriod: time.Duration(c.Int("wait")) * time.Second,
		Width:      c.Int64("width"),
		Height:     c.Int64("height"),
		Scale:      c.Float64("scale"),
	})
	if err != nil {
		return nil
	}

	err = os.MkdirAll(c.String("out-dir"), os.ModePerm)
	if err != nil {
		return errors.New("output directory couldn't be created")
	}

	var path string
	if c.String("out") != "" {
		path = c.String("out")
	} else {
		file := fmt.Sprintf("screenshot-%s.png", time.Now().Format("2006-01-02T15:04:05Z"))
		path = fmt.Sprintf("%s/%s", c.String("out-dir"), file)
	}

	return ioutil.WriteFile(path, buf, 0644)
}
