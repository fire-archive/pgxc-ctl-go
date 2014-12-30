/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

package main

import (
	"bytes"
	"fmt"
	"github.com/davidmz/go-pageant"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"
	"runtime"
)

// http://golang-basic.blogspot.ca/2014/06/step-by-step-guide-to-ssh-using-go.html
// http://play.golang.org/p/kMhHvbl4SG

func getKeyFile() (key ssh.Signer, err error) {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.ssh/id_rsa"
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil {
		return
	}
	return
}

const (
	username = "admin"
	server   = "192.168.1.79:22"
)

func main() {
	var ag agent.Agent
        var auths []ssh.AuthMethod
	if runtime.GOOS == "windows" {
		ag = pageant.New()
                auths = []ssh.AuthMethod{ssh.PublicKeysCallback(ag.Signers)}
	} else {
        	conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()
		ag = agent.NewClient(conn)
        }

	config := &ssh.ClientConfig{
		User: username,
		Auth: auths,
	}
	client, err := ssh.Dial("tcp", server, config)
	if err != nil {
		log.Fatalln("Failed to connect via ssh:", err)
	}

	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/env whoami && /usr/bin/env ifconfig"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
