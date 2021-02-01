package store

import "github.com/xsoroton/go-rules/models"

type Store interface {
	GetRoles() models.RolesMap
	GetRolesParentMap() models.RolesParentMap
	GetUsers() models.UsersMap
	GetUsersRoleMap() models.UsersRoleMap
}
