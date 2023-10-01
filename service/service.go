package service

import (
	"context"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/lisp-calculator/model"
)

type Service struct {
	Ctx                    context.Context
	model.Version          `summon:"type"`
	interfaces.Logger      `summon:"type"`
	Parser[model.LispExpr] `summon:"type"`
	Evaluator
}

func (service *Service) Serve() error {

	return nil
}

// func (service *Service) Eval(src string) result.Result[model.Integer] {
// 	expr := service.Parse(src)
// 	// service.Info("Parsed %#v", expr)
// 	if result.IsErr(expr) {
// 		err := result.GetErr(expr)
// 		// service.Error("service.Parse err: %+v", err)
// 		return result.FromErr[model.Integer](err)
// 	}
// 	return result.FlatMap(expr, func(expr model.LispExpr) result.Result[model.Integer] {
// 		return service.Evaluator.Eval(expr)
// 	})
// }

// func (service *Service) Ver() Version {
// 	return service.Version
// }