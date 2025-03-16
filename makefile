
run:
	go build -o ./bin/app ./cmd/app
	./bin/app

e2e-test:
	go test ./tests -count=1

unit-test:
	go test ./internal/... -count=1
	
gen-swagger-docs:
	swag init -g cmd/app/main.go

docker:
	docker compose --file ./deploy/docker/docker-compose.yaml --env-file .env up --detach
