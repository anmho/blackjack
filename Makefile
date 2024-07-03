
all: deps generate blackjack

.PHONY: deps
deps:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: generate
generate:
	@protoc --go_out=./gen --go_opt=paths=source_relative \
			 --go-grpc_out=./gen --go-grpc_opt=paths=source_relative \
			 ./proto/blackjack/*.proto

.PHONY: blackjack
blackjack: deps generate
	@go build -o ./bin/blackjack main.go


clean:
	@
.PHONY: cleanrm ./bin/*

.PHONY: dev
dev: blackjack
	@./bin/blackjack