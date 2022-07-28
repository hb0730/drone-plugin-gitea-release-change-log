package main

import (
	"fmt"
	"github.com/gookit/gitw"
	"github.com/gookit/gitw/chlog"
	"github.com/gookit/goutil"
	"testing"
)

func TestGitw(t *testing.T) {
	repo := gitw.NewRepo("./")
	cl := chlog.New()
	sha1 := repo.AutoMatchTag("prev")
	sha2 := repo.AutoMatchTag("last")
	cl.FetchGitLog(sha1, sha2)
	cl.Generate()
	t.Log(cl.Changelog())
}
func TestGitChangelog(t *testing.T) {
	cl := chlog.NewWithConfig(chlog.NewDefaultConfig())
	cl.WithConfigFn(func(cfg *chlog.Config) {
		cfg.RepoURL = "https://github.com/hb0730/drone-plugin-gitea-release-change-log"
	})
	repo := gitw.NewRepo("./")
	sha1 := repo.AutoMatchTag("prev")
	sha2 := repo.AutoMatchTag("last")
	// fetch git log
	cl.FetchGitLog(sha1, sha2)

	// do generate
	goutil.PanicIfErr(cl.Generate())

	// dump
	fmt.Println(cl.Changelog())
}

func TestGitTags(t *testing.T) {
	cl := chlog.NewWithConfig(chlog.NewDefaultConfig())
	cl.WithConfigFn(func(cfg *chlog.Config) {
		cfg.RepoURL = "https://github.com/hb0730/drone-plugin-gitea-release-change-log"
	})
	repo := gitw.NewRepo("./")
	sha1, err := repo.Cmd("describe", "--tags", "--abbrev=0").Output()
	if err != nil {
		t.Error(err)
	}
	sha1 = gitw.FirstLine(sha1)
	sha2, err := repo.Cmd("describe", "--tags", "--abbrev=0", fmt.Sprintf("tags/%s^", sha1)).Output()
	if err != nil {
		t.Error(err)
	}
	sha2 = gitw.FirstLine(sha2)
	t.Logf("sh1: %s,sha2: %s \n", sha1, sha2)
}
