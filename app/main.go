package main

import (
	"github.com/KyTama/KyGo/config"
	"github.com/KyTama/KyGo/routes"
	"github.com/KyTama/KyGo/utils"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	config.Listen("localhost:6666", routes.NewRoutes())
}

func Ping(w http.ResponseWriter, r *http.Request) {
	now := utils.Now(7, "2006-01-02 15:04:05")

	utils.HTTPResponse(w, "OK", nil, http.StatusOK, now, r.Body)
	return
}