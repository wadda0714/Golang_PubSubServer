gen-design:
	@echo "Generating design..."
	@goa gen github.com/wadda0714/Golang_PubSubServer/server/design -o ./server 
gen-example:
	#注意！自分でで実装した部分は消えます　
	@echo "Generating example..."
	@rm  ./server/pub_sub_server.go
	@goa example github.com/wadda0714/Golang_PubSubServer/server/design -o ./server
build-server:
	@echo "Building server..."
	@go build -o ./bin/server ./server/cmd/pub_sub_server
