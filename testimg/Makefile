
.PHONY: default
default: main.go Dockerfile 


.PHONY: main.go
main.go: 
	env GOOS=linux GOARCH=amd64 go build -o test ./main.go 
	chmod a+x ./test


.PHONY: Dockerfile
Dockerfile:
	docker build -t testdep:lukeg . 

