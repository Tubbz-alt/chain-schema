package dsl

import chain_schema "github.com/regen-network/chain-schema"

func (m Module) Emit() *chain_schema.ModuleDef {
	res := &chain_schema.ModuleDef{}
	for _, decl := range m.Decls {
		if decl.Struct != nil {
			td := decl.Struct.Emit()
			res.Typedefs = append(res.Typedefs, td)
		} else if decl.Enum != nil {
			e := decl.Enum.Emit()
			res.Typedefs = append(res.Typedefs, e)
		}
	}
	return res
}

func (s Struct) Emit() *chain_schema.TypeDefinition {
	s2 := &chain_schema.Struct{}
	for _, f := range s.Fields {
		s2.Fields = append(s2.Fields, f.Emit())
	}
	res := &chain_schema.TypeDefinition{
		Namespace:    0,
		Name:         s.Name,
		Experimental: false,
		Info:         &chain_schema.TypeDefinition_Struct{Struct: s2},
	}
	return res
}

func (f Field) Emit() *chain_schema.Field {
	res := &chain_schema.Field{
		Name: f.Name,
		Type: nil,
	}
	return res
}

func (e Enum) Emit() *chain_schema.TypeDefinition {
	e2 := &chain_schema.Enum{}
	for _, v := range e.Values {
		e2.Values = append(e2.Values, v.Emit())
	}
	res := &chain_schema.TypeDefinition{
		Namespace:    0,
		Name:         e.Name,
		Experimental: false,
		Info:         &chain_schema.TypeDefinition_Enum{Enum:e2},
	}
	return res
}

func (value EnumValue) Emit() *chain_schema.EnumValue {
	return &chain_schema.EnumValue{
		Name:  value.Name,
		Value: value.Value,
	}
}

