package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

type ValidationFields struct {
	Name     string
	Type     string
	MinValue int
	MaxValue int
	In       []int
	Len      int
	InString []string
	RE       string
	IsMin    bool
	IsMax    bool
	IsIn     bool
	IsLen    bool
	IsRE     bool
}

type ValidationStruct struct {
	Name   string
	Fields []ValidationFields
}

type ValidationPackage struct {
	Package          string
	ValidatedStructs []ValidationStruct
}

func main() {
	filePathes := os.Args[1:]
	for _, filePath := range filePathes {
		fd, err := os.Open(filePath)
		if err != nil {
			log.Fatalln("Error during open file", filePath, err)
		}

		fw, err := os.Create(strings.ReplaceAll(filePath, ".go", "_validation_generated.go"))
		if err != nil {
			log.Fatalln("Error during open file", err)
		}

		fset := token.NewFileSet()
		pfile, err := parser.ParseFile(fset, filePath, fd, parser.AllErrors)
		if err != nil {
			log.Fatalln("Error during parsing file", filePath, err)
		}

		tmpl := template.New("validator")
		tmpl, err = tmpl.Parse(templateFile)
		if err != nil {
			log.Fatalln("error during creating template")
		}

		values := ParseValidatedFile(pfile)

		err = tmpl.Execute(fw, values)
		if err != nil {
			log.Fatalln("error during templating")
		}
		fd.Close()
		fw.Close()
	}
}

func ParseValidatedFile(af *ast.File) ValidationPackage {
	values := newValidationPackage(af.Name.Name)

	derivedTypes := getDerivedTypes(af)

	ast.Inspect(af, func(node ast.Node) bool {
		switch elem := node.(type) { //nolint:gocritic
		case *ast.TypeSpec:
			structType, ok := elem.Type.(*ast.StructType)
			if !ok {
				return true
			}

			structure := newValidationStruct(elem.Name.Name)

			for _, field := range structType.Fields.List {
				if field.Tag == nil {
					continue
				}

				tags := ParseValidateTag(field.Tag.Value)
				if len(tags) == 0 {
					continue
				}

				var typeName string

				if t, ok := field.Type.(*ast.Ident); ok {
					typeName = t.Name
					if derived, ok := derivedTypes[typeName]; ok {
						typeName = derived
					}
				} else if t, ok := field.Type.(*ast.ArrayType); ok {
					typeName = t.Elt.(*ast.Ident).Name
					if derived, ok := derivedTypes[typeName]; ok {
						typeName = derived
					}
					typeName = "[]" + typeName
				}

				validationField := ValidationFields{
					Name: field.Names[0].Name,
					Type: typeName,
				}

				if ok := fillValidationFields(typeName, tags, &validationField); !ok {
					return false
				}

				structure.Fields = append(structure.Fields, validationField)
			}
			if len(structure.Fields) > 0 {
				values.ValidatedStructs = append(values.ValidatedStructs, structure)
			}
		}

		return true
	})

	return values
}

func fillValidationFields(typeName string, tags map[string]string, validationField *ValidationFields) bool {
	switch typeName {
	case
		"int",
		"[]int":
		if val, ok := tags["min"]; ok {
			validationField.IsMin = true
			validationField.MinValue, _ = strconv.Atoi(val)
		}
		if val, ok := tags["max"]; ok {
			validationField.IsMax = true
			validationField.MaxValue, _ = strconv.Atoi(val)
		}
		if val, ok := tags["in"]; ok {
			validationField.IsIn = true
			inValues := strings.Split(val, ",")
			for _, elem := range inValues {
				intValue, err := strconv.Atoi(elem)
				if err != nil {
					return false
				}
				validationField.In = append(validationField.In, intValue)
			}
		}
	case
		"string",
		"[]string":
		if val, ok := tags["len"]; ok {
			validationField.IsLen = true
			validationField.Len, _ = strconv.Atoi(val)
		}
		if val, ok := tags["regexp"]; ok {
			validationField.IsRE = true
			validationField.RE = val
		}
		if val, ok := tags["in"]; ok {
			validationField.IsIn = true
			validationField.InString = strings.Split(val, ",")
		}
	}

	return true
}

//nolint:interfacer
func getDerivedTypes(af *ast.File) map[string]string {
	result := map[string]string{}

	ast.Inspect(af, func(node ast.Node) bool {
		switch elem := node.(type) { //nolint:gocritic
		case *ast.TypeSpec:
			identType, ok := elem.Type.(*ast.Ident)
			if ok {
				result[elem.Name.Name] = identType.Name
			}
		}

		return true
	})

	return result
}

func ParseValidateTag(tag string) map[string]string {
	result := make(map[string]string)

	trimmedTag := strings.Trim(tag, "`")
	tags := strings.Split(trimmedTag, " ")
	for _, elem := range tags {
		splittedTag := strings.SplitN(elem, ":", 2)
		if splittedTag[0] != "validate" {
			continue
		}
		validateTagValue := strings.Trim(splittedTag[1], "\"")
		validateTagValues := strings.Split(validateTagValue, "|")
		for _, validateElem := range validateTagValues {
			splittedValElem := strings.SplitN(validateElem, ":", 2)
			result[splittedValElem[0]] = splittedValElem[1]
		}
	}

	return result
}

func newValidationPackage(name string) ValidationPackage {
	return ValidationPackage{
		Package:          name,
		ValidatedStructs: []ValidationStruct{},
	}
}

func newValidationStruct(name string) ValidationStruct {
	return ValidationStruct{
		Name:   name,
		Fields: []ValidationFields{},
	}
}
