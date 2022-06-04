package golang

import "reflect"

type item struct {
	kind     uint
	name     string
	pkg      string
	path     string
	original string
	deps     items
}

type items map[string]item

type alias string

type imports map[string]alias

var itemKind = struct {
	Func   uint
	Struct uint
	String uint
}{
	1,
	2,
	3,
}

type Field struct {
	Id        string
	Kind      reflect.Kind
	TypeName  string
	FieldName string
	PkgPath   string
}

type Method struct {
	Name string
	In   []Field
	Out  []Field
}

type Type struct {
	Id      string
	Kind    reflect.Kind
	Name    string
	String  string
	PkgPath string
	Fields  []Field
	Methods []Method
}
