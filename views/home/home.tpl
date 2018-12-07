{{ define "content" }}
<h1>{{.Data.PageTitle}}<h1>
<ul>
    {{range .Data.Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
{{ end }}
