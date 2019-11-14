
//{{.Name}} model
type {{Title .Name }} struct {
{{- range $jsonName, $val := .Fields}}
	{{ Title $jsonName }}	{{ (TypeOf $jsonName $val) }}	`json:"{{ $jsonName}}"`
{{- end}}
}