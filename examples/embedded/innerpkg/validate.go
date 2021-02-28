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

	return errs
}

// Validate validates struct accordingly to fields tags
func (b Bar) Validate() []string {
	var errs []string
	if utf8.RuneCountInString(b.Surname) < 1 {
		errs = append(errs, "surname::min_length_is::1")
	}
	if utf8.RuneCountInString(b.Surname) > 1000 {
		errs = append(errs, "surname::max_length_is::1000")
	}
	if e := b.Foo.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}

	return errs
}

// Validate validates struct accordingly to fields tags
func (b BarRef) Validate() []string {
	var errs []string
	if utf8.RuneCountInString(b.Surname) < 1 {
		errs = append(errs, "surname::min_length_is::1")
	}
	if utf8.RuneCountInString(b.Surname) > 1000 {
		errs = append(errs, "surname::max_length_is::1000")
	}
	if b.Foo == nil {
		errs = append(errs, "foo::is_required")
	}
	if b.Foo != nil {
		if e := b.Foo.Validate(); len(e) > 0 {
			errs = append(errs, e...)
		}
	}

	return errs
}
