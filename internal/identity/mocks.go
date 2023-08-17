package identity

import "github.com/stretchr/testify/mock"

type GeneratorMock struct {
	mock.Mock
}

func (g *GeneratorMock) GenerateNewIdentity() string {
	args := g.Called()
	return args.String(0)
}
