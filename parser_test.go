package gojsontokenparser

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestParse(t *testing.T) {
	response := map[string]interface{}{
		"string": "some string",
		"number": 123,
		"object": map[string]interface{}{
			"key": "value",
			"key2": map[string]interface{}{
				"key3": "value3",
			},
		},
		"array": []interface{}{
			"value1",
			"value2",
		},
	}
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	body := map[string]interface{}{
		"name":                "${string}",
		"age":                 "${number}",
		"object":              "${object}",
		"array":               "${array}",
		"arrayElement":        "${array.[0]}",
		"Authorization":       "Bearer ${string}",
		"nested":              "Bearer ${object.key}",
		"nested2":             "Bearer ${object.key2.key3}",
		"all":                 "${.}",
		"nested_object":       "${object.key2}",
		"object_with_message": "Message: ${object.key2}",
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	result, err := Parse(bodyJson, responseJson)
	if err != nil {
		panic(err)
	}
	var indentedResult bytes.Buffer
	err = json.Indent(&indentedResult, []byte(result), "", "    ")
	if err != nil {
		panic(err)
	}
	expected := `{
    "Authorization": "Bearer some string",
    "age": 123,
    "all": {
        "array": [
            "value1",
            "value2"
        ],
        "number": 123,
        "object": {
            "key": "value",
            "key2": {
                "key3": "value3"
            }
        },
        "string": "some string"
    },
    "array": [
        "value1",
        "value2"
    ],
    "arrayElement": "value1",
    "name": "some string",
    "nested": "Bearer value",
    "nested2": "Bearer value3",
    "nested_object": {
        "key3": "value3"
    },
    "object": {
        "key": "value",
        "key2": {
            "key3": "value3"
        }
    },
    "object_with_message": "Message: {\"key3\":\"value3\"}"
}`
	if expected != string(indentedResult.Bytes()) {
		t.Errorf("Expected %s, got %s", expected, string(indentedResult.Bytes()))
	}
}
