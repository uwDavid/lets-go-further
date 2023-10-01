server: 
	go run ./cmd/api

db: 
	docker compose up -d

#migrate create -seq -ext=.sql -dir=./migrations create_movie_table
m-up: 
	migrate -path db/migrations/ -database "postgres://root:secret@localhost:5432/greenlight?sslmode=disable" up

m-down: 
	migrate -path db/migrations/ -database "postgres://root:secret@localhost:5432/greenlight?sslmode=disable" down