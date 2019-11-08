package jsonToStruct

//{{.Name}} model
type {{ .Name }} struct {
{{- range $jsonName, $val := .Fields}}
	{{ Title $jsonName }} {{ (TypeOf $val) }}	`json:"{{ $jsonName}}"`
{{- end}}
}