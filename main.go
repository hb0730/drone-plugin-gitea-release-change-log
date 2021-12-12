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
			EnvVar: "PLUGIN_DEBUG,DEBUG",
		},
		cli.StringFlag{
			Name:   "drone.tag",
			Usage:  "drone tag",
			EnvVar: "DRONE_TAG",
		},
		cli.StringFlag{
			Name:   "drone.repo",
			Usage:  "drone repo",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "drone.repo.name",
			Usage:  "drone repo name",
			EnvVar: "DRONE_REPO_NAME",
		},
		cli.StringFlag{
			Name:   "drone.repo.owner",
			Usage:  "drone repo owner",
			EnvVar: "DRONE_REPO_OWNER",
		},
		cli.StringFlag{
			Name:   "drone.commit",
			Usage:  "drone commit hash",
			EnvVar: "DRONE_COMMIT",
		},
		cli.StringFlag{
			Name:   "gitea.url",
			Usage:  "gite base url",
			EnvVar: "PLUGIN_GITEA_URL,PLUGIN_URL,GITEA_URL,URL",
		},
		cli.StringFlag{
			Name:   "gitea.token",
			Usage:  "git user token",
			EnvVar: "PLUGIN_GITEA_TOKEN,PLUGIN_TOKEN,GITEA_TOKEN,TOKEN",
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
			Tag:        ctx.String("drone.tag"),
			Repo:       ctx.String("drone.repo"),
			RepoName:   ctx.String("drone.repo.name"),
			Owner:      ctx.String("drone.repo.owner"),
			CommitHash: ctx.String("drone.commit"),
		},
		Gitea: Gitea{
			URL:   ctx.String("gitea.url"),
			Token: ctx.String("gitea.token"),
		},
	}
	return plugin.Exec(ctx)
}
