package main

import (
	"github.com/gookit/gitw/chlog"
	"gopkg.in/yaml.v2"
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

func TestChangeLog_ParseConfigStr(t *testing.T) {
	var cfg = chlog.NewDefaultConfig()
	config := ``
	yml := []byte(config)
	if len(yml) > 0 {
		if err := yaml.Unmarshal(yml, &cfg); err != nil {
			t.Error(err)
		}
	}
	config = `title: '## Change Log'
# style allow: simple, markdown(mkdown), ghr(gh-release)
style: gh-release
# group names
names: [Refactor, Fixed, Feature, Update, Other]
# if empty will auto fetch by git remote
#repo_url: https://github.com/gookit/gitw

filters:
  # message length should >= 12
  - name: msg_len
    min_len: 12
  # message words should >= 3
  - name: words_len
    min_len: 3
  - name: keyword
    keyword: format code
    exclude: true
  - name: keywords
    keywords: format code, action test
    exclude: true

# group match rules
# not matched will use 'Other' group.
rules:
  - name: Refactor
    start_withs: [refactor, break]
    contains: ['refactor:']
  - name: Fixed
    start_withs: [fix]
    contains: ['fix:']
  - name: Feature
    start_withs: [feat, new]
    contains: ['feat:']
  - name: Update
    start_withs: [update, 'up:']
    contains: ['update:']`
	yml = []byte(config)
	if len(yml) > 0 {
		if err := yaml.Unmarshal(yml, &cfg); err != nil {
			t.Error(err)
		}
	}
	t.Log(cfg.Title)
}
