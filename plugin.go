package main

import (
	"errors"
	"github.com/urfave/cli"
	"log"
)

type Plugin struct {
	Debug bool
	Drone Drone
	Gitea Gitea
}
type Drone struct {
	Tag        string
	Repo       string
	RepoName   string
	Owner      string
	CommitHash string
}
type Gitea struct {
	URL   string
	Token string
}

func (p *Plugin) Exec(_ *cli.Context, config ChangeLogConfig) error {
	if p.Drone.Tag == "" {
		log.Println("Skipping gitea release change log")
		return nil
	}
	if p.Gitea.URL == "" || p.Gitea.Token == "" {
		return errors.New("gitea url or token missing")
	}
	if p.Debug {
		log.Println("DRONE:{ tag:", p.Drone.Tag, ",repo:", p.Drone.Repo, "owner:", p.Drone.Owner, "commit hash:", p.Drone.CommitHash, "}")
		log.Println("gitea:{ url:", p.Gitea.URL, "token:", p.Gitea.Token, "}")
		log.Println("changelog config:{config:", config.ConfigFile, "sha1:", config.Sha1, "sha2:", config.Sha2, "}")
	}
	changeLog, err := NewChangeLog(p.Gitea.URL, p.Gitea.Token, p.Drone.Tag, p.Debug,
		p.Drone)
	if err != nil {
		return err
	}
	return changeLog.PutRelease(config)
}
