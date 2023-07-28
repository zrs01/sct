{{- range .Entity }}
// -------------------------------- {{ .Name }} -------------------------------
new FormGroup({
  {{- range .Members }}
  {{ lower(.Name[:1]) + .Name[1:] }}: new FormControl<{{ typescriptType(.DataType, .IsCollection) }} | null>(null),
  {{- end }}
});
{{ end }}