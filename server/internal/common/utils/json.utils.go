package utils

import "encoding/json"

func JsonSerialize(obj any) (string, error) {
	str, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	return string(str), nil
}

func JsonDeserialize[TOut any](str string) (*TOut, error) {
	var t TOut
	err := json.Unmarshal([]byte(str), &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
