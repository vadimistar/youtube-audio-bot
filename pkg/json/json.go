package json

import (
	"bytes"
	"encoding/json"
	"io"
)

func Serialize(data interface{}) (io.Reader, error) {
	dataBuffer := new(bytes.Buffer)

	err := json.NewEncoder(dataBuffer).Encode(data)
	if err != nil {
		return nil, err
	}

	return dataBuffer, nil
}

func Deserialize[T any](data io.Reader) (T, error) {
	var deserializedData T

	err := json.NewDecoder(data).Decode(&deserializedData)
	if err != nil {
		return deserializedData, err
	}

	return deserializedData, nil
}
