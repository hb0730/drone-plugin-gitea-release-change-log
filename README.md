# drone-plugin-gitea-release-change-log

Drone plugin for creating and tagging Gitea releases change logs

# example

```yaml
-   name: fetch
    image: alpine/git
    commands:
        - git fetch --tags
-   name: release change log
    image: hb0730/drone-plugin-gitea-release-change-log
    settings:
        debug: true
        url: https://gitea.io
        token: <user token>
        change_log_verbose: true
    when:
        event:
            - tag 
```

# plugin params

* `url`: gitea base url
* `token`: gitea user access token
* `change_log_config`: the YAML config file for generate changelog
* `change_log_config_str`: the YAML string type for generate changelog,priority is lower than file config
* `change_log_tag_type`: sort git tag type,default: 1 creatordate sort, see [gitw](https://github.com/gookit/gitw)
* `change_log_repo_path`: [gitw](https://github.com/gookit/gitw) create repo path
* `change_log_sha1`: The old git sha version. allow: tag name, commit id,default: prev
* `change_log_sha2`: The new git sha version. allow: tag name, commit id,default: last
* `change_log_verbose`: show gitw more information

# change log

see [gitw](https://github.com/gookit/gitw)

![img.png](https://raw.githubusercontent.com/hb0730/drone-plugin-gitea-release-change-log/main/doc/img.png)