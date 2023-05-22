gen-design:
	@echo "Generating design..."
	@goa gen github.com/wadda0714/Golang_PubSubServer/server/design -o ./server 
gen-example:
	#注意！自分で実装した部分は消えます　
	@echo "Generating example..."
	@cp ./server/pub_sub_server.go ./backup/pub_sub_server.go
	@rm  ./server/pub_sub_server.go
	@goa example github.com/wadda0714/Golang_PubSubServer/server/design -o ./server
build-server:
	@echo "Building server..."
	@go build -o ./bin/server ./server/cmd/pub_sub_server

build-server-all:
	@make gen-design
	@make gen-example
	@make build-server

psql:
	@psql -h localhost -p 5430 -U root

sqlboiler:
	@sqlboiler psql --config ./sqlboiler.toml --pkgname userdb --no-hooks --struct-tag-casing camel --output ./pkg/userdb --no-tests --wipe
