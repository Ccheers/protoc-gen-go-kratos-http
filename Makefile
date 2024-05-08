.PHONY: proto
proto:
	 find ./khttp -name '*.proto' |xargs -I {} protoc --proto_path=. --go_out=paths=source_relative:. {}

.PHONY: install
install:fmt
	go install .

.PHONY: example
example: install
	find ./example -name '*.proto' |xargs -I {} protoc --proto_path=. --proto_path=./third_party --go_out=paths=source_relative:. --go-kratos-http_out=paths=source_relative:. {}

.PHONY: fmt
fmt:
	gofumpt -w -l .
	goimports -w -l .