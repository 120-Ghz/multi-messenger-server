package tools

import (
	"encoding/json"
	"io"
)

func ParseBody(body io.ReadCloser) (map[string]interface{}, error) {
	parsedBody, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	var bodyMap map[string]interface{}
	err = json.Unmarshal(parsedBody, &bodyMap)
	if err != nil {
		return nil, err
	}
	return bodyMap, nil
}
