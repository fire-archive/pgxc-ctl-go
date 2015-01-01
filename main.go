/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 * Copyright (c) 2014, K. S. Ernest "iFire" Lee */

package main

import (
	"bytes"
	"fmt"
	"github.com/fire/pgxc-ctl-go/pgxc"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
)

// http://golang-basic.blogspot.ca/2014/06/step-by-step-guide-to-ssh-using-go.html
// http://play.golang.org/p/kMhHvbl4SG
type auth_info struct {
	username string
	server string
}

func execute(ai auth_info, cmds []string) {
	var unixConn net.Conn
	if runtime.GOOS != "windows" {
		var err error
		unixConn, err = net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	client, err := ssh.Dial("tcp", ai.server, pgxc.Config(ai.username, unixConn))
	if err != nil {
		log.Fatalln("Failed to connect:", err)
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	defer client.Close()
	if runtime.GOOS != "windows" {
		defer unixConn.Close()
	}
	
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(strings.Join(cmds, ";")); err != nil {
		fmt.Println(cmds)
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}

func main() {
	var ai auth_info
	ai.username = "admin"
	ai.server   = "192.168.1.81:22"
	var cmds []string
	cmds = append(cmds, "/usr/bin/env ls")
	cmds = append(cmds, "/usr/bin/env ifconfig")
	execute(ai, cmds)
}
