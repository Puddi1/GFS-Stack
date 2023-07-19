run: vite build
	@./bin/server

dev: build
	@./bin/server

build:
	@go build -o bin/server

vite:
	@npm run build


hotreload-run:
	@ls * | entr -r make

hotreload-dev:
	@ls

clean:
	@go mod tidy