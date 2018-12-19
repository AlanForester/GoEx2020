package main

import (
	"gop/api/models"
	"gop/api/srv"
)

func main() {
	api := srv.NewAPI()
	api.SetupLogger()
	api.HandleRecoveries()
	api.LoadMiddleware()
	//glog.Warning("warning")
	//glog.Error("err")
	models.BootstrapRouting(api)

	api.Bootstrap("127.0.0.1:8080")
}
