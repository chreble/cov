// Copyright (c) 2016, Hotolab. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/urfave/cli"

	"github.com/hotolab/cov"
)

// ProfileCommand converts and outputs a golang coverage profile
// to JSON
func ProfileCommand() cli.Command {
	return cli.Command{
		Name:  "profile",
		Usage: "Converts a golang coverage profile to a JSON output",
		Action: func(c *cli.Context) error {
			if c.Args().Get(0) == "" {
				return errors.New("The coverage profile path must be passed to the command")
			}
			r, err := cov.ConvertProfile(c.Args().Get(0))
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
