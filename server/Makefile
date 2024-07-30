postgresinit:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgresstart:
	docker start postgres15

postgresstop:
	docker stop postgres15

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root doodle-guessr

dropdb:
	docker exec -it postgres15 dropdb doodle-guessr

migrateup:
	migrate -path migrations/ -database "postgres://root:password@localhost:5432/doodle-guessr?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations/ -database "postgres://root:password@localhost:5432/doodle-guessr?sslmode=disable" -verbose down
	
.PHONY: postgresinit postgresstart postgresstop postgres createdb dropdb migrateup migratedown