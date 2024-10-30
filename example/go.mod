module github.com/Ccheers/protoc-gen-go-kratos-http/example

go 1.21

require (
	github.com/Ccheers/protoc-gen-go-kratos-http v0.0.4
	github.com/go-kratos/kratos/v2 v2.7.3
	google.golang.org/genproto/googleapis/api v0.0.0-20240506185236-b8a5c65736ae
	google.golang.org/protobuf v1.34.1
)

require (
	github.com/go-kratos/aegis v0.2.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240429193739-8cf5692501f6 // indirect
	google.golang.org/grpc v1.63.2 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/Ccheers/protoc-gen-go-kratos-http => ../
)