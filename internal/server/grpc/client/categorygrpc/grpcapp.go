package categorygrpc

import (
	"context"
	"fmt"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/category"
	mwlogger "github.com/Sanchir01/sandjma_graphql/pkg/lib/logger/middleware/logger"
	grpclog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpcretry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"time"
)

type Client struct {
	api sandjmav1.CategoriesClient
	log *slog.Logger
}

func NewGrpcCategory(ctx context.Context, addr string, retries int, log *slog.Logger) (*Client, error) {
	const op = "grpc.NewCategory"
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithPerRetryTimeout(time.Second * 10),
		grpcretry.WithMax(uint(retries)),
	}
	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}
	cc, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(mwlogger.InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...)),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Client{api: sandjmav1.NewCategoriesClient(cc), log: log}, nil
}

func (g *Client) AllCategory(ctx context.Context) (*sandjmav1.GetAllCategoryResponse, error) {
	const op = "grpc.category.AllCategory"
	name := &sandjmav1.Empty{}
	resp, err := g.api.GetAllCategory(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return resp, nil
}
