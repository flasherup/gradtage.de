INCLUDE = -I=. -I=/Users/valeriyvostrikov/go/src -I=/Users/valeriyvostrikov/go/src/github.com/gogo/protobuf/protobuf -I=/Users/valeriyvostrikov/go/src/github.com/

regenerate:
	protoc $(INCLUDE) --go_out=plugins=grpc:. hourlysvc.proto

