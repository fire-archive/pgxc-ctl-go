package pageant

import (
	"github.com/davidmz/go-pageant"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

var ag agent.Agent

func Init() {
	ag = pageant.New()
}

func Agent() agent.Agent {
	return ag
}

func Auth() (auths []ssh.AuthMethod) {
	auths = []ssh.AuthMethod{ssh.PublicKeysCallback(ag.Signers)}
	return auths
}
