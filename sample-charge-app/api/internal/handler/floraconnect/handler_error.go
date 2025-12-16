package floraconnecthandler

import (
	"errors"

	"github.com/flora/api/internal/usecase"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleError(err error) error {
	switch {
	case errors.Is(err, usecase.ErrInternal):
		return status.Errorf(codes.Internal, "internal server error")
	case errors.Is(err, usecase.ErrNotFound):
		return status.Errorf(codes.NotFound, err.Error())
	case errors.Is(err, usecase.ErrInvalidRequest):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, usecase.ErrInvalidArgument):
		return status.Errorf(codes.InvalidArgument, err.Error())
	case errors.Is(err, usecase.ErrInvalidState):
		return status.Errorf(codes.FailedPrecondition, err.Error())
	default:
		return status.Errorf(codes.Internal, "unknown error")
	}
}
