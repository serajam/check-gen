package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"os"
	"regexp"
	"strings"
	"text/template"

	"github.com/gobuffalo/packr/v2"
)

// templateBuilder builds validation methods for structs
type templateBuilder struct {
	templates    *template.Template
	templateType map[string]string

	methodBuff   *bytes.Buffer
	methodWriter *bufio.Writer

	fileBuff   *bytes.Buffer
	fileWriter *bufio.Writer

	fieldRegexp *regexp.Regexp

	tagReader tagReader
}

// newTemplateBuilder creates new builder
func newTemplateBuilder(reader tagReader) templateBuilder {
	t := templateBuilder{tagReader: reader}
	t.templateType = map[string]string{
		"int":           "int.tpl",
		"int64":         "int.tpl",
		"int32":         "int.tpl",
		"int16":         "int.tpl",
		"int8":          "int.tpl",
		"float32":       "int.tpl",
		"float64":       "int.tpl",
		"string":        "string.tpl",
		"bool":          "bool.tpl",
		"slice_struct":  "slice_struct.tpl",
		"slice_string":  "slice_string.tpl",
		"slice_byte":    "slice_byte.tpl",
		"slice_int":     "slice_int.tpl",
		"slice_int64":   "slice_int.tpl",
		"slice_int32":   "slice_int.tpl",
		"slice_int16":   "slice_int.tpl",
		"slice_int8":    "slice_int.tpl",
		"slice_float32": "slice_int.tpl",
		"slice_float64": "slice_int.tpl",

		// behavior is the same as that of a slice
		"map_struct":  "slice_struct.tpl",
		"map_string":  "slice_string.tpl",
		"map_byte":    "slice_byte.tpl",
		"map_int":     "slice_int.tpl",
		"map_int64":   "slice_int.tpl",
		"map_int32":   "slice_int.tpl",
		"map_int16":   "slice_int.tpl",
		"map_int8":    "slice_int.tpl",
		"map_float32": "slice_int.tpl",
		"map_float64": "slice_int.tpl",

		"struct": "struct.tpl",
	}

	t.methodBuff = &bytes.Buffer{}
	t.fileBuff = &bytes.Buffer{}
	t.methodWriter = bufio.NewWriter(t.methodBuff)
	t.fileWriter = bufio.NewWriter(t.fileBuff)
	t.fieldRegexp = regexp.MustCompile("([a-z0-9])([A-Z])")

	return t
}

// generateFile creates new file in package with validation methods for all structs
func (tb *templateBuilder) generateFile(dir string, ftpl fileTpl) error {
	err := tb.fileWriter.Flush()
	if err != nil {
		return err
	}

	if tb.fileBuff.Len() == 0 {
		fmt.Println("Nothing to generate")
		return nil
	}

	out, err := os.Create(dir + "/validate.go")
	if err != nil {
		return err
	}
	defer func() {
		err := out.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	ftpl.setBody(tb.fileBuff.String())

	err = tb.templates.ExecuteTemplate(out, "file.tpl", ftpl)
	tb.fileBuff.Reset()
	if err != nil {
		return err
	}

	return nil
}

// generateMethod creates validation method based on previously built fieldCheck for fields
func (tb *templateBuilder) generateMethod(name string) error {
	err := tb.methodWriter.Flush()
	if err != nil {
		return err
	}

	err = tb.templates.ExecuteTemplate(
		tb.fileWriter, "method.tpl", funcTpl{
			StructName: name,
			StructRef:  strings.ToLower(string(name[0])),
			Body:       tb.methodBuff.String(),
		},
	)
	tb.methodBuff.Reset()
	if err != nil {
		return err
	}

	_, err = tb.fileWriter.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}

// generateCheck creates fieldCheck for field
func (tb *templateBuilder) generateCheck(field *ast.Field, structName string, fieldType string, composedType string) error {
	var err error
	fieldName := ""

	if len(field.Names) != 0 {
		fieldName = field.Names[0].Name
	} else if v, ok := field.Type.(*ast.Ident); ok {
		fieldName = v.Name
	} else if v, ok := field.Type.(*ast.SelectorExpr); ok {
		fieldName = v.Sel.Name
	} else if v, ok := field.Type.(*ast.StarExpr); ok {
		switch f := v.X.(type) {
		case *ast.Ident:
			fieldName = f.Name
		case *ast.SelectorExpr:
			fieldName = f.Sel.Name
		default:
			return fmt.Errorf("skipping field %+v: could not determine field name", field)
		}
	}

	if field.Tag == nil || !strings.Contains(field.Tag.Value, tagPrefix+":") {
		return fmt.Errorf("skipping field %s: no `"+tagPrefix+"` tag found", fieldName)
	}

	tplData := checkTpl{
		StructRef:      strings.ToLower(string(structName[0])),
		FieldName:      fieldName,
		ErrorFieldName: strings.ToLower(tb.fieldRegexp.ReplaceAllString(fieldName, "${1}_${2}")),
	}

	err = tb.tagReader.readAndSet(field.Tag.Value, &tplData)
	if err != nil {
		return fmt.Errorf("field %s: %s", fieldName, err)
	}

	tplType := fieldType
	if strings.Contains(composedType, "slice") {
		tplType = "slice_" + tplType
	}

	if strings.Contains(composedType, "map") {
		tplType = "map_" + tplType
	}

	tplName, isTypeSupported := tb.templateType[tplType]
	if fieldType == "struct" {
		tplData.CallStruct = tplData.StructRef + "." + fieldName
	}

	if strings.Contains(composedType, "pointer") {
		tplData.DeepChecks.IsRef = true
		tplData.FieldChecks.IsRef = true
	}

	if !isTypeSupported {
		return fmt.Errorf("type of field '%s' not supported", tplType)
	}

	err = tb.templates.ExecuteTemplate(tb.methodWriter, tplName, tplData)
	if err != nil {
		return err
	}

	return nil
}

// load loads templates
func (tb *templateBuilder) load(box *packr.Box) error {
	var t *template.Template
	var err error

	for _, n := range box.List() {
		if t == nil {
			s, _ := box.FindString(n)
			t, err = template.New(n).Parse(s)
			if err != nil {
				return err
			}
		}

		s, _ := box.FindString(n)
		t, err = t.New(n).Parse(s)
		if err != nil {
			return err
		}
	}
	tb.templates = t

	return nil
}
