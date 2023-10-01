package service_test

import (
	"context"
	"testing"

	"github.com/ireina7/lisp-calculator/inject"
	"github.com/ireina7/lisp-calculator/model"
	"github.com/ireina7/lisp-calculator/service"
	"github.com/ireina7/summoner"
)

func TestEval(t *testing.T) {
	inject.Inject()

	service := &service.Service{
		Ctx: context.Background(),
	}
	err := summoner.Inject(service)
	if err != nil {
		panic(err)
	}

	expr := model.List{
		Op: "+",
		Args: []model.LispExpr{
			&model.List{Op: "*", Args: []model.LispExpr{model.Integer(2), model.Integer(3)}},
			model.Integer(1),
		},
	}
	ans := service.Evaluator.Eval(&expr)
	t.Logf("ans: %#v", ans)
}
