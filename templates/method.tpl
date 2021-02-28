
// Validate validates struct accordingly to fields tags
func ({{.StructRef}} {{.StructName}}) Validate() []string {
	var errs []string
{{- .Body }}

	return errs
}