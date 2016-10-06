// Copyright (c) 2016, Hotolab. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cov

import (
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"sort"
	"strings"

	"golang.org/x/tools/cover"

	log "github.com/Sirupsen/logrus"
)

// Report contains information about tested packages, functions and statements
type Report struct {
	// Packages holds all tested packages
	Packages []*Package `json:"packages"`
	// Coverage is the global test coverage percentage
	Coverage float64 `json:"coverage"`
}

func (r *Report) parseProfile(profiles []*cover.Profile) error {
	conv := converter{
		packages: make(map[string]*Package),
	}
	for _, p := range profiles {
		if err := conv.convertProfile(p); err != nil {
			return err
		}
	}
	for _, pkg := range conv.packages {
		r.addPackage(pkg)
	}

	r.computeGlobalCoverage()

	return nil
}

func (r *Report) computeGlobalCoverage() {
	// Loop on each package and determine coverage and TLOC by package
	var gcov float64
	for _, pkg := range r.Packages {
		gcov += pkg.Coverage
	}

	// Report the global # of tested LOCs
	if gcov > 0 {
		r.Coverage = gcov / float64(len(r.Packages))
	}
}

// collectPackages collects ALL packages
func (r *Report) collectPackages() error {
	set := token.NewFileSet()
	dirs, err := packageList("Dir")
	if err != nil {
		return err
	}

	var errs []string
	for _, dir := range dirs {
		pkgs, err := parser.ParseDir(set, dir, nil, 0)
		if err != nil {
			err := fmt.Sprintf("Directory %s returned error: `%s`", dir, err.Error())
			log.Error(err)
			errs = append(errs, err)
		}
		for _, pkg := range pkgs {
			// Ignore test package
			if strings.HasSuffix(pkg.Name, "_test") {
				log.Debugf("Ignoring test package `%s`", pkg.Name)
				continue
			}
			// Craft package path
			path := strings.Replace(dir, os.Getenv("GOPATH")+"/src/", "", 1)

			log.Debugf("path %v / package %v", path, pkg.Name)
			r.addPackage(&Package{
				Name: pkg.Name,
				Path: path,
			})
		}
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	return nil
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

// packageList returns a list of Go-like files or directories from PWD,
func packageList(arg string) ([]string, error) {
	cmd, err := exec.Command("sh", "-c", `go list -f '{{.`+arg+`}}' ./... | grep -v vendor | grep -v Godeps`).CombinedOutput()
	if err != nil {
		return nil, err
	}

	pl := strings.Split(strings.TrimSpace(string(cmd)), "\n")
	log.Debug(pl)

	return pl, nil
}
