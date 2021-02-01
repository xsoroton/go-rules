package store

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/xsoroton/go-rules/models"
)

var (
	store Store

	expectedRoles = models.RolesMap{
		1: models.Role{ID: 1, Name: "System Administrator", Parent: 0},
		2: models.Role{ID: 2, Name: "Location Manager", Parent: 1},
		3: models.Role{ID: 3, Name: "Supervisor", Parent: 2},
		4: models.Role{ID: 4, Name: "Employee", Parent: 3},
		5: models.Role{ID: 5, Name: "Trainer", Parent: 3},
	}

	expectedRolesParentMap = models.RolesParentMap{
		0: []models.Role{{ID: 1, Name: "System Administrator", Parent: 0}},
		1: []models.Role{{ID: 2, Name: "Location Manager", Parent: 1}},
		2: []models.Role{{ID: 3, Name: "Supervisor", Parent: 2}},
		3: []models.Role{{ID: 4, Name: "Employee", Parent: 3}, {ID: 5, Name: "Trainer", Parent: 3}},
	}

	expectedUsers = models.UsersMap{
		1: models.User{ID: 1, Name: "Adam Admin", Role: 1},
		2: models.User{ID: 2, Name: "Emily Employee", Role: 4},
		3: models.User{ID: 3, Name: "Sam Supervisor", Role: 3},
		4: models.User{ID: 4, Name: "Mary Manager", Role: 2},
		5: models.User{ID: 5, Name: "Steve Trainer", Role: 5},
	}

	expectedUsersRolesMap = models.UsersRoleMap{
		1: []models.User{{ID: 1, Name: "Adam Admin", Role: 1}},
		2: []models.User{{ID: 4, Name: "Mary Manager", Role: 2}},
		3: []models.User{{ID: 3, Name: "Sam Supervisor", Role: 3}},
		4: []models.User{{ID: 2, Name: "Emily Employee", Role: 4}},
		5: []models.User{{ID: 5, Name: "Steve Trainer", Role: 5}},
	}
)

func init() {
	store = NewStoreFromJsonFiles("./../jsonData/rules.json", "./../jsonData/users.json")
}

func TestStoreFromJsonFiles(t *testing.T) {
	Convey("Test Store from json files", t, func() {
		Convey("GetRoles", func() {
			roles := store.GetRoles()
			So(roles, ShouldResemble, expectedRoles)
		})
		Convey("GetRolesParentMap", func() {
			roles := store.GetRolesParentMap()
			So(roles, ShouldResemble, expectedRolesParentMap)
		})
		Convey("GetUsers", func() {
			users := store.GetUsers()
			So(users, ShouldResemble, expectedUsers)
		})
		Convey("GetUsersRoleMap", func() {
			users := store.GetUsersRoleMap()
			So(users, ShouldResemble, expectedUsersRolesMap)
		})
	})
}
