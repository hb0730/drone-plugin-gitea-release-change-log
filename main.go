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
			Usage:  "gitea user token",
			EnvVar: "PLUGIN_GITEA_TOKEN,PLUGIN_TOKEN,GITEA_TOKEN,TOKEN",
		},
		cli.StringFlag{
			Name:   "changelog.config",
			Usage:  "the YAML config file for generate changelog",
			EnvVar: "PLUGIN_CHANGE_LOG_CONFIG,CHANGE_LOG_CONFIG,CONFIG",
		},
		cli.StringFlag{
			Name:   "changelog.sha1",
			Usage:  "The old git sha version. allow: tag name, commit id",
			EnvVar: "PLUGIN_CHANGE_LOG_SHA1,CHANGE_LOG_SHA1,SHA1",
			Value:  "prev",
		},
		cli.StringFlag{
			Name:   "changelog.sha2",
			Usage:  "The new git sha version. allow: tag name, commit id",
			EnvVar: "PLUGIN_CHANGE_LOG_SHA2,CHANGE_LOG_SHA2,SHA2",
			Value:  "last",
		},
		cli.BoolFlag{
			Name:   "changelog.verbose",
			Usage:  "show more information",
			EnvVar: "PLUGIN_CHANGE_LOG_VERBOSE,CHANGE_LOG_VERBOSE,VERBOSE",
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
	changeConfig := ChangeLogConfig{
		ConfigFile: ctx.String("changelog.config"),
		Sha1:       ctx.String("changelog.sha1"),
		Sha2:       ctx.String("changelog.sha2"),
		Verbose:    ctx.Bool("changelog.verbose"),
	}
	return plugin.Exec(ctx, changeConfig)
}
