FROM alpine/git
COPY ./drone-plugin-gitea-release-change-log /bin
ENTRYPOINT ["/bin/drone-plugin-gitea-release-change-log"]
