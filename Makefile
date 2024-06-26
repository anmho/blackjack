
all: deps generate blackjack


deps:

generate:
	@protoc --go_out=./gen --go_opt=paths=source_relative \
             --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
             ./proto/blackjack/*.proto

blackjack:
	@go build -o ./bin/blackjack main.go

dev:
	@./bin/blackjack