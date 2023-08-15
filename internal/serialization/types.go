package serialization

// Serializer provides methods for secrets content serialization.
type Serializer interface {
	// Serialize present secret as byte array.
	Serialize(obj interface{}) ([]byte, error)
	// Deserialize present byte array as secret.
	Deserialize(value []byte, obj interface{}) error
}
