package server

import (
	"context"
	"encoding/json"
	"tag-service/pkg/bapi"
	"tag-service/pkg/errorcode"
	pb "tag-service/proto"
)

type TagServer struct {
	pb.UnimplementedTagServiceServer
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")

	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errorcode.TogRPCError(errorcode.ErrorGetTagListFail)
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil, errorcode.TogRPCError(errorcode.Fail)
	}

	return &tagList, nil
}
