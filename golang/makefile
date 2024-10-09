
run:
	go build -o ./bin/app ./cmd/app
	./bin/app

docker-up:
	docker-compose --file ./deploy/docker/docker-compose.yaml --env-file .env up --detach

k8s-apply:
	kubectl apply -f ./deploy/k8s/deployments/postgres.yaml
	kubectl apply -f ./deploy/k8s/deployments/redis.yaml
	kubectl apply -f ./deploy/k8s/deployments/app.yaml

k8s-delete:
	kubectl delete -f ./deploy/k8s/deployments/postgres.yaml
	kubectl delete -f ./deploy/k8s/deployments/redis.yaml
	kubectl delete -f ./deploy/k8s/deployments/app.yaml
	
k8s-kind-up:
	kind create cluster --config ./deploy/k8s/kind/kind.yaml --image=kindest/node:v1.31.0 --name=app
