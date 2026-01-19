dev:
	./node_modules/.bin/tailwindcss -i ./static/css/styles.css -o ./static/css/dist.css --watch & \
	air

start:
	./node_modules/.bin/tailwindcss -i ./static/css/styles.css -o ./static/css/dist.css & \
	export GIN_MODE=release && go run .
