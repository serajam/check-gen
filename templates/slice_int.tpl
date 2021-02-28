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

	{{- if and .Deep (or .DeepChecks.CheckMin .DeepChecks.CheckMax) }}
	for i, v := range {{.StructRef}}.{{.FieldName}} {
	{{- if .DeepChecks.IsRef -}}
		{{- if .DeepChecks.CheckMin }}
		if *v < {{.DeepChecks.Min}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%d::min_value_is::{{.DeepChecks.Min}}", i))
		}
		{{- end -}}
		{{- if .DeepChecks.CheckMax }}
		if *v > {{.DeepChecks.Max}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%d::max_value_is::{{.DeepChecks.Max}}", i))
		}
		{{- end -}}
	{{- else -}}
		{{- if .DeepChecks.CheckMin }}
		if v < {{.DeepChecks.Min}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%d::min_value_is::{{.DeepChecks.Min}}", i))
		}
		{{- end -}}
		{{- if .DeepChecks.CheckMax }}
		if v > {{.DeepChecks.Max}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%d::max_value_is::{{.DeepChecks.Max}}", i))
		}
		{{- end -}}
	{{- end }}
	}
	{{- end -}}
