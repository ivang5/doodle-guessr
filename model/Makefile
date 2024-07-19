modelservicebuild:
	docker build -t model-service .

modelserviceinit:
	docker run --name model-service -p 3000:3000 model-service:latest

modelservicestart:
	docker start model-service

modelservicepurge: serverstop serverremove serverremoveimg

modelservicestop:
	docker stop model-service

modelserviceremove:
	docker rm model-service

modelserviceremoveimg:
	docker rmi model-service


.PHONY: modelservicebuild modelserviceinit modelservicestart modelservicepurge modelservicestop modelserviceremove