
run:
	go build -o ./bin/app ./cmd/app
	./bin/app

api-test:
	go test ./tests -count=1
	
gen-swagger-docs:
	swag init -g cmd/app/main.go
