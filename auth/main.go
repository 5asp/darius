package main

import (
	"net/http"
	"os"

	"github.com/aszeta/darius/auth/account"

	"github.com/go-kit/kit/log"
)

func main() {
	config := account.NewConfig(".")
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "listen", config.App.Addr, "caller", log.DefaultCaller)

	r := account.NewHttpServer(account.NewService(), logger)
	logger.Log("msg", "HTTP", "addr", config.App.Addr)
	logger.Log("err", http.ListenAndServe(config.App.Addr, r))
}
