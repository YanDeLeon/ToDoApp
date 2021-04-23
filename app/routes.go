package app

import (
	base "todoapp"
	"todoapp/app/actions"
	"todoapp/app/middleware"

	"github.com/gobuffalo/buffalo"
	paramlogger "github.com/gobuffalo/mw-paramlogger"
)

// SetRoutes for the application
func setRoutes(root *buffalo.App) {
	root.Use(middleware.Transaction)
	root.Use(paramlogger.ParameterLogger)
	root.Use(middleware.CSRF)

	root.GET("/", actions.List)
	root.ServeFiles("/", base.Assets)

	root.GET("/task/new", actions.New)
	root.PUT("/task/create", actions.Create)
	root.GET("/task/{id}/show", actions.Show)
	root.GET("/task/{id}/edit", actions.Edit)
	root.PUT("/task/{id}/update", actions.Update)
	root.PUT("/task/{id}/update-status/", actions.ChangeStatus)
	root.DELETE("/task/{id}/delete/", actions.Delete)
}
