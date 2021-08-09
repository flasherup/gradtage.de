INCLUDE = -I=. -I=${GO_PATH}/src -I=${GO_PATH}/src/github.com/gogo/protobuf/protobuf

regenerate:
	protoc $(INCLUDE) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false dailysvc.proto

