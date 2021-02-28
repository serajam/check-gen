{{if .FieldChecks.IsRef -}}
	{{if .FieldChecks.Required}}
	if {{.StructRef}}.{{.FieldName}} == nil {
		errs = append(errs, "{{.ErrorFieldName}}::is_required")
	}
	{{- end}}
{{- end}}