NAME=alextonkonogov/crudapi
APP=main

docker_build:
	rm -f ${APP} || true
	go build main.go
	docker build -t ${NAME} .

docker_run:
	docker run -p 5151:5151 ${NAME}

docker_push:
	docker push ${NAME}
