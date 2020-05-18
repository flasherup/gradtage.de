package common

import (
	googlerpc "google.golang.org/grpc"
)

func OpenGRPCConnection(host string) (*googlerpc.ClientConn, error) {
	gc, err := googlerpc.Dial(host, googlerpc.WithInsecure())
	return gc, err
}