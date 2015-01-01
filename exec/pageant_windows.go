package exec

import (
	"github.com/davidmz/go-pageant"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

var ag agent.Agent

func pageant_init() {
	ag = pageant.New()
}

func pagent_agent() agent.Agent {
	return ag
}

func pageant_auth() (auths []ssh.AuthMethod) {
	auths = []ssh.AuthMethod{ssh.PublicKeysCallback(ag.Signers)}
	return auths
}
