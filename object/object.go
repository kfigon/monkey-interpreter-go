package object

import "strconv"

type ObjectType string

const (
	INTEGER ObjectType = "INTEGER"
	BOOLEAN ObjectType = "BOOLEAN"
	NULL    ObjectType = "NULL"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int
}

func (i *Integer) Type() ObjectType {
	return INTEGER
}

func (i *Integer) Inspect() string {
	return strconv.Itoa(i.Value)
}


type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN
}

func (b *Boolean) Inspect() string {
	return strconv.FormatBool(b.Value)
}

type Null struct {}

func (n *Null) Type() ObjectType {
	return NULL
}

func (n *Null) Inspect() string {
	return "null"
}