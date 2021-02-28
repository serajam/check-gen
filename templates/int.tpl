{{- if .FieldChecks.IsRef -}}
	{{- if .FieldChecks.Required }}
	if {{.StructRef}}.{{.FieldName}} == nil || *{{.StructRef}}.{{.FieldName}} == 0 {
		errs = append(errs, "{{.ErrorFieldName}}::is_required")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckMin }}
	if {{.StructRef}}.{{.FieldName}} != nil && *{{.StructRef}}.{{.FieldName}} < {{.FieldChecks.Min}} {
		errs = append(errs, "{{.ErrorFieldName}}::min_value_is::{{.FieldChecks.Min}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckMax }}
	if {{.StructRef}}.{{.FieldName}} != nil && *{{.StructRef}}.{{.FieldName}} > {{.FieldChecks.Max}} {
		errs = append(errs, "{{.ErrorFieldName}}::max_value_is::{{.FieldChecks.Max}}")
	}
	{{- end -}}
{{- else -}}
	{{- if .FieldChecks.Required}}
	if {{.StructRef}}.{{.FieldName}} == 0 {
		errs = append(errs, "{{.ErrorFieldName}}::is_required")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckMin }}
	if {{.StructRef}}.{{.FieldName}} < {{.FieldChecks.Min}} {
		errs = append(errs, "{{.ErrorFieldName}}::min_value_is::{{.FieldChecks.Min}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckMax }}
	if {{.StructRef}}.{{.FieldName}} > {{.FieldChecks.Max}} {
		errs = append(errs, "{{.ErrorFieldName}}::max_value_is::{{.FieldChecks.Max}}")
	}
	{{- end -}}
{{- end -}}
