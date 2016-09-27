// Copyright (c) 2016, Hotolab. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cov

// Package describes a package inner characteristics
type Package struct {
	// Name is the package name
	Name string `json:"name"`
	// Path is the canonical path of the package.
	Path string `json:"path"`
	// Coverage
	Coverage float64 `json:"coverage"`
	// Functions is a list of functions registered with this package.
	Functions []*Function `json:"-"`
}
