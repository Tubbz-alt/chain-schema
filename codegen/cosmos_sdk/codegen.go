package cosmos_sdk

import (
	chain_schema "github.com/regen-network/chain-schema"
	"io"
)

//var moduleTmpl = template.Must(template.New("module").Parse(`
//{{range .Decls}}
//  {{.}}
//{{end}}
//`))

func Emit(td chain_schema.TypeDefinition, wr io.Writer) {

}
