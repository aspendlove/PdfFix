.PHONY: dev start build build-linux docker-build docker-rebuild docker-dev

dev:
	./node_modules/.bin/tailwindcss -i ./static/css/styles.css -o ./static/css/dist.css --watch & \
	air

start:
	./node_modules/.bin/tailwindcss -i ./static/css/styles.css -o ./static/css/dist.css & \
	GIN_MODE=release go run .

build:
	npm install
	./node_modules/.bin/tailwindcss -i ./static/css/styles.css -o ./static/css/dist.css
	go build -o ./bin/pdf-fix .

docker-build: build
	docker build -t pdf-fix:latest .

docker-rebuild: build
	docker build --no-cache -t pdf-fix:latest .

docker-run:
	docker run -p 8080:8080 --name pdf-fix --rm pdf-fix:latest

docker-dev: docker-build docker-run
