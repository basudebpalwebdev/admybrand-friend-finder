build-postgres:
	docker run \
		--name postgres13_3 \
		-e POSTGRES_USER=basu \
		-e POSTGRES_PASSWORD=Basudeb@2021 \
		-p 5432:5432 \
		-d postgres:13.3-alpine3.14

start-postgres:
	docker start postgres13_3

stop-postgres:
	docker stop postgres13_3

destroy-postgres: stop-postgres 
	docker rm postgres13_3

create-db:
	docker exec -it postgres13_3 createdb --username=basu --owner=basu admybrand_friend_finder

drop-db:
	docker exec -it postgres13_3 dropdb --username=basu admybrand_friend_finder

create-migration:
	migrate create -ext sql -dir db/migration -seq $(migration_name)

migrate-up:
	migrate -path db/migration -database "postgresql://basu:Basudeb@2021@localhost:5432/admybrand_friend_finder?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://basu:Basudeb@2021@localhost:5432/admybrand_friend_finder?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test: test-setup
	go test -v -cover ./...

test-setup: drop-db create-db migrate-up sqlc

mock:
	mockgen \
		-destination db/mock/store.go \
		--build_flags=--mod=mod \
		github.com/basudebpalwebdev/admybrand-friend-finder/db/sqlc \
		Store

serve: start-postgres
	air

.PHONY: build-postgres destroy-postgres start-postgres stop-postgres create-db drop-db migrate-up migrate-down sqlc test mock serve
