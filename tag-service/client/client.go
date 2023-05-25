package main

import (
	"context"
	"log"

	pb "tag-service/proto"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	//grpc.WithBlock() =>發起撥號連線時會阻塞等待連接完成
	clientConn, _ := GetClientConn(ctx, "localhost:8004", []grpc.DialOption{grpc.WithBlock()})
	defer clientConn.Close()

	//初始化指定PRC Proto Service的用戶端實例物件
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	//發起指定RPC方法的呼叫
	resp, _ := tagServiceClient.GetTagList(ctx, &pb.GetTagListRequest{})

	log.Printf("resp: %+v", resp)

}

// grpc.DialContext建立指定目標的用戶端連接(非加密模式)
func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	//grpc.WithInsecure禁用此clientConn的傳輸安全性驗證
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}
