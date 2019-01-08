# Server-client solution for storing KV data

**Requirements:** \
Docker (>=18.06.1-ce) \
go (>=1.11.2) 

**Go packages:** \
goimports (go get golang.org/x/tools/cmd/goimports) \
golint (go get github.com/golang/lint/golint)

**Steps to run server:** \
make build -> create docker container with codebase \
make run_server -> start server inside container

**Steps to run client:** \
Run server (take a look "Steps to run server" section) \
make client

**Run tests:** \
make test 

File with information about tests coverage (cover.out) will be generated and put into project directory

**Run code validators:** \
make check
