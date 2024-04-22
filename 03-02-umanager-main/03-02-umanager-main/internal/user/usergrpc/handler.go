package usergrpc

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/pb"
)

var _ pb.UserServiceServer = (*Handler)(nil)

func New(usersRepository usersRepository, timeout time.Duration) *Handler {
	return &Handler{usersRepository: usersRepository, timeout: timeout}
}

type Handler struct {
	pb.UnimplementedUserServiceServer
	usersRepository usersRepository
	timeout         time.Duration
}

func (h Handler) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.Empty, error) {
	user := &pb.User{
		Id:       in.GetId(),
		Username: in.GetUsername(),
		Email:    in.GetEmail(),
	}

	err := h.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error creating user: %v", err)
	}

	return &pb.Empty{}, nil
}

func (h Handler) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	user, err := h.usersRepository.GetUser(ctx, in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user with ID %s not found", in.GetId())
	}
	return user, nil
}

func (h Handler) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	user := &pb.User{
		Id:       in.GetId(),
		Username: in.GetUsername(),
		Email:    in.GetEmail(),
	}
	err := h.usersRepository.UpdateUser(ctx, user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error updating user: %v", err)
	}
	return &pb.Empty{}, nil
}

func (h Handler) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.Empty, error) {
	err := h.usersRepository.DeleteUser(ctx, in.GetId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error deleting user: %v", err)
	}
	return &pb.Empty{}, nil
}

func (h Handler) ListUsers(ctx context.Context, in *pb.Empty) (*pb.ListUsersResponse, error) {
	users, err := h.usersRepository.ListUsers(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error listing users: %v", err)
	}
	return &pb.ListUsersResponse{Users: users}, nil
}
