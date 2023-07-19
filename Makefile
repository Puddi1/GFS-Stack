include .env
export

run: vite build
	@./bin/server

dev: build
	@./bin/server

e:

# @bash -c 'if [[ \"$${DEVELOPMENT}\" == \"true\" ]]; then echo \"true\"; fi'
# e to list envs
# change: awk -F '=' '/^DEVELOPMENT/ { $2 = \"false\" } 1' .env > temp && mv temp .env

build:
	@go build -o bin/server

vite:
	@npm run build

# prob not necessary
hotreload-run:
	@ls * | entr -r make
# reload only backend files
hotreload-dev:
	@ls

clean:
	@go mod tidy