setup:
				go mod tidy
				go install

generate:
						go get github.com/99designs/gqlgen
						go run github.com/99designs/gqlgen generate

start: 
						go run server.go

build:
				go build

dbcreate:
				migrate -path "postgres/migrations" -database "$(DB_URL)" up

dbdrop:
				migrate -path "postgres/migrations" -database "$(DB_URL)" drop