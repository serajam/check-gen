	{{- if .FieldChecks.Required}}
	if {{.StructRef}}.{{.FieldName}} == nil || len({{.StructRef}}.{{.FieldName}}) == 0 {
		errs = append(errs, "{{.ErrorFieldName}}::is_required")
	}
	{{- end -}}

	{{- if .FieldChecks.CheckMin }}
	if len({{.StructRef}}.{{.FieldName}}) < {{.FieldChecks.Min}} {
		errs = append(errs, "{{.ErrorFieldName}}::min_length_is::{{.FieldChecks.Min}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckMax }}
	if len({{.StructRef}}.{{.FieldName}}) > {{.FieldChecks.Max}} {
		errs = append(errs, "{{.ErrorFieldName}}::max_length_is::{{.FieldChecks.Max}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckLen }}
	if len({{.StructRef}}.{{.FieldName}}) != {{.FieldChecks.Len}} {
		errs = append(errs, "{{.ErrorFieldName}}::length_is_not::{{.FieldChecks.Len}}")
	}
	{{- end -}}

	{{if and .Deep .CheckNested}}
	for _, v := range {{.StructRef}}.{{.FieldName}} {
		if e := v.Validate(); len(e) > 0 {
			errs = append(errs, e...)
		}
	}
	{{- end -}}
