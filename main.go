package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
)

var (
	strMap   map[string]string
	strOrder []string
)

func init() {
	strMap = make(map[string]string)
	strOrder = make([]string, 0)

}

//Parser model
type Parser struct {
	Name   string
	Fields map[string]interface{}
}

func main() {

	var (
		data Parser
		body map[string]interface{}
	)

	jsonFile, err := os.Open("input3.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &body)
	data.Name = "User"
	data.Fields = body
	file := "output.go"

	createStruct(data)
	fp, err := os.Create(file)
	if err == nil {
		fp.WriteString("package jsonToStruct\n")
		for i := len(strOrder) - 1; i >= 0; i-- {
			fp.WriteString(strMap[strOrder[i]] + "\n")
		}
	}
	err = exec.Command("gofmt", "-w", file).Run()
	ifError(err)
}

func ifError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func createStruct(data Parser) {
	strMap[data.Name] = getStructStr(data)
	strOrder = append(strOrder, data.Name)
	return
}

func getStructStr(data Parser) string {
	tmpl, _ := template.New("template").Funcs(template.FuncMap{
		"Title": getFieldName,
		"TypeOf": func(k string, v interface{}) string {
			if v == nil {
				return "string"
			}

			rType := reflect.TypeOf(v)
			if isStruct(rType) {
				return getMapFieldName(data.Name, k, v)
			} else if rType.Kind() == reflect.Slice {

				rVal := reflect.ValueOf(v)
				if rVal.Len() > 0 {
					fVal := rVal.Index(0)
					if isStruct(fVal.Elem().Type()) {
						return "[]" + getMapFieldName(data.Name, k, fVal.Interface())
					}
					return "[]" + getType(fVal.Elem().Type())
				}
			}
			return getType(rType)
		},
	}).ParseFiles("template.tpl")

	var buf bytes.Buffer
	tmpl.ExecuteTemplate(&buf, "template.tpl", data)
	return buf.String()
}

func getType(rType reflect.Type) string {
	return strings.ToLower(rType.String())
}

//isStruct will check given ref value is struct type i.e. map[string]interface{}
func isStruct(rType reflect.Type) (isStruct bool) {
	if rType.Kind() == reflect.Map && rType.String() == "map[string]interface {}" {
		isStruct = true
	}
	return
}

//getMapFieldName will return map field name after creating struct
func getMapFieldName(pName string, k string, v interface{}) string {

	name := getFieldName(k)
	if v1, exists := strMap[name]; exists {
		subData := getParserModel(name, v)
		v2 := getStructStr(subData)
		if v1 == v2 {
			return name
		}
	}

	subData := getParserModel(name, v)
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
