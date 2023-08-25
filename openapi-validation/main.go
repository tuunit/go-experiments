package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"gopkg.in/yaml.v2"
)

func convert(i interface{}) interface{} {
	switch x := i.(type) {
	case map[interface{}]interface{}:
		m2 := map[string]interface{}{}
		for k, v := range x {
			m2[k.(string)] = convert(v)
		}
		return m2
	case []interface{}:
		for i, v := range x {
			x[i] = convert(v)
		}
	}
	return i
}

func readSchema() (*spec.Schema, error) {
	file, err := os.ReadFile("schema.yaml")
	if err != nil {
		return nil, err
	}

	var data interface{}

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	raw, err := json.Marshal(convert(data))
	if err != nil {
		return nil, err
	}

	var schema spec.Schema

	err = schema.UnmarshalJSON(raw)
	if err != nil {
		return nil, err
	}

	return &schema, nil
}

func readYaml(filename string) (interface{}, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data interface{}

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return convert(data), nil
}

func validateYaml(schema *spec.Schema, data interface{}) {
	fmt.Println("--- validation ---")
	err := validate.AgainstSchema(schema, data, strfmt.NewFormats())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("validation was successful")
	}
}

func main() {
	schema, err := readSchema()
	if err != nil {
		fmt.Printf("error while reading schema: %s", err)
		return
	}

	file, err := readYaml("wrong.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	validateYaml(schema, file)

	file, err = readYaml("correct.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	validateYaml(schema, file)
}
