package core

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/xsoroton/go-rules/models"
	s "github.com/xsoroton/go-rules/store"
)

func init() {
	store = s.NewStoreFromJsonFiles(
		"./../jsonData/rules.json",
		"./../jsonData/users.json",
	)
}

var (
	expectedResultUser1 = models.Users{
		models.User{ID: 4, Name: "Mary Manager", Role: 2},
		models.User{ID: 3, Name: "Sam Supervisor", Role: 3},
		models.User{ID: 2, Name: "Emily Employee", Role: 4},
		models.User{ID: 5, Name: "Steve Trainer", Role: 5},
	}

	expectedResultUser3 = models.Users{
		models.User{ID: 2, Name: "Emily Employee", Role: 4},
		models.User{ID: 5, Name: "Steve Trainer", Role: 5},
	}

	expectedResultUser4 = models.Users{
		models.User{ID: 3, Name: "Sam Supervisor", Role: 3},
		models.User{ID: 2, Name: "Emily Employee", Role: 4},
		models.User{ID: 5, Name: "Steve Trainer", Role: 5},
	}
)

func TestGetSubOrdinates(t *testing.T) {
	Convey("Test get subordinates", t, func() {
		Convey("User ID 1", func() {
			So(GetSubOrdinates(1), ShouldResemble, expectedResultUser1)
		})
		Convey("User ID 2", func() {
			So(GetSubOrdinates(2), ShouldBeNil)
		})
		Convey("User ID 3", func() {
			So(GetSubOrdinates(3), ShouldResemble, expectedResultUser3)
		})
		Convey("User ID 4", func() {
			So(GetSubOrdinates(4), ShouldResemble, expectedResultUser4)
		})
		Convey("User ID 5", func() {
			So(GetSubOrdinates(5), ShouldBeNil)
		})
	})
}
