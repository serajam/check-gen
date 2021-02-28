package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"
)

// generate analyzes all files for structs and generate validation.go for every found packageCheck with tagged struct
func generate(copyrights string, collector *collector, templateBuilder templateBuilder) {
	for _, pkg := range collector.Packages {
		fileTemplate := fileTpl{
			Copyright: copyrights,

			StandardImports: []string{
				"fmt",
				"unicode",
				"unicode/utf8",
			},

			CustomImports: []string{
				"github.com/google/uuid",
			},
		}
		for _, f := range pkg.Files {
			for _, d := range f.Decls {
				g, ok := d.(*ast.GenDecl)
				if !ok {
					continue
				}

				structs := structSearch(g)
				if len(structs) == 0 {
					continue
				}

				for _, s := range structs {
					atLeastOneField := false

					for _, field := range s.Type.Fields.List {

						pos := collector.FileSet.Position(field.Type.Pos())
						typ := collector.Info.TypeOf(field.Type)

						composedType := ""
						baseName := getType(typ, &composedType)
						fmt.Println("Add validation: ", pos, ": ", baseName, "/", composedType)

						if err := templateBuilder.generateCheck(field, s.Name, baseName, composedType); err != nil {
							fmt.Printf("struct %s: %s\n", s.Name, err)
							continue
						}

						atLeastOneField = true
					}

					if !atLeastOneField {
						continue
					}

					err := templateBuilder.generateMethod(s.Name)
					if err != nil {
						fmt.Printf("struct gen %s: %s\n", s.Name, err)
						continue
					}
				}
			}
		}

		fileTemplate.Package = pkg.Name
		err := templateBuilder.generateFile(pkg.Path, fileTemplate)
		if err != nil {
			fmt.Println("Generation error", err)
		}
	}
}

func getType(p types.Type, elType *string) string {
	switch t := p.(type) {
	case *types.Pointer:
		*elType += "pointer "
		return getType(t.Elem().Underlying(), elType)
	case *types.Struct:
		*elType += "struct "
		return "struct"
	case *types.Array:
		*elType += "array "
		return getType(t.Underlying(), elType)
	case *types.Slice:
		*elType += "slice "
		return getType(t.Elem().Underlying(), elType)
	case *types.Map:
		*elType += "map "
		return getType(t.Elem().Underlying(), elType)
	case *types.Basic:
		*elType += t.String() + " "
		return t.String()
	case *types.Named:
		return getType(t.Underlying(), elType)
	default:
		return "not supported"
	}
}

// structSearch searches for structs with tag comment and returns a map name->ast.StructType
func structSearch(g *ast.GenDecl) []structure {
	var structs []structure
	for _, spec := range g.Specs {
		currType, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		currStruct, ok := currType.Type.(*ast.StructType)
		if !ok {
			continue
		}

		for _, item := range currStruct.Fields.List {
			if item.Tag == nil || !strings.Contains(item.Tag.Value, tagPrefix+":") {
				continue
			}

			fmt.Println("found struct: " + currType.Name.String())
			structs = append(
				structs, structure{
					Name: currType.Name.String(),
					Type: currStruct,
				},
			)
			break
		}
	}

	return structs
}
