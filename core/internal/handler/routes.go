// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "github.com/pjimming/zustacm/core/internal/handler/auth"
	basic "github.com/pjimming/zustacm/core/internal/handler/basic"
	permission "github.com/pjimming/zustacm/core/internal/handler/permission"
	resource "github.com/pjimming/zustacm/core/internal/handler/resource"
	user "github.com/pjimming/zustacm/core/internal/handler/user"
	"github.com/pjimming/zustacm/core/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ping",
				Handler: basic.PingHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/auth/login",
				Handler: auth.AuthLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/auth/captcha",
				Handler: auth.GetAuthCaptchaHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/auth/logout",
					Handler: auth.AuthLogoutHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/user",
					Handler: user.AddUserHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/user",
					Handler: user.GetUserPageHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/resource",
					Handler: resource.AddResourceHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/resource/:id",
					Handler: resource.SaveResourceHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/resource/menu/tree",
					Handler: resource.GetResourceMenuTreeHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/v1/resource/:id",
					Handler: resource.DeleteResourceHandler(serverCtx),
				},
			}...,
		),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuth},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/role/permissions/tree",
					Handler: permission.GetRolePermissionTreeHandler(serverCtx),
				},
			}...,
		),
	)
}
