package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/ireina7/fgo/interfaces"
	"github.com/ireina7/fgo/structs/result"
	"github.com/ireina7/fgo/types"
	"github.com/ireina7/lisp-calculator/model"
	"github.com/ireina7/lisp-calculator/service"
)

type Server struct {
	Ctx               context.Context `summon:"type"`
	interfaces.Logger `summon:"type"`
	// httpServer        http.Server `summon:"none"`
	service.Service
}

func (server *Server) Serve(port int) error {
	http.HandleFunc("/lisp/eval", func(w http.ResponseWriter, req *http.Request) {
		server.Debug("req: %+v", req)
		reader := req.Body
		data, err := io.ReadAll(reader)
		if err != nil {
			server.Error("io.ReadAll err: %+v", err)
			return
		}
		src := string(data)
		ans := server.Eval(src)
		result.MapErr(ans, func(err error) error {
			server.Error("service.Eval err: %+v", err)
			return err
		})
		result.Map(ans, func(i model.Integer) types.Unit {
			_, err := w.Write([]byte(strconv.FormatInt(int64(i), 10)))
			if err != nil {
				server.Error("strconv.FormatInt err: %+v", err)
				return types.MakeUnit()
			}
			return types.MakeUnit()
		})
		server.Info("ok...")
	})
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
