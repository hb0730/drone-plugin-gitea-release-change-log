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
	sha1 := repo.AutoMatchTag("1.0.0")
	sha2 := repo.AutoMatchTag("2.0.0")
	// fetch git log
	cl.FetchGitLog(sha1, sha2)

	// do generate
	goutil.PanicIfErr(cl.Generate())

	// dump
	fmt.Println(cl.Changelog())
}
