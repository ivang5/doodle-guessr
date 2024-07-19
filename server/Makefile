serverbuild:
	docker build -t doodle-guessr-server .

serverinit:
	docker run --name doodle-guessr-server -p 3000:3000 doodle-guessr-server:latest

serverstart:
	docker start doodle-guessr-server

serverpurge: serverstop serverremove serverremoveimg

serverstop:
	docker stop doodle-guessr-server

serverremove:
	docker rm doodle-guessr-server

serverremoveimg:
	docker rmi doodle-guessr-server


.PHONY: serverbuild serverinit serverstart serverpurge serverstop serverremove