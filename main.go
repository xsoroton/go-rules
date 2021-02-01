package main

import (
	"encoding/json"
	"fmt"

	"github.com/xsoroton/go-rules/core"
)

func main() {

	// user ids
	var ids = []int{1, 2, 3, 4, 5}

	for _, id := range ids {
		users := core.GetSubOrdinates(id)
		b, _ := json.Marshal(users)
		fmt.Printf("userId %d: %s\n", id, b)
	}
}
