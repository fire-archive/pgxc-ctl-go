/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. 
 * Copyright (c) 2014, K. S. Ernest "iFire" Lee */

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