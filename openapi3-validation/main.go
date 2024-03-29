package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	ctx := context.Background()
	load := &openapi3.Loader{Context: ctx}
	spec, _ := load.LoadFromFile("schema.yaml")

	if err := spec.Validate(ctx); err != nil {
		fmt.Println(err)
	}

	opts := make([]openapi3.SchemaValidationOption, 0, 4) // 4 potential opts here
	opts = append(opts, openapi3.VisitAsRequest())

	data := `{
		"complete": false,
		"id": 10,
		"petId": 1,
		"quantity": true,
		"shipDate": "2021-07-01T17:00:00Z",
		"status": "placed"
	}`

	value := make(map[string]interface{})

	if err := json.Unmarshal([]byte(data), &value); err != nil {
		fmt.Println(err)
	}

	if err := spec.Components.Schemas["Order"].Value.VisitJSON(value, opts...); err != nil {
		fmt.Println(err)
	}

}
