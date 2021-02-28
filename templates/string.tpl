{{- if .FieldChecks.IsRef -}}
	{{- if .FieldChecks.Required }}
	if {{.StructRef}}.{{.FieldName}} == nil || len(*{{.StructRef}}.{{.FieldName}}) == 0 {
		errs = append(errs, "{{.ErrorFieldName}}::is_required")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckMin }}
	if {{.StructRef}}.{{.FieldName}} != nil && utf8.RuneCountInString(*{{.StructRef}}.{{.FieldName}}) < {{.FieldChecks.Min}} {
		errs = append(errs, "{{.ErrorFieldName}}::min_length_is::{{.FieldChecks.Min}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckMax }}
	if {{.StructRef}}.{{.FieldName}} != nil && utf8.RuneCountInString(*{{.StructRef}}.{{.FieldName}}) > {{.FieldChecks.Max}} {
		errs = append(errs, "{{.ErrorFieldName}}::max_length_is::{{.FieldChecks.Max}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckLen }}
	if {{.StructRef}}.{{.FieldName}} != nil && utf8.RuneCountInString(*{{.StructRef}}.{{.FieldName}}) != {{.FieldChecks.Len}} {
		errs = append(errs, "{{.ErrorFieldName}}::length_is_not::{{.FieldChecks.Len}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckIsDigit }}
	if {{.StructRef}}.{{.FieldName}} != nil {
		for _, n := range *{{.StructRef}}.{{.FieldName}} {
			if !unicode.IsDigit(n) {
				errs = append(errs, "{{.ErrorFieldName}}::is_not_digit")
				break
			}
		}
	}
	{{- end -}}
	{{- if .FieldChecks.CheckIsWord }}
	if {{.StructRef}}.{{.FieldName}} != nil {
		for _, n := range *{{.StructRef}}.{{.FieldName}} {
			if !unicode.IsLetter(n) && !unicode.IsDigit(n) && !unicode.IsPunct(n) && !unicode.IsSpace(n) {
				errs = append(errs, "{{.ErrorFieldName}}::is_not_word")
				break
			}
		}
	}
	{{- end -}}
	{{- if .FieldChecks.UUID }}
	if {{.StructRef}}.{{.FieldName}} != nil {
		if _, err := uuid.Parse(*{{.StructRef}}.{{.FieldName}}); err != nil {
			errs = append(errs, "{{.ErrorFieldName}}::is_not_uuid")
		}
	}
	{{- end -}}
	{{- if .FieldChecks.Password }}
	if {{.StructRef}}.{{.FieldName}} != nil && len(*{{.StructRef}}.{{.FieldName}}) > 0 {
		var {{.FieldName}}Upper, {{.FieldName}}Letter, {{.FieldName}}Number, {{.FieldName}}Spaces, {{.FieldName}}InvalidChar bool
		for _, c := range *{{.StructRef}}.{{.FieldName}} {
			if c < 33 || c > 126 {
				{{.FieldName}}InvalidChar = true
				break
			}
			switch {
			case unicode.IsNumber(c):
				{{.FieldName}}Number = true
			case unicode.IsUpper(c):
				{{.FieldName}}Upper = true
			case unicode.IsLetter(c):
				{{.FieldName}}Letter = true
			case c == ' ':
				{{.FieldName}}Spaces = true
				break
			}
		}
		if !{{.FieldName}}Upper || !{{.FieldName}}Letter || !{{.FieldName}}Number || {{.FieldName}}Spaces || {{.FieldName}}InvalidChar {
			errs = append(errs, "{{.ErrorFieldName}}::invalid_password_format")
		}
	}
	{{- end -}}
	{{- if len .FieldChecks.Phone}}
	if {{.StructRef}}.{{.FieldName}} != nil && len(*{{.StructRef}}.{{.FieldName}}) > 0 {
		if len(*{{.StructRef}}.{{.FieldName}}) == 1 {
			errs = append(errs, "{{.ErrorFieldName}}::invalid_phone_format::wrong_length")
		} else {
			prefix, numbers := (*{{.StructRef}}.{{.FieldName}})[0], (*{{.StructRef}}.{{.FieldName}})[1:]

			if prefix != 43 {
				errs = append(errs, "{{.ErrorFieldName}}::invalid_phone_format::should_be_started_with_'+'")
			}

			if len(numbers) < {{ index .FieldChecks.Phone 0 }} || len(numbers) > {{ index .FieldChecks.Phone 1 }} {
				errs = append(errs, "{{.ErrorFieldName}}::invalid_phone_format::wrong_length")
			}

			for _, c := range numbers {
				if !unicode.IsDigit(c) {
					errs = append(errs, "{{.ErrorFieldName}}::invalid_phone_format::not_a_digits")
					break
				}
			}
		}
	}
    {{- end -}}
{{- else -}}
	{{- if .FieldChecks.Required}}
	if {{.StructRef}}.{{.FieldName}} == "" {
		errs = append(errs, "{{.ErrorFieldName}}::is_required")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckMin }}
	if utf8.RuneCountInString({{.StructRef}}.{{.FieldName}}) < {{.FieldChecks.Min}} {
		errs = append(errs, "{{.ErrorFieldName}}::min_length_is::{{.FieldChecks.Min}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckMax }}
	if utf8.RuneCountInString({{.StructRef}}.{{.FieldName}}) > {{.FieldChecks.Max}} {
		errs = append(errs, "{{.ErrorFieldName}}::max_length_is::{{.FieldChecks.Max}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckLen }}
	if utf8.RuneCountInString({{.StructRef}}.{{.FieldName}}) != {{.FieldChecks.Len}} {
		errs = append(errs, "{{.ErrorFieldName}}::length_is_not::{{.FieldChecks.Len}}")
	}
	{{- end -}}
	{{- if .FieldChecks.CheckIsDigit }}
	for _, n := range {{.StructRef}}.{{.FieldName}} {
		if !unicode.IsDigit(n) {
			errs = append(errs, "{{.ErrorFieldName}}::is_not_digit")
			break
		}
	}
	{{- end -}}
	{{- if .FieldChecks.CheckIsWord }}
	for _, n := range {{.StructRef}}.{{.FieldName}} {
		if !unicode.IsLetter(n) && !unicode.IsDigit(n) && !unicode.IsPunct(n) && !unicode.IsSpace(n) {
			errs = append(errs, "{{.ErrorFieldName}}::is_not_word")
			break
		}
	}
	{{- end -}}
	{{- if .FieldChecks.UUID }}
	if {{.StructRef}}.{{.FieldName}} != "" {
		if _, err := uuid.Parse({{.StructRef}}.{{.FieldName}}); err != nil {
			errs = append(errs, "{{.ErrorFieldName}}::is_not_uuid")
		}
	}
	{{- end -}}
	{{- if .FieldChecks.Password }}
	if len({{.StructRef}}.{{.FieldName}}) > 0 {
		var {{.FieldName}}Upper, {{.FieldName}}Letter, {{.FieldName}}Number, {{.FieldName}}Spaces, {{.FieldName}}InvalidChar bool
		for _, c := range {{.StructRef}}.{{.FieldName}} {
			if c < 33 || c > 126 {
				{{.FieldName}}InvalidChar = true
				break
			}
			switch {
			case unicode.IsNumber(c):
				{{.FieldName}}Number = true
			case unicode.IsUpper(c):
				{{.FieldName}}Upper = true
			case unicode.IsLetter(c):
				{{.FieldName}}Letter = true
			case c == ' ':
				{{.FieldName}}Spaces = true
				break
			}
		}
		if !{{.FieldName}}Upper || !{{.FieldName}}Letter || !{{.FieldName}}Number || {{.FieldName}}Spaces || {{.FieldName}}InvalidChar {
			errs = append(errs, "{{.ErrorFieldName}}::invalid_password_format")
		}
	}
	{{- end -}}
	{{- if len .FieldChecks.Phone}}
	if len({{.StructRef}}.{{.FieldName}}) > 0 {
		if len({{.StructRef}}.{{.FieldName}}) == 1 {
			errs = append(errs, "{{.ErrorFieldName}}::invalid_phone_format::wrong_length")
		} else {
			prefix, numbers := {{.StructRef}}.{{.FieldName}}[0], {{.StructRef}}.{{.FieldName}}[1:]

			if prefix != 43 {
				errs = append(errs, "{{.ErrorFieldName}}::invalid_phone_format::should_be_started_with_'+'")
			}

			if len(numbers) < {{ index .FieldChecks.Phone 0 }} || len(numbers) > {{ index .FieldChecks.Phone 1 }} {
				errs = append(errs, "{{.ErrorFieldName}}::invalid_phone_format::wrong_length")
			}

			for _, c := range numbers {
				if !unicode.IsDigit(c) {
					errs = append(errs, "{{.ErrorFieldName}}::invalid_phone_format::not_a_digits")
					break
				}
			}
		}
	}
    {{- end -}}
{{- end -}}
