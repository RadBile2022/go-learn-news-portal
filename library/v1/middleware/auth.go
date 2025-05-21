package middleware

import (
	"context"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/handling"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/router"
	"net/http"

	"github.com/google/uuid"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return router.RootHandler(func(w http.ResponseWriter, r *http.Request) error {
			ctx := r.Context()

			tokenString, err := parseBearerToken(r)
			if err != nil {
				return err
			}

			claims, err := validateToken(tokenString)
			if err != nil {
				return err
			}

			ctx = AddUserIDToContext(ctx, claims.ID)
			ctx = AddSessionIDToContext(ctx, claims.SessionID)
			ctx = AddRoleToContext(ctx, claims.Role)
			ctx = AddPermissionsToContext(ctx, claims.Permissions)
			ctx = AddInstituteIDToContext(ctx, claims.InstituteID)
			ctx = AddUserTypeToContext(ctx, claims.LoginAs)

			next.ServeHTTP(w, r.WithContext(ctx))
			return nil
		})
	}
}

type contextKey string

const (
	userIDKey       contextKey = "user_id"
	sessionIDKey    contextKey = "session_id"
	roleKey         contextKey = "role"
	permissionsKey  contextKey = "permissions"
	instituteKey    contextKey = "institute"
	userTypeKey     contextKey = "user_type"
	refreshTokenKey contextKey = "refresh_token"
)

func AddSessionIDToContext(ctx context.Context, sessionId string) context.Context {
	return context.WithValue(ctx, sessionIDKey, sessionId)
}

func GetSessionIDFromContext(ctx context.Context, w http.ResponseWriter) (string, error) {
	sessionID, ok := ctx.Value(sessionIDKey).(string)
	if !ok {
		return "", handling.NewHttpError401(nil)
	}
	return sessionID, nil
}

func AddUserIDToContextUint(ctx context.Context, userId uint) context.Context {
	return context.WithValue(ctx, userIDKey, userId)
}

func AddUserIDToContext(ctx context.Context, userId int64) context.Context {
	return context.WithValue(ctx, userIDKey, userId)
}

func GetUserIDFromContext(ctx context.Context) int64 {
	userID, ok := ctx.Value(userIDKey).(int64)
	if !ok {
		return 0
	}

	return userID
}

func GetUserIDFromContextWithErr(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value(userIDKey).(uint)
	if !ok {
		return 0, handling.NewHttpError401(nil)
	}
	return userID, nil
}

func AddRoleToContext(ctx context.Context, role string) context.Context {
	return context.WithValue(ctx, roleKey, role)
}

func AddPermissionsToContext(ctx context.Context, permissions uint) context.Context {
	return context.WithValue(ctx, permissionsKey, permissions)
}

func GetRoleFromContext(ctx context.Context) (string, error) {
	role, ok := ctx.Value(roleKey).(string)
	if !ok {
		return "", handling.NewHttpError401(nil)
	}
	return role, nil
}

func GetPermissionsFromContext(ctx context.Context) (uint, error) {
	permissions, ok := ctx.Value(permissionsKey).(uint)
	if !ok {
		return 0, handling.NewHttpError401(nil)
	}
	return permissions, nil
}

// uuid.MustParse("3592d442-495d-466b-8a70-f2daa1404e57")
func AddInstituteIDToContext(ctx context.Context, instituteID uuid.UUID) context.Context {
	return context.WithValue(ctx, instituteKey, instituteID)
}

func GetInstituteIDFromContext(ctx context.Context, w http.ResponseWriter) (uuid.UUID, error) {
	instituteID, ok := ctx.Value(instituteKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, handling.NewHttpError401(nil)
	}
	return instituteID, nil
}

func AddUserTypeToContext(ctx context.Context, userType string) context.Context {
	return context.WithValue(ctx, userTypeKey, userType)
}

func GetUserTypeFromContext(ctx context.Context, w http.ResponseWriter) (string, error) {
	userType, ok := ctx.Value(userTypeKey).(string)
	if !ok {
		return "", handling.NewHttpError401(nil)
	}
	return userType, nil
}

func AddRefreshTokenToContext(ctx context.Context, refreshToken string) context.Context {
	return context.WithValue(ctx, refreshTokenKey, refreshToken)
}

func GetRefreshTokenFromContext(ctx context.Context, w http.ResponseWriter) (string, error) {
	refreshToken, ok := ctx.Value(refreshTokenKey).(string)
	if !ok {
		return "", handling.NewHttpError401(nil)
	}
	return refreshToken, nil
}
