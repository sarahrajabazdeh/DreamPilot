.PHONY: docs
printf = @printf "%s\t\t%s\n"

help:
	@echo -e "Commands available:\n"
	$(printf) "run" "execute the app"
	$(printf) "build" "build the app in an executable file"

run: 
	go run main.go

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-app .	

docker:

