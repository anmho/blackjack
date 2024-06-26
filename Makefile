
all: deps generate blackjack


deps:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

generate:
	@protoc --go_out=./gen --go_opt=paths=source_relative \
             --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
             ./proto/blackjack/*.proto

blackjack:
	@go build -o ./bin/blackjack main.go


clean:
	@rm ./bin/*

dev:
	@./bin/blackjack