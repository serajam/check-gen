{{- if .FieldChecks.IsRef -}}
	{{- if .FieldChecks.Required }}
	if {{.StructRef}}.{{.FieldName}} == nil {
		errs = append(errs, "{{.ErrorFieldName}}::is_required")
	}
	{{- end -}}
	{{- if .CheckNested }}
	if {{.StructRef}}.{{.FieldName}} != nil {
		if e := {{.CallStruct}}.Validate(); len(e) > 0 {
			errs = append(errs, e...)
		}
	}
	{{- end -}}
{{- else -}}
	{{- if .CheckNested }}
	if e := {{.CallStruct}}.Validate(); len(e) > 0 {
		errs = append(errs, e...)
	}
	{{- end -}}
{{- end -}}