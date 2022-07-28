package main

import (
	"testing"
)

func TestChangeLog_ChangeLogs(t *testing.T) {

	cl := ChangeLog{
		CurrentTag: "2.0.1",
		Debug:      true,
	}
	cl.ChangeLogs(ChangeLogConfig{
		RepoPath: "./",
		Sha1:     "prev",
		Sha2:     "last",
		Verbose:  true,
	})
}
