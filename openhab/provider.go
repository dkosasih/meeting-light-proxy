package openhab

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewOpenhabHandlerFactory,
	wire.Bind(new(OpenhabHandlerCreator), new(*openhabHandlerFactory)),
	NewEndpoints,
)
