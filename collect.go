package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/types"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// collect recursively searches for go files in a specified dir
// collects file set (tokens) information, types information
func (c *collector) collect(path string) error {
	// importer is needed to import custom (Named) types from other packages
	cfg := &types.Config{
		IgnoreFuncBodies: true,
		Error:            func(err error) {},
		Importer:         importer.ForCompiler(c.FileSet, "source", nil),
	}

	paths, err := getPaths(path)
	if err != nil {
		return err
	}

	// collect information
	for _, p := range paths {
		pkgs, err := parser.ParseDir(
			c.FileSet, p, func(fileInfo os.FileInfo) bool {
				// exclude exists validate.go files
				return !strings.Contains(fileInfo.Name(), "validate.go") &&
						!strings.Contains(fileInfo.Name(), "validate_test.go")
			}, parser.AllErrors|parser.ParseComments,
		)
		if err != nil {
			return fmt.Errorf("Collect.ParseDir: %s", err)
		}

		for pkgName, pkg := range pkgs {
			files := make([]*ast.File, 0, len(pkg.Files))
			for _, ff := range pkg.Files {
				files = append(files, ff)
			}

			_, err := cfg.Check(p, c.FileSet, files, c.Info)
			if err != nil {
				return fmt.Errorf("Collect.Check: %s", err)
			}

			// sort for same position files in validate
			sort.Slice(
				files, func(i, j int) bool {
					return files[i].Pos() < files[j].Pos()
				},
			)

			c.Packages = append(
				c.Packages, packageCheck{
					Name:  pkgName,
					Path:  p,
					Files: files,
				},
			)
		}
	}

	return nil
}

func getPaths(path string) ([]string, error) {
	var dirs []string
	err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				dirs = append(dirs, path)
			}
			return nil
		},
	)
	if err != nil {
		return nil, err
	}

	return dirs, nil
}
