// Code generated by check-gen; DO NOT EDIT.
// Package custom contains models and autogenerated validation code
package custom

// Validate validates struct accordingly to fields tags
func (s StructInCustom) Validate() []string {
	var errs []string
	if s.Name == "" {
		errs = append(errs, "name::is_required")
	}

	return errs
}
