package authentication

import (
	"context"
	"github.com/k2rth1k/quick-trade/utilities/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var logger = log.InitZapLog()

func ServerInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// Skip authorize when GetJWT is requested
	if info.FullMethod != "/quick_trade.QuickTrade/Login" && info.FullMethod != "/quick_trade.QuickTrade/CreateUser" {

		valid := ValidateToken(ctx)
		if !valid {
			return nil, status.Error(codes.Unauthenticated, "failed to authenticate")
		}

		ad, err := ExtractTokenMetadata(ctx)
		if err != nil {
			logger.Errorw("failed to extract metadata from context", "error", err)
			return nil, status.Error(codes.Internal, "internal error")
		}
		ctx = context.WithValue(ctx, AccessDetails{}, ad)
	}
	// Calls the handler
	// validateToken function validates the token

	h, err := handler(ctx, req)

	return h, err
}
