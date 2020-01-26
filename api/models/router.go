package models

import (
	"github.com/gin-gonic/gin"
	"source.cloud.google.com/onemo-api//api/controllers"
	"source.cloud.google.com/onemo-api//api/interfaces"
	"source.cloud.google.com/onemo-api//api/srv"
	"gopkg.in/appleboy/gin-jwt.v2"
	"log"
)

type Router struct {
	Api    *gin.Engine
	Auth   *Auth
	routes map[string]interfaces.Routed
}

func (r *Router) Add(path string, route interfaces.Routed) {
	if r.routes == nil {
		r.routes = make(map[string]interfaces.Routed)
	}

	if r.routes[path] == nil {
		r.routes[path] = route
		var group= &r.Api.RouterGroup
		if route.GetApiGroup() != nil {
			group = route.GetApiGroup()
		}

		if route.Get != nil {
			group.GET(path, route.Get)
		}
		if route.Post != nil {
			group.POST(path, route.Post)
		}
		if route.Patch != nil {
			group.PATCH(path, route.Patch)
		}
		if route.Put != nil {
			group.PUT(path, route.Put)
		}
		if route.Delete != nil {
			group.DELETE(path, route.Delete)
		}
	}
}

func (r *Router) makeAuthRoutes(auth *gin.RouterGroup) gin.IRoutes {

	mdw := r.Auth.GetMiddleware()
	auth.POST("/login", mdw.LoginHandler)

	r.Api.NoRoute(mdw.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth.Use(mdw.MiddlewareFunc())
	auth.GET("/refresh_token", mdw.RefreshHandler)
	auth.GET("/hello", helloHandler)
	return auth
}

func (r *Router) MergeRoutes() {
	root := r.Api.Group("/")
	api := root.Group("/api")
	v1 := api.Group("/v1")
	auth := v1.Group("/auth")
	r.makeAuthRoutes(auth)

	//rootCtrl := new(tree.ApiController)
	rootCtrl := new(controllers.ApiController)
	rootCtrl.SetApiGroup(root)
	r.BuildRoutesTree(rootCtrl)
	//r.routes[string] = route

}

func helloHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)

	c.JSON(200, gin.H{
		"userID": claims["id"],
		"claims": claims,
		"text":   "Hello World.",
	})

}

func (r *Router) BuildRoutesTree(ctrl interfaces.Routed) {

	r.Add(ctrl.Path(),ctrl)

	if childs := ctrl.LoadRoutes(); childs != nil {
		group := ctrl.GetApiGroup()
		for _, child := range childs {
			path := child.Path()
			if path != "" {
				child.SetFullPath(ctrl.Path())
				if group != nil {
					child.SetApiGroup(group.Group(ctrl.Path()))
				}
				r.BuildRoutesTree(child)
			}
		}
	}
	//r.routes[string] = route
}

func (r *Router) BootstrapAppRouting() {
	// Install nice.Recovery, passing the handler to call after recovery

	r.MergeRoutes()

}

func BootstrapRouting(api *srv.GinFramework) {
	model := new(Router)
	model.routes = make(map[string]interfaces.Routed)
	model.Api = api.Engine
	model.Auth = NewAuthService()
	model.BootstrapAppRouting()
}
