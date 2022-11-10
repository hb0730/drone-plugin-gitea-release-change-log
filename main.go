package main

import (
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var version = "unknown"

func main() {
	app := &cli.App{
		Name:      "drone-plugin-gitea-release-change-log",
		Usage:     "update/create gitea release change log drone plugin",
		Copyright: "© 2021-now hb0730",
		Action:    run,
		Version:   version,
		Authors: []*cli.Author{
			{
				Name:  "hb0730",
				Email: "huangbing0730@gmail.com",
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "debug",
				Usage:   "debug mode",
				EnvVars: []string{"PLUGIN_DEBUG", "DEBUG"},
			},
			&cli.StringFlag{
				Name:    "drone.tag",
				Usage:   "Provides the tag for the current running build. This value is only populated for tag events and promotion events that are derived from tags",
				EnvVars: []string{"DRONE_TAG"},
			},
			&cli.StringFlag{
				Name:    "drone.repo",
				Usage:   "Provides the full repository name for the current running build",
				EnvVars: []string{"DRONE_REPO"},
			},
			&cli.StringFlag{
				Name:    "drone.repo.name",
				Usage:   "Provides the repository name for the current running build",
				EnvVars: []string{"DRONE_REPO_NAME"},
			},
			&cli.StringFlag{
				Name:    "drone.repo.branch",
				Usage:   "Provides the default repository branch for the current running build.",
				EnvVars: []string{"DRONE_REPO_BRANCH"},
			},
			&cli.StringFlag{
				Name:    "drone.repo.owner",
				Usage:   "Provides the repository namespace for the current running build. The namespace is an alias for the source control management account that owns the repository",
				EnvVars: []string{"DRONE_REPO_OWNER"},
			},
			&cli.StringFlag{
				Name:    "drone.commit",
				Usage:   "Provides the git commit sha for the current running build.",
				EnvVars: []string{"DRONE_COMMIT"},
			},
			&cli.StringFlag{
				Name:    "gitea.url",
				Usage:   "gite base url",
				EnvVars: []string{"PLUGIN_GITEA_URL", "PLUGIN_URL", "GITEA_URL,URL"},
			},
			&cli.StringFlag{
				Name:    "gitea.token",
				Usage:   "gitea user token",
				EnvVars: []string{"PLUGIN_GITEA_TOKEN", "PLUGIN_TOKEN", "GITEA_TOKEN", "TOKEN"},
			},
			&cli.StringFlag{
				Name:    "changelog.config",
				Usage:   "the YAML config file for generate changelog",
				EnvVars: []string{"PLUGIN_CHANGE_LOG_CONFIG", "PLUGIN_LOG_CONFIG", "CHANGE_LOG_CONFIG", "CONFIG"},
			},
			&cli.StringFlag{
				Name:    "changelog.config_str",
				Usage:   "the YAML string type for generate changelog,priority is lower than file config",
				EnvVars: []string{"PLUGIN_CHANGE_LOG_CONFIG_STR", "PLUGIN_LOG_CONFIG_STR", "CHANGE_LOG_CONFIG_STR", "CONFIG_STR"},
			},
			&cli.IntFlag{
				Name:    "changelog.tag_type",
				Usage:   "repo tags sort type,default: 1 creatordate sort",
				EnvVars: []string{"PLUGIN_CHANGE_LOG_TAG_TYPE", "PLUGIN_TAG_TYPE", "CHANGE_LOG_TAG_TYPE", "LOG_TAG_TYPE"},
				Value:   1,
			},
			&cli.StringFlag{
				Name:    "changelog.repo_path",
				Usage:   "git repo path,default:./",
				EnvVars: []string{"PLUGIN_CHANGE_LOG_REPO_PATH", "PLUGIN_REPO_PATH", "CHANGE_LOG_REPO_PATH", "REPO_PATH"},
				Value:   "./",
			},
			&cli.StringFlag{
				Name:    "changelog.sha1",
				Usage:   "The old git sha version. allow: tag name, commit id",
				EnvVars: []string{"PLUGIN_CHANGE_LOG_SHA1", "PLUGIN_SHA1", "CHANGE_LOG_SHA1", "SHA1"},
				Value:   "prev",
			},
			&cli.StringFlag{
				Name:    "changelog.sha2",
				Usage:   "The new git sha version. allow: tag name, commit id",
				EnvVars: []string{"PLUGIN_CHANGE_LOG_SHA2,PLUGIN_SHA2", "CHANGE_LOG_SHA2", "SHA2"},
				Value:   "last",
			},
			&cli.BoolFlag{
				Name:    "changelog.verbose",
				Usage:   "show more information",
				EnvVars: []string{"PLUGIN_CHANGE_LOG_VERBOSE", "PLUGIN_LOG_VERBOSE", "CHANGE_LOG_VERBOSE", "VERBOSE"},
			},
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
			RepoBrand:  ctx.String("drone.repo.branch"),
			Owner:      ctx.String("drone.repo.owner"),
			CommitHash: ctx.String("drone.commit"),
		},
		Gitea: Gitea{
			URL:   ctx.String("gitea.url"),
			Token: ctx.String("gitea.token"),
		},

		ChangeLogConfig: ChangeLogConfig{
			ConfigFile: ctx.String("changelog.config"),
			ConfigStr:  ctx.String("changelog.config_str"),
			TagType:    ctx.Int("changelog.tag_type"),
			RepoPath:   ctx.String("changelog.repo_path"),
			Sha1:       ctx.String("changelog.sha1"),
			Sha2:       ctx.String("changelog.sha2"),
			Verbose:    ctx.Bool("changelog.verbose"),
		},
	}
	return plugin.Exec(ctx)
}
