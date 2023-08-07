package serialization

type Serializer interface {
	Serialize(obj interface{}) ([]byte, error)
	Deserialize(value []byte, obj interface{}) error
}
