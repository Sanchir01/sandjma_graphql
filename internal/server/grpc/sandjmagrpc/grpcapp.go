package sandjmagrpc

import (
	"context"
	"fmt"
	sandjmav1 "github.com/Sanchir01/protos_files_job/gen/go/auth"
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

func NewGrpcAuth(ctx context.Context, api sandjmav1.AuthClient, addr string, retries int, log *slog.Logger) (*Client, error) {
	const op = "grpc.NewSandjma"
	retryOpts := []grpcretry.CallOption{
		grpcretry.WithCodes(codes.NotFound, codes.Aborted, codes.DeadlineExceeded),
		grpcretry.WithPerRetryTimeout(time.Second * 10),
		grpcretry.WithMax(uint(retries)),
	}
	logOpts := []grpclog.Option{
		grpclog.WithLogOnEvents(grpclog.PayloadReceived, grpclog.PayloadSent),
	}
	cc, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			grpclog.UnaryClientInterceptor(InterceptorLogger(log), logOpts...),
			grpcretry.UnaryClientInterceptor(retryOpts...)),
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Client{
		api: sandjmav1.NewAuthClient(cc),
	}, nil
}

func (g *Client) IsUserPhone(ctx context.Context, userId string) (*sandjmav1.LoginResponse, error) {
	const op = "grpc.IsUserPhone"
	resp, err := g.api.Login(ctx, &sandjmav1.LoginRequest{
		Phone:    "1234",
		Password: "test",
	})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return resp, nil
}

func InterceptorLogger(l *slog.Logger) grpclog.Logger {
	return grpclog.LoggerFunc(func(ctx context.Context, lvl grpclog.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
