{{- data := .Data }}
{{- range .Cs }}
// -------------------------------- {{ .Name }} -------------------------------
using Cms.Repo.Entity;

namespace {{ data.Namespace }};

public class {{ .Name }}Dto : {{ .Name }} {
    public new byte[]? RowVer { get; set; }
    {{- range .Members }}
      {{- if .Virtual }}
        {{- if len(.CollectionType) == 0 }}
    public new {{ .DataType }}Dto? {{ .Name }} { get; set; }
        {{- else }}
    public new ICollection<{{ .CollectionType }}Dto>? {{ .Name }} { get; set; }
        {{- end }}
      {{- end }}
   {{- end }}
}
{{ end }}