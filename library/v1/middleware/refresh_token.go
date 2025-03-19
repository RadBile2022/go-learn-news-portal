package middleware

//
//import (
//	"context"
//	"net/http"
//	"slices"
//	"time"
//
//	"github.com/dgrijalva/jwt-go"
//)
//
//func RefreshTokenMiddleware(ctx context.Context, api api.AuthMiddleware) func(http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return router.RootHandler(func(w http.ResponseWriter, r *http.Request) error {
//			tokenString, err := parseBearerToken(r)
//			if err != nil {
//				return err
//			}
//
//			claims := &RefreshPayload{}
//			token, err := jwt.ParseWithClaims(tokenString, claims, getJWTKeyFunc())
//
//			if err != nil || !token.Valid {
//				return handling.NewHttpError401(err)
//			}
//
//			if claims.SessionID == "" || claims.Type != PayloadTypeRefresh {
//				return handling.NewHttpError401(err)
//			}
//
//			user := entity.NewUserFromID(claims.ID)
//			if err := api.FindUserByID(ctx, user); err != nil {
//				return handling.NewHttpError401(err)
//			}
//
//			if user.RememberToken == nil {
//				return handling.NewHttpError401(err)
//			}
//
//			switch user.AccountStatus {
//			case entity.UserAccountStatusSuspended:
//				return handling.NewHttpError401(nil)
//			}
//
//			sessions, err := api.GetSessions(ctx, user.ID)
//			if err != nil {
//				return err
//			}
//
//			sessionIdx := slices.IndexFunc(sessions, func(s *entity.DeviceLog) bool {
//				return s.SessionID == claims.SessionID
//			})
//
//			// if sessionIdx < 0 {
//			// 	return handling.NewHttpError401(nil)
//			// }
//
//			if sessionIdx > -1 {
//				session := sessions[sessionIdx]
//				session.LastActivityAt = time.Now()
//				api.UpdateSession(ctx, session)
//			}
//
//			ctx = AddUserIDToContext(ctx, claims.ID)
//			ctx = AddSessionIDToContext(ctx, claims.SessionID)
//			ctx = AddUserTypeToContext(ctx, claims.LoginAs)
//			ctx = AddRefreshTokenToContext(ctx, tokenString)
//
//			next.ServeHTTP(w, r.WithContext(ctx))
//			return nil
//		})
//	}
//}
