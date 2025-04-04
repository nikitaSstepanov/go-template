
run:
	go build -o ./bin/app ./cmd/app
	./bin/app

docker:
	docker compose up -d

e2e-test:
	go test ./tests -count=1

unit-test:
	go test ./internal/... -count=1
	
gen-mocks:
	mockgen -source=internal/usecase/pkg/user/type.go -destination=internal/usecase/pkg/user/mocks/user/mocks.go
	mockgen -source=internal/usecase/pkg/auth/type.go -destination=internal/usecase/pkg/auth/mocks/auth/mocks.go
	mockgen -source=internal/controller/http/v1/pkg/account/type.go -destination=internal/controller/http/v1/pkg/account/mocks/account/mocks.go
	mockgen -source=internal/controller/http/v1/pkg/auth/type.go -destination=internal/controller/http/v1/pkg/auth/mocks/auth/mocks.go

gen-swagger-docs:
	swag init -g cmd/app/main.go
