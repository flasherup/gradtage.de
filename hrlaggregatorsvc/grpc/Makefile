INCLUDE = -I=. -I=${GO_PATH}/src -I=${GO_PATH}/src/github.com/gogo/protobuf/protobuf

regenerate:
	protoc $(INCLUDE) --go_out=plugins=grpc:. hrlaggregatorsvc.proto

