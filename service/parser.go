package service

import (
	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/result"
	"github.com/ireina7/lisp-calculator/model"
)

type Parser[A any] interface {
	Parse(string) result.Result[A]
}

type lispParser struct {
	interfaces.Logger
}

func (parser *lispParser) Parse(src string) result.Result[model.LispExpr] {

	return result.FromErr[model.LispExpr](nil)
}
