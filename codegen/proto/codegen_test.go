package proto

import (
	"fmt"
	"github.com/regen-network/chain-schema/codegen"
	"github.com/regen-network/chain-schema/dsl"
	"github.com/stretchr/testify/require"
	"testing"
)

var src1 string = `
struct Test {
  x: uint64
  y: string
}

enum ABC { A B C }
`

func TestSimple(t *testing.T)  {
	var res dsl.Module
	err := dsl.Parser.ParseString(src1, &res)
	require.NoError(t, err)
	m := res.Emit()
	w := codegen.NewWriter()
	EmitModule(w, m)
	fmt.Println(w.String())
}

