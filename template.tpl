
{{$structName:= Title .Name}}
//{{$structName}} model
type {{$structName}} struct {
{{- range $jsonName, $val := .Fields}}
	{{ $cField := Title $jsonName -}}
	{{ $cType := (TypeOf $structName $jsonName $val) -}}
	{{ $cField }} {{ $cType }} `json:"{{ $jsonName }}"`
	{{- $cHash := print "field_" $cField "_type_" $cType "_json_" $jsonName }}
	{{- Hash $structName $cHash}}
{{- end}}
}