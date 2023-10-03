package service_test

import (
	"fmt"
	"testing"

	"github.com/ireina7/fgo/structs/result"
	"github.com/ireina7/lisp-calculator/model"
	"github.com/ireina7/lisp-calculator/service"
)

func TestParser(t *testing.T) {
	parser := service.NewLispParser()
	res := parser.Parse("(/ (* 1 4) 2)")
	t.Logf("%#v", result.Map(res, func(e model.LispExpr) string {
		return fmt.Sprintf("%#v", *(e.(*model.List)))
	}))
}
