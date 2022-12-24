setup:
				go mod tidy
				go install

generate:
						go get github.com/99designs/gqlgen
						go run github.com/99designs/gqlgen generate

start: 
						go run server.go