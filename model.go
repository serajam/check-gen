package main

import (
	"go/ast"
	"go/token"
	"go/types"
	"strings"
)

type (
	collector struct {
		Info     *types.Info
		FileSet  *token.FileSet
		Packages []packageCheck
	}

	packageCheck struct {
		Name  string
		Path  string
		Files []*ast.File
	}

	structure struct {
		Name string
		Type *ast.StructType
	}

	fileTpl struct {
		Body,
		Package,
		Copyright string

		StandardImports []string
		CustomImports   []string
	}

	checkTpl struct {
		StructRef      string
		FieldName      string
		ErrorFieldName string
		CallStruct     string

		CheckNested bool
		Deep        bool
		DeepChecks  fieldCheck
		FieldChecks fieldCheck
	}

	funcTpl struct {
		StructRef  string
		StructName string
		Body       string
	}

	fieldCheck struct {
		IsRef          bool
		Required       bool
		RequiredEither []string
		CheckMax       bool
		Max            float64
		CheckMin       bool
		Min            float64
		CheckLen       bool
		Len            int
		CheckIsDigit   bool
		CheckIsWord    bool
		UUID           bool
		Password       bool
		Phone          []string
	}
)

// setBody and remove unused imports
func (r *fileTpl) setBody(body string) {
	var filtered []string
	for _, pkgName := range r.StandardImports {
		pkgWords := strings.Split(pkgName, "/")
		if strings.Contains(body, pkgWords[len(pkgWords)-1]) {
			filtered = append(filtered, pkgName)
		}
	}
	r.StandardImports = filtered
	filtered = nil

	for _, pkgName := range r.CustomImports {
		pkgWords := strings.Split(pkgName, "/")
		if strings.Contains(body, pkgWords[len(pkgWords)-1]) {
			filtered = append(filtered, pkgName)
		}
	}

	r.CustomImports = filtered
	r.Body = body
}
