package model

type Integer int

func (Integer) lispExpr() {}

type Name string

func (Name) lispExpr() {}

type List struct {
	Op   Name
	Args []LispExpr
}

func (*List) lispExpr() {}

type LispExpr interface {
	lispExpr()
}
