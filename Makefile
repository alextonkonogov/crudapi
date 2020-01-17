APP=main

docker_build:
	rm -f ${APP} || true
	go build main.go
	docker build -t alextonkonogov/crudapi .

docker_run:
	docker run -p 5151:5151 alextonkonogov/crudapi

docker_push:
	docker push alextonkonogov/crudapi
