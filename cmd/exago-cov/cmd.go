// Copyright (c) 2016, Hotolab. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"

	. "github.com/exago/cov/config"
	"github.com/exago/cov"
)

var (
	App *cli.App
)

// Initialize commandline app.
func init() {
	App = cli.NewApp()

	// For fancy output on console
	App.Name = "exago cov"
	App.Usage = `Check -h`
	App.Author = "Hotolab <dev@hotolab.com>"

	// Version is injected at build-time
	App.Version = ""

	App.Action = func(c *cli.Context) {
		r, err := cov.ConvertProfile(Config.Profile)
		if err != nil {
			log.Fatal(err)
		}

		bytes, err := json.Marshal(r)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(bytes))
	}

	InitializeConfig()
	InitializeLogging(Config.LogLevel)
}

func main() {
	if err := App.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// InitializeLogging sets logrus log level.
func InitializeLogging(logLevel string) {
	// If log level cannot be resolved, exit gracefully
	if logLevel == "" {
		log.Warning("Log level could not be resolved, fallback to fatal")
		log.SetLevel(log.FatalLevel)
		return
	}
	// Parse level from string
	lvl, err := log.ParseLevel(logLevel)

	if err != nil {
		log.WithFields(log.Fields{
			"passed":  logLevel,
			"default": "fatal",
		}).Warn("Log level is not valid, fallback to default level")
		log.SetLevel(log.FatalLevel)
		return
	}

	log.SetLevel(lvl)
	log.WithFields(log.Fields{
		"level": logLevel,
	}).Debug("Log level successfully set")
}
