
run:
	go build -o ./bin/app ./cmd/app
	./bin/app

docker-up:
	docker-compose --file ./deploy/docker/docker-compose.yaml --env-file .env up --detach

k8s-setup-dbs:
	kubectl apply -f ./deploy/k8s/network/namespace/postgres.yaml
	kubectl apply -f ./deploy/k8s/network/namespace/redis.yaml

	kubectl apply -f ./deploy/k8s/config/secret/postgres.yaml
	kubectl apply -f ./deploy/k8s/config/secret/redis.yaml

	helm install postgresql bitnami/postgresql-ha -f ./deploy/k8s/helm/postgresql/values.yaml -n postgres
	helm install redis bitnami/redis -f ./deploy/k8s/helm/redis/values.yaml -n redis

k8s-setup-app:
	kubectl apply -f ./deploy/k8s/deployment/app.yaml

	kubectl apply -f ./deploy/k8s/service/app.yaml
	
k8s-kind-up:
	kind create cluster --config ./deploy/k8s/kind/kind.yaml --image=kindest/node:v1.31.0 --name=app

k8s-kind-down:
	kind delete cluster --name=app
