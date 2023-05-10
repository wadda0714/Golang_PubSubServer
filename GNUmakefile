gen-design:
	@echo "Generating design..."
	@goa gen github.com/wadda0714/Golang_PubSubServer/server/design
gen-example:
	@echo "Generating example..."
	@rm  pub_sub_server.go
	@goa example github.com/wadda0714/Golang_PubSubServer/server/design
build-server:
	@echo "Building server..."
	@go build -o ./bin/server ./cmd/pub_sub_server
