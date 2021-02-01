`GetSubOrdinates()` located in `core/core.go` 

# RUN

### Local Run

To run
```bash
make run
```

###Build and run in Docker container
```bash
make run-in-docker
```

results
```
userId 1: [{"Id":4,"Name":"Mary Manager","Role":2},{"Id":3,"Name":"Sam Supervisor","Role":3},{"Id":2,"Name":"Emily Employee","Role":4},{"Id":5,"Name":"Steve Trainer","Role":5}]
userId 2: null
userId 3: [{"Id":2,"Name":"Emily Employee","Role":4},{"Id":5,"Name":"Steve Trainer","Role":5}]
userId 4: [{"Id":3,"Name":"Sam Supervisor","Role":3},{"Id":2,"Name":"Emily Employee","Role":4},{"Id":5,"Name":"Steve Trainer","Role":5}]
userId 5: null
```

# Test

### Run Test

To run test
```bash
make test
```

```       
go test ./... -v
?       github.com/xsoroton/go-rules    [no test files]
=== RUN   TestGetSubOrdinates

  Test get subordinates 
    User ID 1 ✔
    User ID 2 ✔
    User ID 3 ✔
    User ID 4 ✔
    User ID 5 ✔


5 total assertions

--- PASS: TestGetSubOrdinates (0.00s)
PASS
ok      github.com/xsoroton/go-rules/core       (cached)
?       github.com/xsoroton/go-rules/models     [no test files]
=== RUN   TestStoreFromJsonFiles

  Test Store from json files 
    GetRoles ✔
    GetRolesParentMap ✔
    GetUsers ✔
    GetUsersRoleMap ✔


4 total assertions

--- PASS: TestStoreFromJsonFiles (0.00s)
PASS
ok      github.com/xsoroton/go-rules/store      (cached)
```

### Run test cover
To calculate test cover
```bash
make test-cover
```

```
go test -covermode=count -coverprofile=coverage.out ./...
?       github.com/xsoroton/go-rules    [no test files]
ok      github.com/xsoroton/go-rules/core       0.003s  coverage: 89.3% of statements
?       github.com/xsoroton/go-rules/models     [no test files]
ok      github.com/xsoroton/go-rules/store      0.002s  coverage: 97.1% of statements
?       github.com/xsoroton/go-rules/util       [no test files]
```
