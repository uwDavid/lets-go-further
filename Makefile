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

USER='{"name": "Test User", "email": "test@example.com", "password": "pa55word"}'
USER2='{"name": "", "email": "bob@invalid.", "password": "pass"}'
USER3='{"name": "Charlie Charles", "email": "charles@example.com", "password": "pa55word"}'
test-user-register:
	curl -i -d $(USER3) localhost:4000/v1/users
	
TOKEN='{"token": "Z7BUOADCNEWLSVMD4JIJTXO7YI"}'
test-activation: 
	curl -X PUT -d $(TOKEN) localhost:4000/v1/users/activated

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

mailserver:
	docker run -d -p 1025:1025 -p 8025:8025 mailhog/mailhog

# vendor: tidy and vendor dependencies
.PHONY: vendor
vendor: 
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# audit: tidy, vendor dependencies and format, vet, test all code
# include vendor as a prerequisite
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	# staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

# ============
# BUILD
# ============
current_time = $(shell date --iso-8601=seconds)
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

.PHONY: build/api
build/api:
	@echo 'Buliding cmd/api...'
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/api ./cmd/api
