package authgrpc

import (
	"context"
	"fmt"
	sandjmav1 "github.com/Sanchir01/protos_files_job/pkg/gen/golang/auth"
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
	api sandjmav1.AuthClient
	log *slog.Logger
}

func NewGrpcAuth(ctx context.Context, addr string, retries int, log *slog.Logger) (*Client, error) {
	const op = "grpc.NewSandjma"
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

	return &Client{
		api: sandjmav1.NewAuthClient(cc),
		log: log,
	}, nil
}

func (g *Client) IsUserPhone(ctx context.Context, userId string) (*sandjmav1.LoginResponse, error) {
	const op = "grpc.IsUserPhone"
	resp, err := g.api.Login(ctx, &sandjmav1.LoginRequest{
		Phone:    userId,
		Password: "test",
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return resp, nil
}
