// Copyright (c) 2016, Hotolab. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/hotolab/cov"
)

// RepoCommand prepares a coverage file
// converts and outputs to JSON
func RepoCommand() cli.Command {
	return cli.Command{
		Name:  "repo",
		Usage: "Generates a golang coverage profile and convert it to a JSON output",
		Action: func(c *cli.Context) error {
			repo := c.Args().Get(0)
			if repo == "" {
				return errors.New("The repository must be passed to the command")
			}
			if err := changeDir(repo); err != nil {
				return err
			}

			r, err := cov.ConvertRepository(repo)
			if err != nil {
				return err
			}

			bytes, err := json.Marshal(r)
			if err != nil {
				return err
			}

			fmt.Println(string(bytes))
			return nil
		},
	}
}

// changeDir takes the repository passed to the function and
// makes it the current working directory
func changeDir(repo string) error {
	dir := fmt.Sprintf("%s/src/%s", os.Getenv("GOPATH"), repo)
	return os.Chdir(dir)
}
