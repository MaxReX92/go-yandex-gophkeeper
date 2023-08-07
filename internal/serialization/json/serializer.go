package json

import (
	"encoding/json"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type jsonSerializer struct {
}

func (j *jsonSerializer) Serialize(obj interface{}) ([]byte, error) {
	result, err := json.Marshal(obj)
	if err != nil {
		return nil, logger.WrapError("serialize object", err)
	}

	return result, nil
}

func (j *jsonSerializer) Deserialize(value []byte, obj interface{}) error {
	err := json.Unmarshal(value, obj)
	if err != nil {
		return logger.WrapError("deserialize object", err)
	}

	return nil
}
