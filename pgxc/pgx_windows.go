package pgxc

import (
        "github.com/fire/pgxc-ctl-go/pageant"
	"golang.org/x/crypto/ssh"
)

func Config(username string) (*ssh.ClientConfig){
	pageant.Init()
	config := &ssh.ClientConfig{
		User: username,
		Auth: pageant.Auth(),
	}
        return config
}