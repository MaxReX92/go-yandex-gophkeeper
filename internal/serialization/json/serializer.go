package json

import (
	"encoding/json"

	"github.com/MaxReX92/go-yandex-gophkeeper/pkg/logger"
)

type jsonSerializer struct{}

// NewSerializer creates a new instance of json secrets serializer instance.
func NewSerializer() *jsonSerializer {
	return &jsonSerializer{}
}

func (j *jsonSerializer) Serialize(obj interface{}) ([]byte, error) {
	result, err := json.Marshal(obj)
	logger.InfoFormat("SERIALIZED: %s", string(result))
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
