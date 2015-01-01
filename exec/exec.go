package exec

import (
	"bytes"
	"fmt"
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
	if err := session.Run(strings.Join(cmds, ";")); err != nil {
		fmt.Println(cmds)
		panic("Failed to run: " + err.Error())
	}
	fmt.Println(b.String())
}
