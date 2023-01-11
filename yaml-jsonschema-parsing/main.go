package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

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
	file, err := ioutil.ReadFile("nginx-deployment.yaml")
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

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(string(b))
}
