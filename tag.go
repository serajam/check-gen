package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// type of tags values
const (
	tagPrefix = "check" // basic prefix
	required  = "required"
	deep      = "deep"  // for slices
	check     = "check" // for nested structs
	flen      = "len"
	max       = "max"
	min       = "min"
	digit     = "digit"
	word      = "word"
	uuid      = "uuid"
	password  = "password"
	phone     = "phone"
)

// tagReader holds logic for analyzing field tags
type tagReader struct {
	tagRegexp *regexp.Regexp
}

// newTagReader creates new tag reader
func newTagReader() tagReader {
	return tagReader{tagRegexp: regexp.MustCompile(tagPrefix + `:".*"`)}
}

// readAndSet reads tags and sets specific field or slice value validation
func (t tagReader) readAndSet(tag string, tpl *checkTpl) error {
	validTag := t.tagRegexp.FindString(tag)
	tagParts := strings.Split(validTag, ":")
	if len(tagParts) == 1 {
		return fmt.Errorf("incorrect tag format")
	}

	checks := strings.Split(tagParts[1], ",")

	checkField := fieldCheck{}
	checkDeep := fieldCheck{}

	for _, c := range checks {
		var value interface{}
		checkType := strings.Trim(c, `"`)

		if strings.Contains(checkType, "=") {
			parts := strings.Split(checkType, "=")
			if len(parts) == 1 {
				return errors.New("incorrect tag format")
			}

			checkType = parts[0]
			value = parts[1]
		}

		if tpl.Deep {
			setValidation(checkType, value, &checkDeep, tpl)
		} else {
			setValidation(checkType, value, &checkField, tpl)
		}
	}

	tpl.FieldChecks = checkField
	tpl.DeepChecks = checkDeep

	return nil
}

func setValidation(checkType string, value interface{}, fieldCheck *fieldCheck, tpl *checkTpl) {
	switch checkType {
	case required:
		fieldCheck.Required = true
	case check:
		tpl.CheckNested = true
	case flen:
		lenFunc(fieldCheck, value)
	case min:
		minFunc(fieldCheck, value)
	case max:
		maxFunc(fieldCheck, value)
	case deep:
		tpl.Deep = true
	case digit:
		fieldCheck.CheckIsDigit = true
	case word:
		fieldCheck.CheckIsWord = true
	case uuid:
		fieldCheck.UUID = true
	case password:
		fieldCheck.Password = true
	case phone:
		phoneFunc(fieldCheck, value)
	}
}

func minFunc(fieldCheck *fieldCheck, value interface{}) {
	valStr, ok := value.(string)
	if !ok {
		fmt.Printf("failed to cast min value: %v\n", value)
		return
	}

	res, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		fmt.Printf("failed to parse min value: %v\n", valStr)
		return
	}

	fieldCheck.CheckMin = true
	fieldCheck.Min = res
}

func maxFunc(fieldCheck *fieldCheck, value interface{}) {
	valStr, ok := value.(string)
	if !ok {
		fmt.Printf("failed to cast max value: %v\n", value)
		return
	}

	res, err := strconv.ParseFloat(valStr, 64)
	if err != nil {
		fmt.Printf("failed to parse max value: %v\n", valStr)
		return
	}

	fieldCheck.CheckMax = true
	fieldCheck.Max = res
}

func phoneFunc(fieldCheck *fieldCheck, value interface{}) {
	val, ok := value.(string)
	if !ok {
		fmt.Printf("failed to cast phone value: %v\n", value)
		return
	}

	split := strings.Split(val, "~")

	if len(split) < 2 {
		fmt.Printf("invalid phone values: %v\n", value)
		return
	}

	fieldCheck.Phone = split
}

func lenFunc(fieldCheck *fieldCheck, value interface{}) {
	valStr, ok := value.(string)
	if !ok {
		fmt.Printf("failed to cast len value: %v\n", value)
		return
	}

	res, err := strconv.Atoi(valStr)
	if err != nil {
		fmt.Printf("failed to parse len value: %v\n", valStr)
		return
	}

	fieldCheck.CheckLen = true
	fieldCheck.Len = res
}
