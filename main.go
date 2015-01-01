/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright (c) 2014, K. S. Ernest "iFire" Lee */

package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/fire/pgxc-ctl-go/exec"
	"log"
)

func main() {
	var ai exec.Auth_info
	ai.Username = "admin"
	ai.Server = "192.168.1.81:22"

	app := cli.NewApp()
	app.Name = "pgxc"
	app.Usage = "Controls a postgresqlxl cluster"
	app.EnableBashCompletion = true
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
	}
	app.Run(os.Args)
}
