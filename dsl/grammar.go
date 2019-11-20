package dsl

import (
	"github.com/alecthomas/participle"
)

type Module struct {
	Decls []*Decl `@@*`
}

type Decl struct {
	Struct *Struct ` @@`
	Enum   *Enum   `| @@`
}

type Struct struct {
	Name   string   `"struct" @Ident`
	Fields []*Field `"{" { @@ } "}"`
}

type Field struct {
	Name string   `@Ident`
	Type *TypeRef `":" @@`
}

type TypeRef struct {
	Type string `@Ident`
	Many bool   `[ @"@" ]`
}

type Enum struct {
	Name   string      `"enum" @Ident`
	Values []EnumValue `"{" { @@ } "}"`
}

type EnumValue struct {
	Name  string `@Ident`
	Value int32
}

var (
	Parser = participle.MustBuild(&Module{})
)
