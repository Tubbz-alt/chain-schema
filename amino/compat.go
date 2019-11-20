package amino

import (
	"github.com/tendermint/go-amino"
)

type codec struct {
	amino.Codec
}

var _ amino.Codec = codec{}


