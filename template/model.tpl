{{- range .Ts }}
export interface {{ .Name }} {
  {{- range .Members }}
  {{ lower(.Name[:1]) + .Name[1:] }}?: {{ .DataType }};
  {{- end }}
}
{{ end }}
