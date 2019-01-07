build :
	docker build -t go_course:1.0 .
run_server :
	docker run --name go_course go_course:1.0
check :
	go vet ./codebase/*
	golint ./codebase/*
	goimports ./codebase/*
test :
	go test ./codebase/*  -coverprofile=cover.out
client	:
	docker exec -it go_course /go/src/kv_storage/codebase/connect
