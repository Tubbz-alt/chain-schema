package dsl

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var src1 string = `
struct Test {
  x: uint64
}

enum ABC { A B C }

table Balance {
  account: bytes
  asset: bytes
  balance: decimal
  @primary_key(account, asset)
}

table Group {
  id: auto uint64
  desc: string
  total_weight: decimal
}
`

func TestSimple(t *testing.T) {
	var res Module
	err := Parser.ParseString(src1, &res)
	require.NoError(t, err)
}
