// Code generated by check-gen; DO NOT EDIT.
// Package innerpkg contains models and autogenerated validation code
package innerpkg

import (
	"unicode/utf8"
)

// Validate validates struct accordingly to fields tags
func (f Foo) Validate() []string {
	var errs []string
	if utf8.RuneCountInString(f.Name) < 1 {
		errs = append(errs, "name::min_length_is::1")
	}
	if utf8.RuneCountInString(f.Name) > 1000 {
		errs = append(errs, "name::max_length_is::1000")
	}
	if e := f.Bar.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (f FooRef) Validate() []string {
	var errs []string
	if utf8.RuneCountInString(f.Name) < 1 {
		errs = append(errs, "name::min_length_is::1")
	}
	if utf8.RuneCountInString(f.Name) > 1000 {
		errs = append(errs, "name::max_length_is::1000")
	}
	if f.Bar != nil {
		if e := f.Bar.Validate(); len(e) > 0 {
			errs = append(errs, e...)
		}
	}

	return errs
}
