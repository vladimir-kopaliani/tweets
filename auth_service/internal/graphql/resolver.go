package graphql

import (
	"github.com/vladimir-kopaliani/tweets/auth_service/internal/interfaces"
)

type Resolver struct {
	Service interfaces.Servicer
	Logger  interfaces.Logger
}
