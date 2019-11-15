package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var (
	strMap   map[string]string
	hashMap  map[string]string
	strOrder []string
	tmplStr  string
)

func init() {
	strMap = make(map[string]string)
	hashMap = make(map[string]string)
	strOrder = make([]string, 0)
	initTmpl()
}

//Parser model
type Parser struct {
	Name    string
	Fields  map[string]interface{}
	strMap  map[string]string
	hashMap map[string]string
}

//getJSONData will return json data
func getJSONData(file string, raw int) (data map[string]interface{}, err error) {

	var (
		byteValue []byte
		jsonValue string
	)
	if raw == 1 {
		fmt.Println("Enter the json between tilt (~): ")
		fmt.Scanf("%q", &jsonValue)
		err = json.Unmarshal([]byte(jsonValue), &data)
	} else if byteValue, err = readJSONFile(file); err == nil {
		err = json.Unmarshal(byteValue, &data)
	}
	return
}

//readJSONFile will read json file
func readJSONFile(file string) (byteValue []byte, err error) {
	var jsonFile *os.File
	if jsonFile, err = os.Open(file); err == nil {
		defer jsonFile.Close()
		byteValue, err = ioutil.ReadAll(jsonFile)
	}
	return
}

func main() {

	var data Parser

	ip := flag.String("ip", "input.json", "Input File")
	op := flag.String("op", "output.go", "Output File")
	name := flag.String("name", "User", "Structure Name")
	raw := flag.Int("raw", 0, "1 to enable console json input")
	flag.Parse()

	curPkg := getCurPkg()
	body, err := getJSONData(*ip, *raw)
	ifError(err)
	data.Name = *name
	data.Fields = body
	createStruct(data)
	fp, err := os.Create(*op)
	if err == nil {
		fp.WriteString("package " + curPkg + "\n")
		for i := len(strOrder) - 1; i >= 0; i-- {
			fp.WriteString(strMap[strOrder[i]] + "\n")
		}
	}
	err = exec.Command("gofmt", "-w", *op).Run()
	ifError(err)
	fmt.Println(*op, "Created")
}

//getCurPkg will return current pkg name
func getCurPkg() (curPkg string) {
	dir, err := os.Getwd()
	ifError(err)
	curPkg = filepath.Base(dir)
	return
}

//ifError will print error if not nil
func ifError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createStruct(data Parser) {
	var hash string
	fName := getFieldName(data.Name)
	strMap[fName], hash = getStructStr(data)
	hashMap[hash] = fName
	strOrder = append(strOrder, data.Name)
	return
}

//getStructStr will execute template and return string and struct hash
func getStructStr(data Parser) (string, string) {

	hash := ""
	hFn := getHashFn(&hash)
	tmpl, err := template.New("template").Funcs(template.FuncMap{
		"Title": getFieldName,
		"TypeOf": func(p string, k string, v interface{}) string {
			if v == nil {
				return "string"
			}

			rType := reflect.TypeOf(v)
			if isStruct(rType) {
				return getSubStructType(p, k, v)
			} else if rType.Kind() == reflect.Slice {

				rVal := reflect.ValueOf(v)
				if rVal.Len() > 0 {
					fVal := rVal.Index(0)
					if isStruct(fVal.Elem().Type()) {
						return "[]" + getSubStructType(p, k, fVal.Interface())
					}
					return "[]" + getType(v, fVal.Elem().Type())
				}
			}
			return getType(v, rType)
		},
		"Hash": hFn,
	}).Parse(tmplStr)
	ifError(err)
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	ifError(err)
	// tmpl.ExecuteTemplate(&buf, "template.tpl", data)
	return buf.String(), hash
}

func getHashFn(gHash *string) func(string, string) string {
	return func(name string, hash string) string {
		*gHash += hash
		return ""
	}
}

func getType(v interface{}, rType reflect.Type) string {
	if rType.Kind() == reflect.Float64 {
		if isIntegral(v.(float64)) {
			return "int"
		}
	}
	return strings.ToLower(rType.String())
}

//isIntegral will check float is int
func isIntegral(val float64) bool {
	return val == float64(int(val))
}

//isStruct will check given ref value is struct type i.e. map[string]interface{}
func isStruct(rType reflect.Type) (isStruct bool) {
	if rType.Kind() == reflect.Map && rType.String() == "map[string]interface {}" {
		isStruct = true
	}
	return
}

//getSubStructType will return map field name after creating struct
func getSubStructType(p string, k string, v interface{}) string {

	name := getFieldName(k)
	subData := getParserModel(name, v)
	newStr, hash := getStructStr(subData)

	//if any struct already exists with same hash
	if name, exists := hashMap[hash]; exists {
		return name
	}

	//if struct with same name already exists
	if curStr, exists := strMap[name]; exists {
		if curStr == newStr {
			return name
		}
		name = getFieldName(p + k)
	}

	subData = getParserModel(name, v)
	createStruct(subData)
	return name
}

func getParserModel(name string, v interface{}) Parser {
	return Parser{
		Name:   name,
		Fields: v.(map[string]interface{}),
	}
}

//getFieldName will return field name in Camel Case
func getFieldName(k string) (f string) {
	f = strcase.ToCamel(k)
	r := strings.NewReplacer(
		"Id", "ID",
		"id", "ID",
		"Api", "API",
		"Http", "HTTP",
	)
	f = r.Replace(f)
	return
}

func initTmpl() {
	tmplStr = `
{{$structName:= Title .Name}}
//{{$structName}} model
type {{$structName}} struct {
{{- range $jsonName, $val := .Fields}}
	{{ $cField := Title $jsonName -}}
	{{ $cType := (TypeOf $structName $jsonName $val) -}}
	{{ $cField }} {{ $cType }} ` + "`json:\"{{ $jsonName }}\"`" + `
	{{- $cHash := print "field_" $cField "_type_" $cType "_json_" $jsonName }}
	{{- Hash $structName $cHash}}
{{- end}}
}`
}
