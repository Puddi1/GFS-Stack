run: vite build
	@./bin/server

dev: build
	@./bin/server

build:
	@go build -o bin/server

vite:
	@npm run build

hotreload:
	@ls * | entr -r make

clean:
	@go mod tidy