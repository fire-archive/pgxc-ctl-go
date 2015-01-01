/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. 
 * Copyright (c) 2014, K. S. Ernest "iFire" Lee */

package exec

import (
	"bytes"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
)

// http://golang-basic.blogspot.ca/2014/06/step-by-step-guide-to-ssh-using-go.html
// http://play.golang.org/p/kMhHvbl4SG
type Auth_info struct {
	Username string
	Server   string
}

func Execute(ai Auth_info, cmds []string) {
	var unixConn net.Conn
	if runtime.GOOS != "windows" {
		var err error
		unixConn, err = net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
		if err != nil {
			log.Fatalf(err.Error())
		}
	}
	client, err := ssh.Dial("tcp", ai.Server, Config(ai.Username, unixConn))
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
	expanded_cmds := strings.Join(cmds, ";")
	log.Println(expanded_cmds)
	if err := session.Run(expanded_cmds); err != nil {
		log.Println("Failed to run: " + err.Error())
	}
	log.Println(b.String())
}
