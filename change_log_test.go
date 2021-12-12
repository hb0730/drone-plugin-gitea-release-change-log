package main

import (
	"github.com/hb0730/git-change-log"
	"testing"
)

func TestChangeLogs(t *testing.T) {
	changelog, err := git.GetChangeLogs("", "1.0.0-beta")
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Log(changelog)
}
