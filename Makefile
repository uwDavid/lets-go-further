server: 
	go run ./cmd/api

c-m: 
	migrate create -seq -ext sql -dir db/migrations -seq $(name)

db: 
	docker compose up -d

#migrate create -seq -ext=.sql -dir=./migrations create_movie_table
m-up: 
	migrate -path db/migrations/ -database "postgres://root:secret@localhost:5432/greenlight?sslmode=disable" up

m-down: 
	migrate -path db/migrations/ -database "postgres://root:secret@localhost:5432/greenlight?sslmode=disable" down

BODY='{"title":"Black Panther","year":2018,"runtime":"134 mins","genres":["sci-fi","action","adventure"]}'
test-post:
	curl -i -d $(BODY) localhost:4000/v1/movies

USER='{"name": "Alice Smith", "email": "alice@example.com", "password": "pa55word"}'
USER2='{"name": "", "email": "bob@invalid.", "password": "pass"}'
USER3='{"name": "Bob Jones", "email": "bob@example.com", "password": "pa55word"}'
test-user-register:
	curl -i -d $(USER3) localhost:4000/v1/users
	
test-put: 
	curl -X PUT -d $(BODY) localhost:4000/v1/movies/2

test-show:
	curl localhost:4000/v1/movies/4

test-delete:
	curl -X DELETE localhost:4000/v1/movies/7

test-patch: 
	curl -X PATCH -d '{"year":1985}' localhost:4000/v1/movies/4

test-list:
	curl localhost:4000/v1/movies

