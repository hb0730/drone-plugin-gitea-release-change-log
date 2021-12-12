package main

import (
	"github.com/urfave/cli"
	"log"
)

type Plugin struct {
	Debug bool
	Drone Drone
}
type Drone struct {
	Tag string
}

func (p *Plugin) Exec(ctx *cli.Context) error {
	if p.Drone.Tag == "" {
		log.Println("Skipping gitea release change log")
	}

	return nil
}
