package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/buger/jsonparser"
)

func main() {
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
		"name":          "${string}",
		"age":           "${number}",
		"object":        "${object}",
		"array":         "${array}",
		"Authorization": "Bearer ${string}",
		"nested":        "Bearer ${object.key}",
		"nested2":       "Bearer ${object.key2.key3}",
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile(`\$\{(.*?)\}`)
	matches := r.FindAllStringSubmatch(string(bodyJson), -1)
	result := string(bodyJson)
	for _, match := range matches {
		keys := strings.Split(match[1], ".")
		token := match[0]
		val, typ, _, err := jsonparser.Get(responseJson, keys...)
		if err != nil {
			panic(err)
		}
		switch typ {
		case jsonparser.Object:
			fallthrough
		case jsonparser.Array:
			fallthrough
		case jsonparser.Number:
			newToken := fmt.Sprintf("\"%s\"", token)
			if strings.Contains(result, newToken) {
				token = newToken
			}
		default:
		}
		result = strings.Replace(result, token, string(val), -1)
	}
	var indentedResult bytes.Buffer
	err = json.Indent(&indentedResult, []byte(result), "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(indentedResult.Bytes()))
}
