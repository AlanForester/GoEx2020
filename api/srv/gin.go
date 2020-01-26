package srv

import (
	"context"
	"github.com/RaMin0/gin-health-check"
	"github.com/danielkov/gin-helmet"
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/szuecs/gin-glog"
	"/api/models"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type GinFramework struct {
	Engine *gin.Engine
	Srv    *http.Server
}

func (g *GinFramework) Bootstrap(addr string)  {

	g.Srv = &http.Server{
		Addr:    addr,
		Handler: g.Engine,
	}

	go func() {
		// service connections
		if err := g.Srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := g.Srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

func (g *GinFramework) SetupLogger() {
	g.Engine.Use(ginglog.Logger(3 * time.Second))
	//glog.Warning("warning")
	//glog.Error("err")
	//glog.Info("info")
}

func (g *GinFramework) LoadMiddleware() {
	// Security middlewares
	g.Engine.Use(helmet.Default())

	g.Engine.Use(models.CORS())
	// Check health
	// curl -iL -XGET -H "X-Health-Check: 1" http://localhost
	g.Engine.Use(healthcheck.Default())
	//g.Engine.Use(limit.CIDR("172.18.0.0/16"))
}

func (g *GinFramework) HandleRecoveries() {
	// Install nice.Recovery, passing the handler to call after recovery
	g.Engine.Use(nice.Recovery(g.recoveryHandler))
}

func (g *GinFramework) recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(500, err)
}

func NewAPI() *GinFramework {
	model := new(GinFramework)
	model.Engine = gin.New()
	return model
}
