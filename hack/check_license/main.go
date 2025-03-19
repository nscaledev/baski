/*
Copyright 2025 Nscale.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

const (
	// goApache2LicenseHeader is an exact match for a license header.
	goApache2LicenseHeader = `/*
Copyright 2025 Nscale.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/`
)

var (
	// errFail tells you there was an error detected.
	errFail = errors.New("errors detected")

	// errNoComments tells you that you've not commented anything.
	errNoComments = errors.New("file contains no comments")

	// errFirstCommentNotLicense tells you that the first comment isn't a license.
	errFirstCommentNotLicense = errors.New("first comment not a valid license")
)

// glob does a recursive walk of the working directory, returning all files that
// match the provided extension e.g. ".go".
func glob(extension string) ([]string, error) {
	var files []string

	appendFileWithExtension := func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}

		if strings.Contains(path, "pkg/mock") {
			return nil
		}

		if filepath.Ext(path) != extension {
			return nil
		}

		files = append(files, path)

		return nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	if err := filepath.Walk(wd, appendFileWithExtension); err != nil {
		return nil, err
	}

	return files, nil
}

// checkGoLicenseInComments checks the AST for a license header anywhere the top
// level (because autogenerated code does what it likes).
func checkGoLicenseInComments(path string, file *ast.File) error {
	if len(file.Comments) == 0 {
		return fmt.Errorf("%s: %w\n", path, errNoComments)
	}

	for _, comment := range file.Comments {
		for _, item := range comment.List {
			if item.Text == goApache2LicenseHeader {
				return nil
			}
		}
	}

	return fmt.Errorf("%s: %w\n", path, errFirstCommentNotLicense)
}

// checkGoLicenseFile parses a source file and checks there is a license header in there.
func checkGoLicenseFile(path string) error {
	fset := token.NewFileSet()

	file, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	if err := checkGoLicenseInComments(path, file); err != nil {
		return err
	}

	return nil
}

// checkGoLicense finds all go source files in the working directory, then parses them
// into an AST and checks there is a license header in there.
func checkGoLicense() error {
	paths, err := glob(".go")
	if err != nil {
		return err
	}

	var hasErrors bool

	for _, path := range paths {
		if strings.Contains(path, "generated") {
			continue
		}
		if err := checkGoLicenseFile(path); err != nil {
			fmt.Println(err)

			hasErrors = true
		}
	}

	if hasErrors {
		return errFail
	}

	return nil
}

// main runs any license checkers over the code.
func main() {
	if err := checkGoLicense(); err != nil {
		os.Exit(1)
	}
}
