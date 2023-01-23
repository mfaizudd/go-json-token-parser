package gojsontokenparser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/buger/jsonparser"
)

func Parse(input []byte, data []byte) (string, error) {
	r := regexp.MustCompile(`\$\{(.*?)\}`)
	matches := r.FindAllStringSubmatch(string(input), -1)
	result := string(input)
	for _, match := range matches {
		var keys []string
		if match[1] == "." {
			keys = []string{}
		} else {
			keys = strings.Split(match[1], ".")
		}
		token := match[0]
		val, typ, _, err := jsonparser.Get(data, keys...)
		if err != nil {
			return "", err
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
			} else {
				val = []byte(strings.Replace(string(val), "\"", "\\\"", -1))
			}
		default:
		}
		result = strings.Replace(result, token, string(val), -1)
	}
	return result, nil
}
