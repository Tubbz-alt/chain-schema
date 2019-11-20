package proto

import (
	"fmt"
	chain_schema "github.com/regen-network/chain-schema"
	"github.com/regen-network/chain-schema/codegen"
)

func EmitModule(writer codegen.Writer, def *chain_schema.ModuleDef) {
	for _, td := range def.Typedefs {
		EmitTypeDefinition(writer, td)
		writer.Write("\n")
	}
}

func EmitTypeDefinition(writer codegen.Writer, x *chain_schema.TypeDefinition) {
	switch info := x.Info.(type) {
	case *chain_schema.TypeDefinition_Struct:
		EmitStruct(writer, x, info.Struct)
	case *chain_schema.TypeDefinition_Enum:
		EmitEnum(writer, x, info.Enum)
	}
}

func EmitStruct(writer codegen.Writer, typeDef *chain_schema.TypeDefinition, str *chain_schema.Struct) {
	writer.Write(fmt.Sprintf("message %s {\n", typeDef.Name))
	for _, field := range str.Fields {
		EmitField(writer, field)
	}
	writer.Write(fmt.Sprintf("}\n"))
}

func EmitField(writer codegen.Writer, field *chain_schema.Field) {
	writer.Write(fmt.Sprintf("  %s %s = %d;\n",
		"bytes",
		field.Name,
		field.ProtoField,
	))
}

func EmitEnum(writer codegen.Writer, definition *chain_schema.TypeDefinition, enum *chain_schema.Enum) {
	writer.Write(fmt.Sprintf("enum %s {\n", definition.Name))
	for _, val := range enum.Values {
		writer.Write(fmt.Sprintf("  %s = %d;\n", val.Name, val.Value))
	}
	writer.Write(fmt.Sprintf("}\n"))
}

