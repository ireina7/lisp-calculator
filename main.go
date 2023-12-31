package main

import (
	"github.com/ireina7/lisp-calculator/inject"
	"github.com/ireina7/summoner"
)

func main() {
	inject.Instances()

	server := &Server{
		// httpServer: http.Server{
		// 	Addr: fmt.Sprintf(":%d", 8080),
		// },
	}
	err := summoner.Inject(server)
	if err != nil {
		panic(err)
	}

	err = server.Serve(8080)
	if err != nil {
		panic(err)
	}
}
