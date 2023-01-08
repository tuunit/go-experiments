package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/santhosh-tekuri/jsonschema"
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
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	schema, err := ioutil.ReadFile("apiobject.schema.json")
	if err != nil {
		log.Fatal(err)
	}

	//data := make(map[interface{}]interface{})
	var data interface{}

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	data = convert(data)

	compiler := jsonschema.NewCompiler()
	if err := compiler.AddResource("schema.json", strings.NewReader(string(schema))); err != nil {
		log.Fatal(err)
	}

	s, err := compiler.Compile("schema.json")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.ValidateInterface(data); err != nil {
		log.Fatal(err)
	}

	fmt.Println("validation successful")
}
