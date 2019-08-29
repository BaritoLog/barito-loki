package loki

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func onStoreErrorGrpc(err error) error {
	return status.Errorf(codes.FailedPrecondition, err.Error())
}
