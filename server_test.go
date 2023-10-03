package main

import (
	"testing"

	"github.com/ireina7/lisp-calculator/inject"
	"github.com/ireina7/summoner"
)

func TestServer(t *testing.T) {
	inject.Instances()

	server := &Server{}
	err := summoner.Inject(server)
	if err != nil {
		panic(err)
	}

	src := "(* 2 4)"
	ans := server.Eval(src)
	t.Logf("ans: %#v", ans)
}
