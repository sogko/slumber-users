package users

import (
	"github.com/sogko/slumber/domain"
	"strings"
)

const (
	ListUsers      = "ListUsers"
	CountUsers     = "CountUsers"
	GetUser        = "GetUser"
	CreateUser     = "CreateUser"
	UpdateUsers    = "UpdateUsers"
	DeleteAllUsers = "DeleteAllUsers"
	ConfirmUser    = "ConfirmUser"
	UpdateUser     = "UpdateUser"
	DeleteUser     = "DeleteUser"
)
const defaultBasePath = "/api/users"

func (resource *Resource) generateRoutes(basePath string) *domain.Routes {
	if basePath == "" {
		basePath = defaultBasePath
	}
	var baseRoutes = domain.Routes{
		domain.Route{
			Name:           ListUsers,
			Method:         "GET",
			Pattern:        "/api/users",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleListUsers_v0,
			},
			ACLHandler: resource.HandleListUsersACL,
		},
		domain.Route{
			Name:           CountUsers,
			Method:         "GET",
			Pattern:        "/api/users/count",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleCountUsers_v0,
			},
			ACLHandler: resource.HandleCountUsersACL,
		},
		domain.Route{
			Name:           CreateUser,
			Method:         "POST",
			Pattern:        "/api/users",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleCreateUser_v0,
			},
			ACLHandler: resource.HandleCreateUserACL,
		},
		domain.Route{
			Name:           UpdateUsers,
			Method:         "PUT",
			Pattern:        "/api/users",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleUpdateUsers_v0,
			},
			ACLHandler: resource.HandleUpdateUsersACL,
		},
		domain.Route{
			Name:           DeleteAllUsers,
			Method:         "DELETE",
			Pattern:        "/api/users",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleDeleteAllUsers_v0,
			},
			ACLHandler: resource.HandleDeleteAllUsersACL,
		},
		domain.Route{
			Name:           GetUser,
			Method:         "GET",
			Pattern:        "/api/users/{id}",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleGetUser_v0,
			},
			ACLHandler: resource.HandleGetUserACL,
		},
		/*
			Method for email confirmation has to be GET because
			link to confirm email has to be click-able from email content
			(You can't add a POST/PUT body)
		*/
		domain.Route{
			Name:           ConfirmUser,
			Method:         "GET",
			Pattern:        "/api/users/{id}/confirm",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleConfirmUser_v0,
			},
			ACLHandler: resource.HandleConfirmUserACL,
		},
		domain.Route{
			Name:           UpdateUser,
			Method:         "PUT",
			Pattern:        "/api/users/{id}",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleUpdateUser_v0,
			},
			ACLHandler: resource.HandleUpdateUserACL,
		},
		domain.Route{
			Name:           DeleteUser,
			Method:         "DELETE",
			Pattern:        "/api/users/{id}",
			DefaultVersion: "0.0",
			RouteHandlers: domain.RouteHandlers{
				"0.0": resource.HandleDeleteUser_v0,
			},
			ACLHandler: resource.HandleDeleteUserACL,
		},
	}

	routes := domain.Routes{}

	for _, route := range baseRoutes {
		r := domain.Route{
			Name:           route.Name,
			Method:         route.Method,
			Pattern:        strings.Replace(route.Pattern, defaultBasePath, basePath, -1),
			DefaultVersion: route.DefaultVersion,
			RouteHandlers:  route.RouteHandlers,
			ACLHandler:     route.ACLHandler,
		}
		routes = routes.Append(&domain.Routes{r})
	}
	resource.routes = &routes
	return resource.routes
}
