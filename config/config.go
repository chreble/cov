// Copyright (c) 2016, Hotolab. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/exago/envconfig"
)

// Config holds the configuration object
var Config cfg

type cfg struct {
	Profile  string `envconfig:"PROFILE" required:"true"`
	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}

// InitializeConfig will populate Config variable from environment variables
func InitializeConfig() {
	if err := envconfig.Process("", &Config); err != nil {
		log.Fatal(err)
	}
}
