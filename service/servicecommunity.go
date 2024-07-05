package servicecommunity

import (
	pb "GreenThumb/Community-Service/genproto/community"
	"GreenThumb/Community-Service/storage/postgres"
	"context"
	"database/sql"
	"log"
)

type CommunityServer struct {
	pb.UnimplementedCommunityServiceServer
	db        *sql.DB
	community *postgres.Community
}

func NewCommunityServer(db *sql.DB, community *postgres.Community) *CommunityServer {
	return &CommunityServer{db: db, community: community}
}

func (s *CommunityServer) CreateCommunity(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	community, err := s.community.CreateCommunity(req)
	if err != nil {
		return nil, err
	}
	return community, nil
}

func (s *CommunityServer) GetCommunityByID(ctx context.Context, req *pb.GetCommunityRequest) (*pb.GetCommunityResponse, error) {
	getcommunity, err := s.community.GetCommunity(req)
	if err != nil {
		log.Fatal("Xatolik?")
		return nil, err
	}
	return getcommunity, nil
}

func (s *CommunityServer) UpdateCommunityByID(ctx context.Context, req *pb.UpdateCommunityRequest) (*pb.UpdateCommunityResponse, error) {
	updatecommunity, err := s.community.UpdateCommunity(req)
	if err != nil {
		return nil, err
	}
	return updatecommunity, nil
}

func (s *CommunityServer) DeleteCommunityByID(ctx context.Context, req *pb.DeleteCommunityRequest) (*pb.DeleteCommunityResponse, error) {
	deletecommunity, err := s.community.DeleteCommunity(req)
	if err != nil {
		return nil, err
	}
	return deletecommunity, nil
}

func (s *CommunityServer) GetAllCommunities(ctx context.Context, req *pb.GetCommunitiesRequest) (*pb.GetCommunitiesResponse, error) {
	getcommunity, err := s.community.GetAllCommunities(req)
	if err != nil {
		return nil, err
	}
	return getcommunity, nil
}

func (s *CommunityServer) JoinCommunityMember(ctx context.Context, req *pb.JoinMemberRequest) (*pb.JoinMemberResponse, error) {
	joincommunity, err := s.community.CreateCommunityMember(req)
	if err != nil {
		return nil, err
	}
	return joincommunity, nil
}

func (s *CommunityServer) LeftCommunityMember(ctx context.Context, req *pb.LeftMemberRequest) (*pb.LeftMemberResponse, error) {
	leftcommunity, err := s.community.DeleteCommunityMember(req)
	if err != nil {
		return nil, err
	}
	return leftcommunity, nil
}

func (s *CommunityServer) JoinCommunityEvent(ctx context.Context, req *pb.JoinEventRequest) (*pb.JoinEventResponse, error) {
	joinevent, err := s.community.CreateEvent(req)
	if err != nil {
		return nil, err
	}

	return joinevent, nil
}

func (s *CommunityServer) GetEventByID(ctx context.Context, req *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	getevent, err := s.community.GetEvent(req)
	if err != nil {
		return nil, err
	}

	return getevent, nil
}

func (s *CommunityServer) JoinCommunityForum(ctx context.Context, req *pb.JoinForumRequest) (*pb.JoinForumResponse, error) {
	joinforum, err := s.community.CreateForum(req)
	if err != nil {
		return nil, err
	}
	return joinforum, nil
}

func (s *CommunityServer) GetForumByID(ctx context.Context, req *pb.GetForumRequest) (*pb.GetForumResponse, error) {
	getforumbyid, err := s.community.GetForumBYID(req)
	if err != nil {
		return nil, err
	}
	return getforumbyid, nil
}

func (s *CommunityServer) AddCommentForum(ctx context.Context, req *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	addcommunityforum, err := s.community.CreateComment(req)
	if err != nil {
		return nil, err
	}
	return addcommunityforum, nil
}

func (s *CommunityServer) GetForumCommentByID(ctx context.Context, req *pb.GetForumCommentRequest) (*pb.GetForumCommentResponse, error) {
	getcommentforum, err := s.community.GetCommentByUserID(req)
	if err != nil {
		return nil, err
	}
	return getcommentforum, nil
}
