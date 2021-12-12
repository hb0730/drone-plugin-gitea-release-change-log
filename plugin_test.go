package main

import (
	"os"
	"testing"
)

func TestPlugin_Exec(t *testing.T) {
	p := &Plugin{
		Debug: true,
		Drone: Drone{
			Tag:        os.Getenv("tag"),
			Repo:       os.Getenv("repo"),
			Owner:      os.Getenv("owner"),
			CommitHash: os.Getenv("commitHash"),
		},
		Gitea: Gitea{
			Token: os.Getenv("token"),
			URL:   os.Getenv("url"),
		},
	}
	p.Exec(nil)
}
