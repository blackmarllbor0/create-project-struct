PROJECT_NAME = create-project-struct
PROJECT_PATH = cmd/$(PROJECT_NAME).go
PROGRAM_NAME = cps

.PHONE: run
run:
	go run $(PROJECT_PATH)


.PHONE: build
build:
	go build -o $(PROGRAM_NAME) $(PROJECT_PATH)


.PHONY: tests
tests:
	go test ./...


.PHONE: exec
exec:
	make build && sudo mv cps /usr/local/bin/cps


.PHONE: rm-program
rm-program:
	sudo rm /usr/local/bin/cps


.PHONE: test
test:
	sudo make rm-program && make exec


.PHONE: clear
clear:
	sh ./script/clear.sh