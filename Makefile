run:
	./bin/air

generate:
	go run github.com/arsmn/fastgql
	go run generate.go
	