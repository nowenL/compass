package main

import (
	"github.com/urfave/cli"
	"os"
)

const (
	usage = `
Compass(github.com/weiwei04/compass) is a front end for tiller, it implements all tiller grpc api and integrate with Helm-Registry(github.com/caicloud/helm-registry) and provide a extra grpc api: InstallCompassRelease, it will download chart from helm-registry and send a InstallReleaseRequest to tiller

Avaiable Commands:
  compass install [CHART] [OPTIONS]
  `
)

func main() {
	app := cli.NewApp()
	app.Name = "compass"
	app.Usage = usage
	app.Commands = []cli.Command{
		installReleaseCommand,
	}
	app.Run(os.Args)
}