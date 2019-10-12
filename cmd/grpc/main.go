package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/brunoluiz/snapurl/internal/handler"
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
		cli.StringFlag{
			Name:   "gateway-enabled",
			Value:  "true",
			Usage:  "grpc gateway enabled",
			EnvVar: "GATEWAY_ENABLED",
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
		if err := handler.StartGRPCServer(grpcAddress); err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		if err := handler.StartGRPCGateway(gatewayAddress, grpcAddress); err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}
