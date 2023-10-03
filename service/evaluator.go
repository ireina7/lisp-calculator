package service

import (
	"fmt"

	"github.com/ireina7/fgo/structs/result"
	"github.com/ireina7/fgo/structs/slice"
	"github.com/ireina7/lisp-calculator/model"
)

type Integer = model.Integer

type Evaluate[Input, Output any] interface {
	Eval(Input) result.Result[Output]
}
type Evaluator struct {
	// interfaces.Traversable[result.ResultKind, slice.SliceKind, model.LispExpr, Integer] `summon:"type"`
}

func NewEvaluator() *Evaluator {
	return &Evaluator{}
}

func (eval *Evaluator) Eval(expr model.LispExpr) result.Result[Integer] {
	switch expr := expr.(type) {
	case model.Integer:
		return result.From(expr)
	case model.Name:
		return result.FromErr[Integer](fmt.Errorf("name is not Integer"))
	case *model.List:
		var f func(Integer, Integer) Integer
		switch expr.Op {
		case "+":
			f = func(a, b Integer) Integer {
				return a + b
			}
		case "-":
			f = func(a, b Integer) Integer {
				return a - b
			}
		case "*":
			f = func(a, b Integer) Integer {
				return a * b
			}
		case "/":
			f = func(a, b Integer) Integer {
				return a / b
			}
		default:
			return result.FromErr[Integer](fmt.Errorf("unknown operator: %v", expr.Op))
		}
		traverse := slice.TraverseResult[model.LispExpr, Integer]{}
		args := traverse.Traverse(
			expr.Args,
			func(e model.LispExpr) result.Result[Integer] {
				return eval.Eval(e)
			},
		)

		return result.AndThen(
			args,
			func(xs slice.Slice[Integer]) result.Result[Integer] {
				var ans Integer
				xs.ForEach(func(i int, x Integer) {
					if i == 0 {
						ans = x
						return
					}
					ans = f(ans, x)
				})
				return result.From(ans)
			},
		)
	default:
		return result.FromErr[Integer](fmt.Errorf("invalid expr type: %#v", expr))
	}
}
