// Copyright (c) 2016, Hotolab. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cov

import (
	"sort"

	"golang.org/x/tools/cover"
)

// Report contains information about tested packages, functions and statements
type Report struct {
	// Packages holds all tested packages
	Packages []*Package
	// TLOC contains the sum of all TLOCs
	TLOC int64
	// Coverage is the global test coverage percentage
	Coverage float64
}

// ConvertProfile converts a given profile to a Report struct
func ConvertProfile(filename string) (report *Report, e error) {
	profiles, err := cover.ParseProfiles(filename)
	if err != nil {
		return nil, err
	}

	r := &Report{}
	return r.parseProfile(profiles)
}

func (r *Report) parseProfile(profiles []*cover.Profile) (report *Report, e error) {
	conv := converter{
		packages: make(map[string]*Package),
	}
	for _, p := range profiles {
		if err := conv.convertProfile(p); err != nil {
			return nil, err
		}
	}
	for _, pkg := range conv.packages {
		r.addPackage(pkg)
	}
	r.computeGlobalCoverage()

	return r, nil
}

func (r *Report) computeGlobalCoverage() {
	// Loop on each package and determine coverage and TLOC by package
	var gcov float64
	var tloc int64
	for _, pkg := range r.Packages {
		gcov += pkg.Coverage
		tloc += pkg.TLOC
	}

	// Report the global # of tested LOCs
	r.TLOC = tloc
	if gcov > 0 {
		r.Coverage = gcov / float64(len(r.Packages))
	}
}

// AddPackage adds a package coverage information
func (r *Report) addPackage(p *Package) {
	i := sort.Search(len(r.Packages), func(i int) bool {
		return (r.Packages)[i].Name >= p.Name
	})
	if i < len(r.Packages) && (r.Packages)[i].Name == p.Name {
		(r.Packages)[i].Accumulate(p)
	} else {
		head := (r.Packages)[:i]
		tail := append([]*Package{p}, (r.Packages)[i:]...)
		r.Packages = append(head, tail...)
	}
}
