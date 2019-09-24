package testutils

import (
	"encoding/json"
	"io/ioutil"
)

// ReadFileAsJSON reads a file and converts it to a json structure
func ReadFileAsJSON(path string) (deserializedFile map[string]interface{}, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := map[string]interface{}{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
