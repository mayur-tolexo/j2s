
{{$structName:= .Name}}
//{{$structName}} model
type {{$structName}} struct {
{{- range $jsonName, $val := .Fields}}
	{{ Title $jsonName }} {{ (TypeOf $structName $jsonName $val) }} `json:"{{ $jsonName}}"`
{{- end}}
}