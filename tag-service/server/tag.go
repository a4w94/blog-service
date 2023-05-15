package server

import (
	"context"
	"encoding/json"
	"tag-service/pkg/bapi"
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
		// errocode.ErrorGetTagListFail.WithDetails(err.Error())
		// errorcode.ErrorGetTagListFail.WithDetails(err.Error())
		// return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		//return nil, errcode.TogRPCError(errcode.ErrorGetTagListFail)
	}

	return &tagList, nil
}
