package main

import (
	"github.com/apex/log"
	"github.com/hb0730/git-change-log"
	"testing"
)

func TestChangeLogs(t *testing.T) {
	changelog, err := git.GetChangeLogs("", "1.0.0-beta", log.DebugLevel)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(changelog)
}
