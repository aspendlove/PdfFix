dev:
	./node_modules/.bin/tailwindcss -i ./static/css/styles.css -o ./static/css/dist.css --watch & \
	air

start:
	./node_modules/.bin/tailwindcss -i ./static/css/styles.css -o ./static/css/dist.css & \
	export GIN_MODE=release && go run .

build:
	./node_modules/.bin/tailwindcss -i ./static/css/styles.css -o ./static/css/dist.css
	export GIN_MODE=release && \
	export CGO_ENABLED=0 && \
	export GOOS=linux && \
	go build -o ./bin .

docker-build: build
	docker build -t pdf-fix:latest .

docker-rebuild: build
	docker build --no-cache -t pdf-fix:latest .

docker-dev: docker-build
	docker run -p 8080:8080 --name pdf-fix --rm pdf-fix:latest
