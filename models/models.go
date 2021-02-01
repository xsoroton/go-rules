package models

// NOTE: use of maps will be much faster, specially with big data, remove needs in loops.

const SystemAdminParenID = 0

type Roles []Role

// map[<RoleID>]User
type RolesMap map[int]Role

// map[<ParentId>][]Role
type RolesParentMap map[int][]Role

type Role struct {
	ID     int    `json:"Id"`
	Name   string `json:"Name"`
	Parent int    `json:"Parent"`
}

type Users []User

// map[<userID>]User
type UsersMap map[int]User

// map[<RoleId>][]User
type UsersRoleMap map[int][]User

type User struct {
	ID   int    `json:"Id"`
	Name string `json:"Name"`
	Role int    `json:"Role"`
}
