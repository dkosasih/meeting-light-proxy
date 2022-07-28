package albums

import (
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	NewAlbumHandlerFactory,
	wire.Bind(new(AlbumHandlerCreator), new(*albumHandlerFactory)),
	NewEndpoints,
)
