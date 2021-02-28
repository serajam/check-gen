package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"os"
	"path/filepath"

	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/pflag"
)

const version = "1.0.0"

func main() {
	copyright := pflag.StringP("copyright", "c", "", "copyright for generated files")
	versionFlag := pflag.BoolP("version", "v", false, "print version")

	pflag.Parse()
	if *versionFlag {
		fmt.Printf("%s\n", version)
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		fmt.Println("Please specify path.to look for types")
		fmt.Println("For example: \ncheck-gen pkg/api/v1/models")
		os.Exit(0)
	}

	path := os.Args[1]
	fmt.Println(path)
	if path == "" {
		fmt.Println("Please specify path.to look for types")
		fmt.Println("For example: \ncheck-gen pkg/api/v1/models")
		os.Exit(0)
	}

	// includes templates in binary
	box := packr.New("templates", "./templates")

	templateBuilder := newTemplateBuilder(newTagReader())
	err := templateBuilder.load(box)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	collector := &collector{
		FileSet:  token.NewFileSet(),
		Info:     &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)},
		Packages: []packageCheck{},
	}

	err = collector.collect(filepath.Clean(path))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(collector.Packages) == 0 {
		fmt.Println("no go files found")
		os.Exit(0)
	}

	generate(*copyright, collector, templateBuilder)
}
