package json

import (
	json "encoding/json"

	"github.com/pkg/errors"
)

func Stringify(v interface{}) (string, error) {
	byteData, err := json.Marshal(v)
	if err != nil {
		return "", errors.WithStack(err)
	}

	return string(byteData), nil
}

func Parse(s string, v interface{}) error {
	if err := json.Unmarshal([]byte(s), v); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
