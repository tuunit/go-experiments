package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-openapi/spec"
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

func main() {
	file, err := os.ReadFile("schema.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data interface{}

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	raw, err := json.Marshal(convert(data))
	if err != nil {
		fmt.Println(err)
		return
	}

	var schema spec.Schema

	err = schema.UnmarshalJSON(raw)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(schema.Properties["name"].Example)

}
