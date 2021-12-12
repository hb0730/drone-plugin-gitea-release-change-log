package main

import (
	"code.gitea.io/sdk/gitea"
	"errors"
	"github.com/apex/log"
	"github.com/hb0730/git-change-log"
)

type ChangeLog struct {
	Debug      bool
	git        *gitea.Client
	CurrentTag string
	Drone      Drone
}

func NewChangeLog(url, token, currentTag string, debug bool, drone Drone) (ChangeLog, error) {
	var changelog ChangeLog
	if url == "" || token == "" {
		return changelog, errors.New("gitea url or token missing")
	}
	if currentTag == "" {
		return changelog, errors.New("current tag missing")
	}
	client, err := gitea.NewClient(url, gitea.SetToken(token))
	if err != nil {
		return changelog, err
	}
	changelog.git = client
	changelog.CurrentTag = currentTag
	changelog.Drone = drone
	changelog.Debug = debug
	return changelog, nil
}
func (l ChangeLog) PutRelease() error {
	if l.Drone.Repo == "" || l.Drone.Owner == "" {
		return errors.New("gitea repo or repo owner missing")
	}
	changelog, err := l.getChangeLogs()
	if err != nil {
		return err
	}
	release, resp, err := l.git.GetReleaseByTag(l.Drone.Owner, l.Drone.RepoName, l.CurrentTag)
	if resp.StatusCode == 404 {
		option := gitea.CreateReleaseOption{
			TagName: l.CurrentTag,
			Target:  l.Drone.CommitHash,
			Title:   l.CurrentTag,
			Note:    changelog,
		}
		_, _, err = l.git.CreateRelease(l.Drone.Owner, l.Drone.RepoName, option)
	} else if resp.StatusCode == 200 {
		option := gitea.EditReleaseOption{

			TagName:      release.TagName,
			Target:       release.Target,
			Title:        release.Title,
			Note:         changelog,
			IsDraft:      &release.IsDraft,
			IsPrerelease: &release.IsPrerelease,
		}
		_, _, err = l.git.EditRelease(l.Drone.Owner, l.Drone.RepoName, release.ID, option)
	} else {
		return err
	}
	return err
}
func (l ChangeLog) getChangeLogs() (string, error) {
	var level log.Level
	if l.Debug {
		level = log.DebugLevel
	} else {
		level = log.InfoLevel
	}
	return git.GetChangeLogs("", l.CurrentTag, level)
}
