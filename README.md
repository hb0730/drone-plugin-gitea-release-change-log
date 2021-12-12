# drone-plugin-gitea-release-change-log
Drone plugin for creating and tagging Gitea releases change logs

# example
```yaml
- name: fetch
  image: alpine/git
  commands:
    - git fetch --tags
- name: release change log
  image: hb0730/drone-plugin-gitea-release-change-log
  settings:
    debug: true
    url: https://gitea.io
    token: <user token>
  when:
    event:
      - tag 
```
# plugin params
* `url`: gitea base url
* `token`: gitea user access token