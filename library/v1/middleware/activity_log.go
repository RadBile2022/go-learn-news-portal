package middleware

//
//import (
//	"context"
//	"log/slog"
//	"time"
//
//	activitylog "github.com/dev-digitalproject/go-simpp-auth-core/internal/pkg/activity_log"
//)
//
//func LogUserActivity(ctx context.Context, logger *activitylog.ActivityLogger, eventType string, metadata interface{}) {
//	authUserID, err := GetUserIDFromContext(ctx)
//	if err != nil {
//		slog.Error("Error getting user id from context", slog.Any("error", err))
//		return
//	}
//	proxy, err := GetRequestInfoFromContext(ctx)
//	if err != nil {
//		slog.Error("Error getting request info from context", slog.Any("error", err))
//		return
//	}
//	logger.Log(activitylog.ActivityLogEntry{
//		UserID:           authUserID,
//		ActionType:       eventType,
//		Metadata:         metadata,
//		Timestamp:        time.Now(),
//		RequestIP:        proxy.IPAddr,
//		RequestUserAgent: proxy.UserAgent,
//		RequestEndpoint:  proxy.Endpoint,
//	})
//}
