package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
	"log"
	"os"
)

var version = "unknown"

func main() {
	app := cli.NewApp()
	app.Name = "drone-plugin-gitea-release-change-log"
	app.Description = "update/create gitea release change log drone plugin"
	app.Copyright = "© 2021-now hb0730"
	app.Action = run
	app.Version = version
	app.Authors = []cli.Author{
		{
			Name:  "hb0730",
			Email: "huangbing0730@gmail.com",
		},
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "debug",
			EnvVar: "DEBUG",
		},
		cli.StringFlag{
			Name:   "drone.tag",
			Usage:  "drone tag",
			EnvVar: "DRONE_TAG",
		},
	}
	if _, err := os.Stat("/run/drone/env"); err == nil {
		godotenv.Overload("/run/drone/env")
	}
	if err := app.Run(os.Args); nil != err {
		log.Fatal(err)
	}
}
func run(ctx *cli.Context) error {
	plugin := &Plugin{
		Debug: ctx.Bool("debug"),
		Drone: Drone{
			ctx.String("drone.tag"),
		},
	}
	return plugin.Exec(ctx)
}
