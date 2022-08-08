package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
)

type Plugin struct {
	Debug           bool
	Drone           Drone
	Gitea           Gitea
	ChangeLogConfig ChangeLogConfig
}
type Drone struct {
	Tag        string
	Repo       string
	RepoName   string
	Owner      string
	CommitHash string
}

func (d Drone) toString() string {
	return fmt.Sprintf(
		"DRONE: {tag: %s,repo:%s ,repo_name: %s,owner: %s,commit_hash: %s}",
		d.Tag, d.Repo, d.RepoName, d.Owner, d.CommitHash)
}

type Gitea struct {
	URL   string
	Token string
}

func (g Gitea) toString() string {
	return fmt.Sprintf("GITEA: {url: %s,token: %s}", g.URL, g.Token)
}

type ChangeLogConfig struct {
	ConfigFile string
	ConfigStr  string
	TagType    int
	RepoPath   string
	Sha1       string
	Sha2       string
	Verbose    bool
}

func (c ChangeLogConfig) toString() string {
	return fmt.Sprintf("CHANGE_LOG_CONFIG: {config_file: %s,config_str: %s,tag_sort:%d, repo_path: %s,sha1: %s,sha2: %s,verbose: %t }",
		c.ConfigFile, c.ConfigStr, c.TagType, c.RepoPath, c.Sha1, c.Sha2, c.Verbose)
}

func (p *Plugin) Exec(_ *cli.Context) error {
	if p.Drone.Tag == "" {
		log.Println("Skipping gitea release change log")
		return nil
	}
	if p.Debug {
		log.Println(p.Drone.toString())
		log.Println(p.Gitea.toString())
		log.Println(p.ChangeLogConfig.toString())
	}
	changeLog, err := NewChangeLog(*p)
	if err != nil {
		return err
	}
	return changeLog.PutRelease(p.ChangeLogConfig)
}
