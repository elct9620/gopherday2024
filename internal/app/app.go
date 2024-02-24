package app

import "github.com/google/wire"

var DefaultSet = wire.NewSet(
	ProvideBoltDB,
)
