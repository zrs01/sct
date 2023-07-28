{{- data := .Data }}
{{- range .Entity}}
{{- entity := .Name }}
/* -------------------------------- INTERFACE ------------------------------- */
using Cpas.Common;

namespace Cpas.{{ data.Team }}.Repo;

public interface I{{ .Name }}Service : ICRUDService<{{ .Name }}, int> {
}

/* --------------------------------- SERVICE -------------------------------- */
using Cpas.Common;
using Cpas.{{ data.Team }}.Repo;
using Microsoft.Extensions.DependencyInjection;

namespace {{ data.Namespace }};

public class {{ .Name }}Service : CRUDService<{{ data.Team }}DbContext, {{ .Name }}, int>, I{{ .Name }}Service {
    private readonly {{ data.Team }}DbContext _context;
    private readonly IServiceProvider _serviceProvider;

    public {{ .Name }}Service(IServiceProvider serviceProvider, {{ data.Team }}DbContext context) : base(serviceProvider, context) {
        _context = serviceProvider.GetRequiredService<{{ data.Team }}DbContext>();
        _serviceProvider = serviceProvider;
    }

{{- if .IsContainsVirtual || .IsContainsCollection }}

    protected override async Task<SearchResult<CpaRpt>> SearchActionAsync(SearchCriteria criteria) {
        return await base.VersatileSearchAsync(criteria, async (query) => {
            foreach (var option in (criteria.options ?? new Dictionary<string, string>())) {
  {{- range .Members }}
    {{- if .IsVirtual && !.IsCollection }}
                if (string.Equals(option.Key, "{{ .Name }}", StringComparison.OrdinalIgnoreCase)) {
                    query = query.Where(x => _context.{{ .Name }}.RSql(option.Value).Select(x => x.{{ .Name }}Id).Contains(x.{{ .Name }}Id));
                }
    {{- end }}
    {{- if .IsCollection }}
                if (string.Equals(option.Key, "{{ .Name }}", StringComparison.OrdinalIgnoreCase)) {
                    query = query.Where(x => _context.{{ .Name }}.RSql(option.Value).Select(x => x.{{ entity }}Id).Contains(x.{{ entity }}Id));
                }
    {{- end }}
  {{- end }}
            }
            await Task.CompletedTask;
            return query;
        });
    }
{{- end }}
{{- if .IsContainsCollection }}

    public override async Task<CpaRpt> UpdateActionAsync<T>(int id, T dto) {
        var typedDto = dto as {{ entity }}Dto;
        if (typedDto != null) {
  {{- range .Members }}
    {{- if .IsCollection }}
            if (typedDto.{{ .Name }} != null) {
                var daoSet = await _context.{{ .Name }}.Where(x => x.{{ entity }}Id == typedDto.{{ entity }}Id).ToListAsync();
                await _serviceProvider.GetRequiredService<I{{ .Name }}Service>().UpdateCollection(typedDto.{{ .Name }}, daoSet);
            }
    {{- end }}
  {{- end }}
        }
        return await base.UpdateActionAsync(id, dto);
    }
{{- end }}
{{- if .IsContainsCollection }}

    public override async Task DeleteActionAsync<T>(int id) {
  {{- range .Members }}
    {{- if .IsCollection }}
        foreach (var dao in await _context.{{ .Name }}.Where(x => x.{{ .Name }}Id == id).ToListAsync()) {
            await _serviceProvider.GetRequiredService<I{{ .Name }}Service>().DeleteActionAsync<{{ .Name }}>(dao.{{ .Name }}Id);
        }
    {{- end }}
  {{- end }}
        await base.DeleteActionAsync<T>(id);
    }
{{- end }}
}

/* ------------------------- Add below to Program.cs ------------------------ */
builder.Services.AddScoped<Cpas.{{ data.Team }}.Repo.I{{ .Name }}Service, Cpas.{{ data.Team }}.Service.{{ .Name }}Service>();

/* ------------------------------- CONTROLLER ------------------------------- */
using Cpas.{{ data.Team }}.Repo;
using Microsoft.AspNetCore.Mvc;

namespace Cpas.App.Controllers.{{ data.Team }};

[Authorize]
[Route("api/{{ lower(data.Team) }}/[controller]")]
public class {{ .Name }}Controller : RestController<{{ .Name }}Dto, {{ .Name }}, int, I{{ .Name }}Service> {
    public {{ .Name }}Controller(IServiceProvider serviceProvider) : base(serviceProvider) { }
}

/* ----------------------------- ANGULAR SERVICE ---------------------------- */
import { Injectable } from '@angular/core';
import { BaseCRUDService } from 'src/app/share/services/base-crud.service';
import { SystemInfo } from 'src/app/share/utilities/system-info';

@Injectable({
  providedIn: 'root'
})
export class {{ data.Team }}{{ .Name }}Service extends BaseCRUDService<{{ .Name }}> {

  constructor() {
    super(SystemInfo.data.team + '/{{ .Name }}');
  }
}

{{ end }}