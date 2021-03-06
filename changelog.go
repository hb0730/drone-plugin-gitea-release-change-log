package main

import (
	"code.gitea.io/sdk/gitea"
	"errors"
	"fmt"
	gitw "github.com/gookit/gitw"
	"github.com/gookit/gitw/chlog"
	"github.com/gookit/goutil/cliutil"
	"github.com/gookit/goutil/dump"
	"github.com/gookit/goutil/fsutil"
	"gopkg.in/yaml.v2"
)

type ChangeLog struct {
	Debug      bool
	gitea      *gitea.Client
	repo       *gitw.Repo
	CurrentTag string
	Drone      Drone
}

func NewChangeLog(plugin Plugin) (ChangeLog, error) {
	var changelog ChangeLog
	giteaConf := plugin.Gitea
	if giteaConf.URL == "" || giteaConf.Token == "" {
		return changelog, errors.New("gitea url or token missing")
	}
	drone := plugin.Drone
	if drone.Tag == "" {
		return changelog, errors.New("current tag missing")
	}
	client, err := gitea.NewClient(giteaConf.URL, gitea.SetToken(giteaConf.Token))
	if err != nil {
		return changelog, err
	}
	changelog.gitea = client
	changelog.repo = gitw.NewRepo(plugin.ChangeLogConfig.RepoPath)
	changelog.CurrentTag = drone.Tag
	changelog.Drone = drone
	changelog.Debug = plugin.Debug
	return changelog, nil
}
func (l ChangeLog) PutRelease(config ChangeLogConfig) error {
	if l.Drone.Repo == "" || l.Drone.Owner == "" {
		return errors.New("gitea repo or repo owner missing")
	}
	changelog, err := l.getChangeLogs(config)
	if err != nil {
		return err
	}
	release, resp, err := l.gitea.GetReleaseByTag(l.Drone.Owner, l.Drone.RepoName, l.CurrentTag)
	if resp.StatusCode == 404 {
		option := gitea.CreateReleaseOption{
			TagName: l.CurrentTag,
			Target:  l.Drone.CommitHash,
			Title:   l.CurrentTag,
			Note:    changelog,
		}
		_, _, err = l.gitea.CreateRelease(l.Drone.Owner, l.Drone.RepoName, option)
	} else if resp.StatusCode == 200 {
		option := gitea.EditReleaseOption{

			TagName:      release.TagName,
			Target:       release.Target,
			Title:        release.Title,
			Note:         changelog,
			IsDraft:      &release.IsDraft,
			IsPrerelease: &release.IsPrerelease,
		}
		_, _, err = l.gitea.EditRelease(l.Drone.Owner, l.Drone.RepoName, release.ID, option)
	} else {
		return err
	}
	return err
}
func (l ChangeLog) getChangeLogs(config ChangeLogConfig) (string, error) {
	cfg := chlog.NewDefaultConfig()
	err := loadConfig(config, l.repo, cfg)
	if err != nil {
		return "", err
	}
	cl := chlog.NewWithConfig(cfg)

	sha1 := l.repo.AutoMatchTag(config.Sha1)
	sha2 := l.repo.AutoMatchTag(config.Sha2)
	cl.FetchGitLog(sha1, sha2)
	err = cl.Generate()
	if err != nil {
		return "", err
	}
	changelog := cl.Changelog()
	if l.Debug {
		fmt.Printf("change_log: %s \n", changelog)
	}
	return changelog, nil
}
func loadConfig(config ChangeLogConfig, repo *gitw.Repo, cfg *chlog.Config) error {
	yml := fsutil.ReadExistFile(config.ConfigFile)
	if len(yml) > 0 {
		err := yaml.Unmarshal(yml, cfg)
		if err != nil {
			return err
		}
	}
	if cfg.RepoURL == "" {
		cfg.RepoURL = repo.DefaultRemoteInfo().URLOfHTTPS()
	}

	if config.Verbose {
		cfg.Verbose = true
		cliutil.Cyanln("Changelog Config:")
		dump.NoLoc(cfg)
		fmt.Println()
	}
	return nil
}
