{{- range .Ts }}
// -------------------------------- {{ .Name }} -------------------------------
new FormGroup({
  {{- range .Members }}
  {{ lower(.Name[:1]) + .Name[1:] }}: new FormControl<{{ raw(.DataType) }}{{ .Value == "null" ? " | null" : "" }}>({{ raw(.Value) }}{{ .Value != "null" ? ", { nonNullable: true }": "" }}),
  {{- end }}
});
{{ end }}