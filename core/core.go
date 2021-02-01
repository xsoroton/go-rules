package core

import (
	"log"

	"github.com/xsoroton/go-rules/models"
	s "github.com/xsoroton/go-rules/store"
	"github.com/xsoroton/go-rules/util"
)

var store s.Store

func init() {
	store = s.NewStoreFromJsonFiles(
		util.EnvString("RULES_JSON_FILE", "./jsonData/rules.json"),
		util.EnvString("USERS_JSON_FILE", "./jsonData/users.json"),
	)
}

// todo: add LRU Cache with ttl to boost performance

func GetSubOrdinates(userId int) (users models.Users) {
	user, found := store.GetUsers()[userId]
	if !found {
		return // User not exist
	}
	// Get all subordinates roles
	subordinatesIds := getSubordinatesRoles(user.Role)
	// get users by roles ids
	usersRoleMap := store.GetUsersRoleMap()
	for _, id := range subordinatesIds {
		v, ok := usersRoleMap[id]
		if ok {
			users = append(users, v...)
		}
	}
	return
}

func getSubordinatesRoles(roleId int) (subordinatesRolesIds []int) {
	roles := store.GetRoles()
	role, ok := roles[roleId]
	if !ok {
		log.Fatal("role id is not fund")
	}
	var lookUpSubs = []int{role.ID}
	parentRoleMap := store.GetRolesParentMap()
	// find ALL subordinates in tree
	for len(lookUpSubs) > 0 {
		var next []int
		for _, kidId := range lookUpSubs {
			kidsRoles, ok := parentRoleMap[kidId]
			if ok {
				for _, r := range kidsRoles {
					next = append(next, r.ID)
					subordinatesRolesIds = append(subordinatesRolesIds, r.ID)
				}
			}
		}
		lookUpSubs = next
	}
	return
}

// ... not need it, because tree is sorted ;)
func isAdmin(role models.Role) bool {
	return role.ID == models.SystemAdminParenID
}
