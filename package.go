// Copyright (c) 2016, Hotolab. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cov

import "fmt"

// Package describes a package inner characteristics
type Package struct {
	// Name is the canonical path of the package.
	Name string
	// Coverage
	Coverage float64
	// TLOC is the number of lines of code which are tested (LOC)
	TLOC int64
	// Functions is a list of functions registered with this package.
	Functions []*Function
}

// Accumulate will accumulate the coverage information from the provided
// Package into this Package.
func (p *Package) Accumulate(p2 *Package) error {
	if p.Name != p2.Name {
		return fmt.Errorf("Names do not match: %q != %q", p.Name, p2.Name)
	}
	if len(p.Functions) != len(p2.Functions) {
		return fmt.Errorf("Function counts do not match: %d != %d", len(p.Functions), len(p2.Functions))
	}
	for i, f := range p.Functions {
		err := f.Accumulate(p2.Functions[i])
		if err != nil {
			return err
		}
	}

	return nil
}
