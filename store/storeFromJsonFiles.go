package store

import (
	"io/ioutil"

	"github.com/xsoroton/go-rules/models"
	"github.com/xsoroton/go-rules/util"
)

// Json File locations
type dataFromFile struct {
	rulesFile string
	usersFile string
}

// Implement Store interface from local json files
func NewStoreFromJsonFiles(rulesFile, usersFile string) Store {
	return &dataFromFile{
		rulesFile: rulesFile,
		usersFile: usersFile,
	}
}

// ROLES

func (d *dataFromFile) GetRoles() models.RolesMap {
	roles := getRoles(d)
	rolesMap := make(models.RolesMap, len(roles))
	for _, role := range roles {
		rolesMap[role.ID] = role
	}
	return rolesMap
}

func (d *dataFromFile) GetRolesParentMap() models.RolesParentMap {
	roles := getRoles(d)
	rolesParentMap := make(models.RolesParentMap)
	for _, role := range roles {
		mRoles, ok := rolesParentMap[role.Parent]
		if !ok {
			rolesParentMap[role.Parent] = []models.Role{role}
		} else {
			rolesParentMap[role.Parent] = append(mRoles, role)
		}
	}
	return rolesParentMap
}

func getRoles(d *dataFromFile) (roles models.Roles) {
	data, err := ioutil.ReadFile(d.rulesFile)
	util.FatalOnError(err)
	util.FatalUnlessUnmarshal(data, &roles)
	return
}

// USER

func (d *dataFromFile) GetUsers() models.UsersMap {
	users := getUsers(d)
	usersMap := make(models.UsersMap, len(users))
	for _, user := range users {
		usersMap[user.ID] = user
	}
	return usersMap
}

func (d *dataFromFile) GetUsersRoleMap() models.UsersRoleMap {
	users := getUsers(d)
	usersRoleMap := make(models.UsersRoleMap)
	for _, user := range users {
		mUsers, ok := usersRoleMap[user.Role]
		if !ok {
			usersRoleMap[user.Role] = []models.User{user}
		} else {
			usersRoleMap[user.Role] = append(mUsers, user)
		}
	}
	return usersRoleMap
}

func getUsers(d *dataFromFile) (users models.Users) {
	data, err := ioutil.ReadFile(d.usersFile)
	util.FatalOnError(err)
	util.FatalUnlessUnmarshal(data, &users)
	return
}
