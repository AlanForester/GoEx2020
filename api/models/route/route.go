package route

import (
	"github.com/gin-gonic/gin"
	"source.cloud.google.com/onemo-api//api/interfaces"
)

type Route struct {
	Protected bool
	fullPath string
	RouteGroup *gin.RouterGroup
	Handler func(ctx *gin.Context)
}

func (r *Route) SetHandler(handler func(ctx *gin.Context)) {
	r.Handler = handler
}

func (r *Route) GetHandler() func(ctx *gin.Context) {
	return r.Handler
}

func (r *Route) LoadRoutes() []interfaces.Routed {
	return nil
}

func (r *Route) SetApiGroup(group *gin.RouterGroup) {
	r.RouteGroup = group
}

func (r *Route) GetApiGroup() *gin.RouterGroup {
	return r.RouteGroup
}

func (r *Route) SetFullPath(path string) {
	r.fullPath = path
}

func (r *Route) GetFullPath() string {
	return r.fullPath
}

func (r *Route) Path() string {
	return ""
}

func (r *Route) Get(ctx *gin.Context)  {
}

func (r *Route) Post(ctx *gin.Context)  {
}

func (r *Route) Put(ctx *gin.Context)  {
}

func (r *Route) Patch(ctx *gin.Context)  {
}

func (r *Route) Delete(ctx *gin.Context)  {
}