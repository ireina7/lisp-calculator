package main

import (
	"context"

	"github.com/ireina7/lisp-calculator/inject"
	"github.com/ireina7/lisp-calculator/service"
	"github.com/ireina7/summoner"
)

func main() {
	inject.Instances()

	service := &service.Service{
		Ctx: context.Background(),
	}
	err := summoner.Inject(service)
	if err != nil {
		panic(err)
	}
	err = service.Serve()
	if err != nil {
		panic(err)
	}
}
