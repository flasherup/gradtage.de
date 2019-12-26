INCLUDE = -I=. -I=/Users/valeriyvostrikov/go/src -I=/Users/valeriyvostrikov/go/src/github.com/gogo/protobuf/protobuf

regenerate:
	protoc $(INCLUDE) --go_out=plugins=grpc:. stationssvc.proto

