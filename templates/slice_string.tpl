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
	{{- if and .Deep (or .DeepChecks.CheckIsDigit .DeepChecks.UUID .DeepChecks.CheckIsWord .DeepChecks.CheckMin .DeepChecks.CheckMax .DeepChecks.CheckLen) }}
	for i, v := range {{.StructRef}}.{{.FieldName}} {
	{{- if .DeepChecks.IsRef -}}
		{{- if .DeepChecks.CheckMin }}
		if utf8.RuneCountInString(*v) < {{.DeepChecks.Min}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::min_length_is::{{.DeepChecks.Min}}", i))
		}
		{{- end -}}
		{{- if .DeepChecks.CheckMax }}
		if utf8.RuneCountInString(*v) > {{.DeepChecks.Max}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::max_length_is::{{.DeepChecks.Max}}", i))
		}
		{{- end -}}
		{{- if .DeepChecks.CheckLen }}
		if utf8.RuneCountInString(*v) != {{.DeepChecks.Len}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::length_is_not::{{.DeepChecks.Len}}", i))
		}
		{{- end -}}
		{{- if .DeepChecks.CheckIsDigit }}
		if v != nil {
			for _, n := range *v {
				if !unicode.IsDigit(n) {
					errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::is_not_digit", i))
					break
				}
			}
		}
		{{- end -}}
		{{- if .DeepChecks.CheckIsWord }}
		if v != nil {
			for _, n := range *v {
				if !unicode.IsLetter(n) && !unicode.IsDigit(n) && !unicode.IsPunct(n) && !unicode.IsSpace(n) {
					errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::is_not_word", i))
					break
				}
			}
		}
		{{- end -}}
		{{- if .DeepChecks.UUID }}
		if v != nil {
			if _, err := uuid.Parse(*v); err != nil {
				errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::is_not_uuid", i))
			}
		}
		{{- end -}}
	{{- else -}}
		{{- if .DeepChecks.CheckMin }}
		if utf8.RuneCountInString(v) < {{.DeepChecks.Min}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::min_length_is::{{.DeepChecks.Min}}", i))
		}
		{{- end -}}
		{{- if .DeepChecks.CheckMax }}
		if utf8.RuneCountInString(v) > {{.DeepChecks.Max}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::max_length_is::{{.DeepChecks.Max}}", i))
		}
		{{- end -}}
		{{- if .DeepChecks.CheckLen }}
		if utf8.RuneCountInString(v) != {{.DeepChecks.Len}} {
			errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::length_is_not::{{.DeepChecks.Len}}", i))
		}
		{{- end -}}
		{{- if .DeepChecks.CheckIsDigit }}
		for _, n := range v {
			if !unicode.IsDigit(n) {
				errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::is_not_digit", i))
				break
			}
		}
		{{- end -}}
		{{- if .DeepChecks.CheckIsWord }}
		for _, n := range v {
			if !unicode.IsLetter(n) && !unicode.IsDigit(n) && !unicode.IsPunct(n) && !unicode.IsSpace(n) {
				errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::is_not_word", i))
				break
			}
		}
		{{- end -}}
		{{- if .DeepChecks.UUID }}
		if v != "" {
			if _, err := uuid.Parse(v); err != nil {
				errs = append(errs, fmt.Sprintf("{{.ErrorFieldName}}_%v::is_not_uuid", i))
			}
		}
		{{- end -}}
	{{- end }}
	}
	{{- end -}}
