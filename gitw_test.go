package main

import (
	"github.com/gookit/gitw"
	"github.com/gookit/gitw/chlog"
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
