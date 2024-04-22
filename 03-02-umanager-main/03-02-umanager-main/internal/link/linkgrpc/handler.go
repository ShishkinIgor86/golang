package linkgrpc

import (
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/pb"
)

var _ pb.LinkServiceServer = (*Handler)(nil)

func New(linksRepository linksRepository, timeout time.Duration) *Handler {
	return &Handler{linksRepository: linksRepository, timeout: timeout}
}

type Handler struct {
	pb.UnimplementedLinkServiceServer
	linksRepository linksRepository
	timeout         time.Duration
}

func (h Handler) GetLinkByUserID(ctx context.Context, id *pb.GetLinksByUserId) (*pb.ListLinkResponse, error) {
	links, err := h.linksRepository.GetLinksByUserID(ctx, id.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, "error getting links by user ID")
	}
	return &pb.ListLinkResponse{Links: links}, nil
}

func (h Handler) CreateLink(ctx context.Context, request *pb.CreateLinkRequest) (*pb.Empty, error) {
	err := h.linksRepository.CreateLink(ctx, &Link{
		UserID: request.Link.UserId,
		URL:    request.Link.Url,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "error creating link")
	}
	return &pb.Empty{}, nil
}

func (h Handler) GetLink(ctx context.Context, request *pb.GetLinkRequest) (*pb.Link, error) {
	link, err := h.linksRepository.GetLinkByID(ctx, request.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "link not found")
	}

	return &pb.Link{
		Id:     link.ID,
		UserId: link.UserID,
		Url:    link.URL,
	}, nil
}

func (h Handler) UpdateLink(ctx context.Context, request *pb.UpdateLinkRequest) (*pb.Empty, error) {
	err := h.linksRepository.UpdateLink(ctx, &Link{
		ID:     request.Link.Id,
		UserID: request.Link.UserId,
		URL:    request.Link.Url,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, "error updating link")
	}

	return &pb.Empty{}, nil
}

func (h Handler) DeleteLink(ctx context.Context, request *pb.DeleteLinkRequest) (*pb.Empty, error) {
	err := h.linksRepository.DeleteLink(ctx, request.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "error deleting link")
	}
	return &pb.Empty{}, nil
}

func (h Handler) ListLinks(ctx context.Context, request *pb.Empty) (*pb.ListLinkResponse, error) {
	links, err := h.linksRepository.GetAllLinks(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "error getting all links")
	}
	var pbLinks []*pb.Link
	for _, link := range links {
		pbLinks = append(pbLinks, &pb.Link{
			Id:     link.ID,
			UserId: link.UserID,
			Url:    link.URL,
		})
	}

	return &pb.ListLinkResponse{Links: pbLinks}, nil
}
