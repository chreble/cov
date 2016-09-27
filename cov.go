// Copyright (c) 2016, Hotolab. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cov

import (
	"os"

	"golang.org/x/tools/cover"
)

// ConvertRepository converts a given repository to a Report struct
func ConvertRepository(repo string) (*Report, error) {
	r := &Report{}
	err := r.collectPackages()
	if err != nil {
		return nil, err
	}

	p, err := createProfile()
	if err != nil {
		return nil, err
	}
	defer os.Remove(p.Name())

	profiles, err := cover.ParseProfiles(p.Name())
	if err != nil {
		return nil, err
	}

	if err = r.parseProfile(profiles); err != nil {
		return nil, err
	}

	return r, nil
}

// ConvertProfile converts a given profile to a Report struct
func ConvertProfile(filename string) (*Report, error) {
	profiles, err := cover.ParseProfiles(filename)
	if err != nil {
		return nil, err
	}

	r := &Report{}
	if err = r.parseProfile(profiles); err != nil {
		return nil, err
	}

	return r, nil
}
