package main

import (
	"fmt"
	"io"
	"strings"
	"text/template"
)

//go:generate go-bindata ./templates/
func newTemplateWithFuncMaps(funcMaps template.FuncMap) *template.Template {
	tpl := template.New("file").Funcs(funcMaps)
	filenames := []string{
		"templates/file.tgo",
		"templates/publisher.tgo",
		"templates/subscriber.tgo",
	}
	for _, filename := range filenames {
		tpl = template.Must(tpl.Parse(string(MustAsset(filename))))
	}
	return tpl
}

func generate(w io.Writer, pkg string, typ string, args []string) error {
	eventNames, argTypesMap, err := parseArgs(args)
	if err != nil {
		return err
	}
	tpl := newTemplateWithFuncMaps(template.FuncMap{
		"printArgTypesOf": func(eventName string) string {
			argTypes, ok := argTypesMap[eventName]
			if !ok {
				return ""
			}
			items := []string{}
			for _, v := range argTypes {
				items = append(items, fmt.Sprintf("%s %s", v.Name, v.TypeName))
			}
			return strings.Join(items, ", ")
		},
		"printArgsOf": func(eventName string) string {
			argTypes, ok := argTypesMap[eventName]
			if !ok {
				return ""
			}
			items := []string{}
			for _, v := range argTypes {
				items = append(items, v.Name)
			}
			return strings.Join(items, ", ")
		},
	})
	return tpl.Execute(w, map[string]interface{}{
		"PackageName":       pkg,
		"PublisherTypeName": typ,
		"EventNames":        eventNames,
	})
}

type argType struct {
	Name string

	TypeName string
}

func parseArgs(args []string) (eventNames []string, argTypesMap map[string][]*argType, err error) {
	eventNames = []string{}
	argTypesMap = make(map[string][]*argType)

	for _, arg := range args {
		// Hoge|a:string,b:interface{}
		if strings.ContainsRune(arg, '|') {
			specs := strings.SplitN(arg, "|", 2)
			eventName := strings.TrimSpace(specs[0])
			eventNames = append(eventNames, eventName)
			argsSpecs := strings.Split(specs[1], ",")
			for _, argSpec := range argsSpecs {
				nameAndType := strings.SplitN(argSpec, ":", 2)
				argTypesMap[eventName] = append(argTypesMap[eventName], &argType{
					Name:     strings.TrimSpace(nameAndType[0]),
					TypeName: strings.TrimSpace(nameAndType[1]),
				})
			}
		} else {
			eventNames = append(eventNames, arg)
		}
	}
	return
}
