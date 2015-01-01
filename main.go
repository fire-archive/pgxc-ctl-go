/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright (c) 2014, K. S. Ernest "iFire" Lee */

package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fire/pgxc-ctl-go/exec"
	"log"
	"os"
)

type gtm struct {
	Location string
}

func main() {
	var ai exec.Auth_info
	ai.Username = "admin"
	ai.Server = "192.168.1.81:22"

	app := cli.NewApp()
	app.Name = "pgxc"
	app.Usage = "Controls a postgresqlxl cluster"
	app.EnableBashCompletion = true
	profileAddPgxc := "PATH=$PATH:/usr/local/pgsql/bin"
	app.Commands = []cli.Command{
		{
			Name:      "run",
			ShortName: "r",
			Usage:     "run a shell command",
			Action: func(c *cli.Context) {
				var cmds []string
				cmds = append(cmds, c.Args().First())
				log.Print(cmds)
				exec.Execute(ai, cmds)
			},
		},
		{
			Name:  "start",
			Usage: "add new nodes",
			Subcommands: []cli.Command{
				{
					Name:  "gtm",
					Usage: "start gtm",
					Subcommands: []cli.Command{
						{
							Name:  "master",
							Usage: "start gtm master node",
							Action: func(c *cli.Context) {
								if c.Args().First() != "" {
									var a exec.Auth_info
									a.Username = ai.Username
									a.Server = c.Args().First()
									var g gtm
									g.Location = "/home/" + ai.Username + "/pgxcgo/nodes/gtm"
									var cmds []string
									cmds = append(cmds, profileAddPgxc)
									cmds = append(cmds, "/usr/bin/env gtm_ctl -Z gtm start -D " + g.Location+" &2>1")
									exec.Execute(a, cmds)
								} else {
									fmt.Println("Usage: start gtm master localhost:80")
								}
							},
						},
					},
				},
			},
		},
		{
			Name:  "init",
			Usage: "init new nodes",
			Subcommands: []cli.Command{
				{
					Name:  "gtm",
					Usage: "initialize a new gtm (will remove existing directory)",
					Subcommands: []cli.Command{
						{
							Name:  "master",
							Usage: "initialize gtm master node (will remove existing directory)",
							Action: func(c *cli.Context) {
								if c.Args().First() != "" {
									var a exec.Auth_info
									a.Username = ai.Username
									a.Server = c.Args().First()
									var cmds []string
									var g gtm
									g.Location = "/home/" + ai.Username + "/pgxcgo/nodes/gtm"
									var mkdircmds []string
									mkdircmds = append(mkdircmds, "/usr/bin/env mkdir -p "+g.Location)
									exec.Execute(a, mkdircmds)
									cmds = append(cmds, profileAddPgxc)
									cmds = append(cmds, "/usr/bin/env initgtm -Z gtm -D "+g.Location+" 2>&1")
									exec.Execute(a, cmds)
								} else {
									fmt.Println("Usage: init gtm master localhost:80")
								}
							},
						},
					},
				},
			},
		},
		{
			Name:  "stop",
			Usage: "stop node",
			Subcommands: []cli.Command{
				{
					Name:  "gtm",
					Usage: "stop gtm",
					Subcommands: []cli.Command{
						{
							Name:  "master",
							Usage: "stop master node",
							Action: func(c *cli.Context) {
								if c.Args().First() != "" {
									var a exec.Auth_info
									a.Username = ai.Username
									a.Server = c.Args().First()
									var g gtm
									g.Location = "/home/" + ai.Username + "/pgxcgo/nodes/gtm"
									var cmds []string
									cmds = append(cmds, profileAddPgxc)
									cmds = append(cmds, "/usr/bin/env gtm_ctl -Z gtm stop -D " + g.Location +" &2>1")
									exec.Execute(a, cmds)
								} else {
									fmt.Println("Usage: stop gtm master localhost:80")
								}
							},
						},
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
