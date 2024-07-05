package postgres

import (
	pb "GreenThumb/Community-Service/genproto/community"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Community struct {
	DB *sql.DB
}

func NewCommunity(db *sql.DB) *Community {
	return &Community{
		DB: db,
	}
}

func (c *Community) CreateCommunity(req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id := uuid.NewString()
	query := `INSERT INTO communities (id,name, description, location) VALUES ($1, $2, $3,$4)`
	_, err := c.DB.Exec(query, id, req.Name, req.Description, req.Location)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Message: "Successfully created"}, nil
}

func (c *Community) GetCommunity(req *pb.GetCommunityRequest) (*pb.GetCommunityResponse, error) {
	query := `SELECT name,description,location FROM communities WHERE community_id = $1`
	row := c.DB.QueryRow(query, req.CommunityID)

	resp := &pb.GetCommunityResponse{}

	err := row.Scan(&resp.Name, &resp.Description, &resp.Location)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Community) UpdateCommunity(req *pb.UpdateCommunityRequest) (*pb.UpdateCommunityResponse, error) {
	query := `UPDATE community SET name = $1, description = $2, location = $3 where community_id = $4`
	_, err := c.DB.Exec(query, req.Name, req.Description, req.Location, req.CommunityID)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCommunityResponse{Success: true}, nil

}

func (c *Community) DeleteCommunity(req *pb.DeleteCommunityRequest) (*pb.DeleteCommunityResponse, error) {

	DeletedAt := time.Now().Unix()

	query := `update from community set deleted_at = $1 where community_id = $2`
	_, err := c.DB.Exec(query, DeletedAt, req.CommunityID)
	if err != nil {
		return &pb.DeleteCommunityResponse{Success: false}, err
	}
	return &pb.DeleteCommunityResponse{Success: true}, nil
}

func (c *Community) GetAllCommunities(req *pb.GetCommunitiesRequest) (*pb.GetCommunitiesResponse, error) {
	query := `SELECT * FROM community where len(description) > 40`
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comms []*pb.Community
	for rows.Next() {
		comm := &pb.Community{}
		err = rows.Scan(&comm.Name, &comm.Description, &comm.Location)
		if err != nil {
			return nil, err
		}
		comms = append(comms, comm)
	}
	return &pb.GetCommunitiesResponse{Communities: comms}, nil
}

func (c *Community) CreateCommunityMember(req *pb.JoinMemberRequest) (*pb.JoinMemberResponse, error) {
	query := `insert into community_members(community_id, use_id) values ($1, $2)`
	_, err := c.DB.Exec(query, req.CommunityID, req.UserID)
	if err != nil {
		return &pb.JoinMemberResponse{
			Success: false,
		}, err
	}
	return &pb.JoinMemberResponse{Success: true}, nil
}

func (c *Community) DeleteCommunityMember(req *pb.LeftMemberRequest) (*pb.LeftMemberResponse, error) {
	query := `delete from community_members where community_id = $1`
	_, err := c.DB.Exec(query, req.CommunityID)
	if err != nil {
		return &pb.LeftMemberResponse{Success: false}, err
	}
	return &pb.LeftMemberResponse{Success: true}, nil

}

func (c *Community) CreateEvent(req *pb.JoinEventRequest) (*pb.JoinEventResponse, error) {
	id := uuid.NewString()
	newtime := time.Now()
	query := `insert into events (id,community_id, name , description, type,location,start_time,end_time) values ($1, $2, $3, $4, $5,$6,$7,$8)`
	_, err := c.DB.Exec(query, id, req.CommunityID, req.Name, req.Description, req.Type, req.Location, newtime, newtime)
	if err != nil {
		return &pb.JoinEventResponse{Success: false}, nil
	}
	return &pb.JoinEventResponse{Success: true}, nil
}

func (c *Community) GetEvent(req *pb.GetEventsRequest) (*pb.GetEventsResponse, error) {
	query := `SELECT * FROM events WHERE event_id = $1`
	row := c.DB.QueryRow(query, req.EventID)

	event := &pb.GetEventsResponse{}

	err := row.Scan(&event.Event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (c *Community) CreateForum(req *pb.JoinForumRequest) (*pb.JoinForumResponse, error) {
	id := uuid.NewString()
	newtime := time.Now()
	query := `insert into forum_posts (id,community_id, user_id ,title, content,created_at) values ($1, $2, $3, $4, $5, $6)`
	_, err := c.DB.Exec(query, id, req.CommunityID, req.UserID, req.Title, req.Content, newtime)
	if err != nil {
		return &pb.JoinForumResponse{Success: false}, nil
	}
	return &pb.JoinForumResponse{Success: true}, nil
}

func (c *Community) GetForumBYID(req *pb.GetForumRequest) (*pb.GetForumResponse, error) {
	query := `select * from forums where forum_id = $1`
	row := c.DB.QueryRow(query, req.ForumID)

	var forum *pb.GetForumResponse

	err := row.Scan(&forum.Forum.Id, &forum.Forum.CommunityID, &forum.Forum.Title, &forum.Forum.Content)
	if err != nil {
		return nil, nil
	}
	return forum, nil
}

func (c *Community) CreateComment(req *pb.AddCommentRequest) (*pb.AddCommentResponse, error) {
	query := `insert into Comments (user_id, forum_id, content) values($1,$2,$3)`
	_, err := c.DB.Exec(query, req.Request.UserID, req.Request.ForumID, req.Request.Content)
	if err != nil {
		return &pb.AddCommentResponse{Success: false}, err
	}
	return &pb.AddCommentResponse{Success: true}, nil
}

func (c *Community) GetCommentByUserID(req *pb.GetForumCommentRequest) (*pb.GetForumCommentResponse, error) {
	query := `select * from comments where user_id = $1`
	row := c.DB.QueryRow(query, req.UserID)

	var com *pb.GetForumCommentResponse

	err := row.Scan(&com.Comment.ForumID, &com.Comment.Content)
	if err != nil {
		return nil, err
	}
	return com, nil
}
