package inject

import (
	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/lisp-calculator/model"
	"github.com/ireina7/lisp-calculator/service"
	"github.com/ireina7/summoner"
)

func Inject() {
	summoner.Given[model.Version](model.Version("0.1.0"))
	summoner.Given[service.Parser[model.LispExpr]](service.NewLispParser())
	summoner.Given[interfaces.Logger](interfaces.NewPreludeLogger(nil))
	summoner.Given[service.Evaluate[model.LispExpr, model.Integer]](service.NewEvaluator())
	// summoner.Given[interfaces.Traversable[result.ResultKind, slice.SliceKind, model.LispExpr, model.Integer]](&SliceTraversable[result.ResultKind, model.LispExpr, model.Integer]{
	// 	Functor:  slice.NewSliceFunctor[model.LispExpr, types.HKT[result.ResultKind, model.Integer]](),
	// 	FunctorF: result.NewResultFunctor[model.Integer, types.Unit](),
	// 	Pure:     result.NewResultApplicative[types.HKT[slice.SliceKind, model.Integer]](),
	// })
}

// type SliceTraversable[F_, A, B any] struct {
// 	Functor  functor.Functor[slice.SliceKind, A, types.HKT[F_, B]]
// 	FunctorF functor.Functor[F_, B, types.Unit]
// 	Pure     interfaces.Applicative[F_, types.HKT[slice.SliceKind, B]]
// }

// func (self *SliceTraversable[F_, A, B]) Traverse(
// 	xs types.HKT[slice.SliceKind, A],
// 	f func(A) types.HKT[F_, B],
// ) types.HKT[F_, types.HKT[slice.SliceKind, B]] {

// 	ys := self.Functor.Fmap(xs, func(a A) types.HKT[F_, B] {
// 		return f(a)
// 	}).(slice.Slice[types.HKT[F_, B]])
// 	zs := slice.Empty[B]()
// 	for _, y := range ys {
// 		self.FunctorF.Fmap(y, func(b B) types.Unit {
// 			zs = zs.Append(b)
// 			return types.MakeUnit()
// 		})
// 	}
// 	return self.Pure.Pure(zs)
// }
