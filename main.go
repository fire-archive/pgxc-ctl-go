/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright (c) 2014, K. S. Ernest "iFire" Lee */

package main

import (
	"github.com/fire/pgxc-ctl-go/pgxc"
)

func main() {
	var ai pgxc.Auth_info
	ai.Username = "admin"
	ai.Server = "192.168.1.81:22"
	var cmds []string
	cmds = append(cmds, "/usr/bin/env ls")
	cmds = append(cmds, "/usr/bin/env ifconfig")
	pgxc.Execute(ai, cmds)
}
