# # Development and Production commands. # #
# Production build is default command: make
run: vite build
	@./bin/server

dev: build
	@./bin/server


# # Scripts commands # #
init:
	@chmod +x ./scripts/bash/init.sh
	@./scripts/bash/init.sh
e:
	@chmod +x ./scripts/bash/environment.sh
	@./scripts/bash/environment.sh
p: 
	@flyctl launch
	@flyctl deploy


# # Utils commands # #
build:
	@go build -o bin/server

vite:
	@npm run build

hot:
	@chmod +x ./scripts/bash/hotReloadGo.sh
	@./scripts/bash/hotReloadGo.sh

clean:
	@go mod tidy