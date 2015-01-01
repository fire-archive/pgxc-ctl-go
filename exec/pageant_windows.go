/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. 
 * Copyright (c) 2014, K. S. Ernest "iFire" Lee */

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
