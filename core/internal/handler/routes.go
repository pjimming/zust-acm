// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	auth "github.com/pjimming/zustacm/core/internal/handler/auth"
	basic "github.com/pjimming/zustacm/core/internal/handler/basic"
	codeforces "github.com/pjimming/zustacm/core/internal/handler/codeforces"
	competition "github.com/pjimming/zustacm/core/internal/handler/competition"
	permission "github.com/pjimming/zustacm/core/internal/handler/permission"
	record "github.com/pjimming/zustacm/core/internal/handler/record"
	resource "github.com/pjimming/zustacm/core/internal/handler/resource"
	role "github.com/pjimming/zustacm/core/internal/handler/role"
	sysdict "github.com/pjimming/zustacm/core/internal/handler/sysdict"
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
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/assign/roles",
					Handler: user.AssignRolesHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/user/detail",
					Handler: user.GetUserDetailHandler(serverCtx),
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
					Path:    "/api/v1/role",
					Handler: role.AddRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/role/page",
					Handler: role.GetRolePageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/role/all",
					Handler: role.GetRoleAllHandler(serverCtx),
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
					Path:    "/api/v1/codeforces/sync/user/:id",
					Handler: codeforces.SyncUserCfHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/api/v1/codeforces/sync/all",
					Handler: codeforces.SyncCfAllHandler(serverCtx),
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
					Path:    "/api/v1/competition",
					Handler: competition.AddCompetitionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/competition",
					Handler: competition.GetCompetitionHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/competition/:id",
					Handler: competition.UpdateCompetitionHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/v1/competition/:id",
					Handler: competition.DeleteCompetitionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/competition/all",
					Handler: competition.GetCompetitionAllHandler(serverCtx),
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
					Path:    "/api/v1/record",
					Handler: record.AddRecordHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/record",
					Handler: record.GetRecordHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/record/:id",
					Handler: record.UpdateRecordHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/v1/record/:id",
					Handler: record.DeleteRecordHandler(serverCtx),
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
					Path:    "/api/v1/sysdict",
					Handler: sysdict.AddSysDictHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/api/v1/sysdict",
					Handler: sysdict.GetSysDictHandler(serverCtx),
				},
				{
					Method:  http.MethodPut,
					Path:    "/api/v1/sysdict/:id",
					Handler: sysdict.UpdateSysDictHandler(serverCtx),
				},
				{
					Method:  http.MethodDelete,
					Path:    "/api/v1/sysdict/:id",
					Handler: sysdict.DeleteSysDictHandler(serverCtx),
				},
			}...,
		),
	)
}
