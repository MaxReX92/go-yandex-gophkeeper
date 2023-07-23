package rand

type randomGenerator struct {
}

func NewGenerator() *randomGenerator {
	return &randomGenerator{}
}

func (r *randomGenerator) GenerateNewIdentity() string {
	//TODO implement me
	panic("implement me")
}
