PROJECT_NAME = create-project-struct
PROJECT_PATH = cmd/$(PROJECT_NAME).go
PROGRAM_NAME = cps

run:
	go run $(PROJECT_PATH)

build:
	go build -o $(PROGRAM_NAME) $(PROJECT_PATH)

exec:
	make build && sudo mv cps /usr/local/bin/cps

rm-program:
	sudo rm /usr/local/bin/cps

test:
	sudo make rm-program && make exec