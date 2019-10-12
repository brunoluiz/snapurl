package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/brunoluiz/snapurl/server"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	a := cli.NewApp()

	a.Name = "snapurl server"
	a.Usage = "URLSnap server API"
	a.Description = "URLSnap server API"
	a.Action = start

	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "port",
			Value:  "5000",
			Usage:  "grpc server port",
			EnvVar: "PORT",
		},
		cli.BoolFlag{
			Name:   "gateway-disable",
			Usage:  "disable grpc gateway",
			EnvVar: "GATEWAY_DISABLE",
		},
		cli.StringFlag{
			Name:   "gateway-port",
			Value:  "4000",
			Usage:  "grpc gateway port",
			EnvVar: "GATEWAY_PORT",
		},
	}

	if err := a.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func start(c *cli.Context) error {
	grpcAddress := fmt.Sprintf(":%s", c.String("port"))
	gatewayAddress := fmt.Sprintf(":%s", c.String("gateway-port"))

	go func() {
		if err := server.StartGRPCServer(grpcAddress); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if c.Bool("gateway-disable") {
			return
		}

		if err := server.StartGRPCGateway(gatewayAddress, grpcAddress); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}
