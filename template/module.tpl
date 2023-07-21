{{- typescripts := .Ts }}
{{- range i, Cs := .Cs }}
{{- Ts := typescripts[i] }}
// ------------------------------------ DTO -----------------------------------
using Cms.Repo.Entity;

namespace Cms.Repo.Model;

public class {{ Cs.Name }}Dto : {{ Cs.Name }} {
    public new byte[]? RowVer { get; set; }
    {{- range Cs.Members }}
      {{- if .Virtual }}
        {{- if len(.CollectionType) == 0 }}
    public new {{ .DataType }}Dto? {{ .Name }} { get; set; }
        {{- else }}
    public new ICollection<{{ .CollectionType }}Dto>? {{ .Name }} { get; set; }
        {{- end }}
      {{- end }}
   {{- end }}
}

// ----------------------------- SERVICE INTERFACE ----------------------------
using Cms.Repo.Entity;

namespace Cms.Service.Support;

public interface I{{ Cs.Name }}Service : ICRUDService<{{ Cs.Name }}, int> {
}

// -------------------------- SERVICE IMPLEMENTATION --------------------------
using Cms.Repo;
using Cms.Repo.Entity;
using Cms.Service.Support;
using Microsoft.Extensions.DependencyInjection;

namespace {{ Cs.Namespace }};

public class {{ Cs.Name }}Service : CRUDService<{{ Cs.Name }}, int>, I{{ Cs.Name }}Service {
    private readonly CmsDbContext _context;
    private readonly IServiceProvider _serviceProvider;

    public {{ Cs.Name }}Service(IServiceProvider serviceProvider) : base(serviceProvider) {
        _context = serviceProvider.GetRequiredService<CmsDbContext>();
        _serviceProvider = serviceProvider;
    }
}

// ------------------------- ADD BELOW TO 'Program.cs' ------------------------
builder.Services.AddScoped<I{{ Cs.Name }}Service, {{ Cs.Name }}Service>();

// -------------------------------- CONTROLLER --------------------------------
using Cms.App.Auth;
using Cms.Repo.Entity;
using Cms.Repo.Model;
using Cms.Service.Support;
using Microsoft.AspNetCore.Mvc;

namespace Cms.App.Controllers;

[Authorize]
[Route("api/[controller]")]
public class {{ Cs.Name }}Controller : RestController<{{ Cs.Name }}Dto, {{ Cs.Name }}, int, I{{ Cs.Name }}Service> {
    public {{ Cs.Name }}Controller(IServiceProvider serviceProvider) : base(serviceProvider) { }
}

# ------------------------------- ANGULAR MODEL ------------------------------ #
export interface {{ Ts.Name }} {
  {{- range Ts.Members }}
  {{ lower(.Name[:1]) + .Name[1:] }}?: {{ .DataType }};
  {{- end }}
}

// ------------------------------ ANGULAR SERVICE -----------------------------
import { Injectable } from '@angular/core';
import { BaseCRUDService } from 'src/app/share/services/base-crud.service';

@Injectable({
  providedIn: 'root'
})
export class {{ Ts.Name }}Service extends BaseCRUDService<{{ Ts.Name }}> {

  constructor() {
    super('{{ Ts.Name }}');
  }

}
{{ end }}