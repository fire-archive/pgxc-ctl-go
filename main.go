/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright (c) 2014, K. S. Ernest "iFire" Lee */

package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/fire/pgxc-ctl-go/exec"
)

func main() {
	app := cli.NewApp()
	app.Name = "pgxc"
	app.Usage = "Controls a postgresqlxl cluster"
	app.Action = func(c *cli.Context) {
		var ai exec.Auth_info
		ai.Username = "admin"
		ai.Server = "192.168.1.81:22"
		var cmds []string
		cmds = append(cmds, "/usr/bin/env whoami")
		cmds = append(cmds, "/usr/bin/env ifconfig")
		exec.Execute(ai, cmds)
	}
	app.Run(os.Args)
}
