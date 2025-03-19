package middleware_core

//
//import (
//	"context"
//	"net/http"
//	"strings"
//
//	"github.com/dgrijalva/jwt-go"
//	"github.com/google/uuid"
//	"github.com/spf13/viper"
//)
//
//type payloadType int
//
//const (
//	PayloadTypeAccess payloadType = iota
//	PayloadTypeRefresh
//)
//
//type Payload struct {
//	ID          uint        `json:"id"`
//	Role        string      `json:"role"`
//	Permissions uint        `json:"permissions"`
//	SessionID   string      `json:"sid"`
//	InstituteID uuid.UUID   `json:"institute_id"`
//	LoginAs     string      `json:"login_as"`
//	Type        payloadType `json:"type"`
//	jwt.StandardClaims
//}
//
//type contextKey string
//
//const (
//	authUserKey contextKey = "authenticated_user"
//	proxykey    contextKey = "proxy_info"
//)
//
//func Authentication(next http.Handler) http.Handler {
//	fn := func(w http.ResponseWriter, r *http.Request) {
//		bearerToken, err := parseBearerToken(r)
//		if err != nil {
//			response.Error(err, w)
//			return
//		}
//
//		claims, err := validateToken(bearerToken)
//		if err != nil {
//			response.Error(err, w)
//			return
//		}
//
//		proxy := entity.ProxyRequestInfo{
//			Authorization: r.Header.Get("authorization"),
//			UserAgent:     r.UserAgent(),
//			RemoteIP:      r.RemoteAddr,
//			Endpoint:      r.URL.Path,
//		}
//
//		authService := spesial.NewAuthService(&proxy)
//
//		user := &entity.User{}
//
//		hasReadUserPerm := permission.NewPermission(claims.Permissions).HasAll(permission.ReadUser)
//
//		if hasReadUserPerm {
//			user, err = authService.FindUserByID(claims.ID)
//		} else {
//			err = authService.GetMe(user)
//		}
//
//		if err != nil {
//			response.Error(err, w)
//			return
//		}
//
//		user.Permissions = claims.Permissions
//		user.InstituteID = claims.InstituteID.String()
//		user.InstituteUID = claims.InstituteID
//		user.Role = claims.Role
//		user.Type = claims.LoginAs
//
//		ctx := AddUserToContext(r.Context(), user)
//		ctx = AddProxyInfoToContext(ctx, proxy)
//
//		next.ServeHTTP(w, r.WithContext(ctx))
//	}
//	return http.HandlerFunc(fn)
//}
//
//func getJWTKeyFunc() jwt.Keyfunc {
//	jwtKey := viper.GetString("JWT_KEY")
//	return func(t *jwt.Token) (interface{}, error) {
//		return []byte(jwtKey), nil
//	}
//}
//
//func parseBearerToken(r *http.Request) (string, error) {
//	bearerToken := r.Header.Get("Authorization")
//	if bearerToken == "" {
//		return "", handling.NewHttpError401(nil)
//	}
//
//	var tokenString string
//	if len(strings.Split(bearerToken, " ")) == 2 {
//		tokenString = strings.Split(bearerToken, " ")[1]
//	}
//	return tokenString, nil
//}
//
//func validateToken(tokenStr string) (*Payload, error) {
//	claims := &Payload{}
//	token, err := jwt.ParseWithClaims(tokenStr, claims, getJWTKeyFunc())
//
//	if err != nil || !token.Valid {
//		return nil, handling.NewHttpError401(err)
//	}
//
//	if claims.SessionID == "" || claims.Type != PayloadTypeAccess {
//		return nil, handling.NewHttpError401(err)
//	}
//
//	return claims, nil
//}
//
//func AddProxyInfoToContext(ctx context.Context, proxyInfo entity.ProxyRequestInfo) context.Context {
//	return context.WithValue(ctx, proxykey, proxyInfo)
//}
//
//func AddUserToContext(ctx context.Context, e *entity.User) context.Context {
//	return context.WithValue(ctx, authUserKey, *e)
//}
//
//func GetProxyInfoFromContext(ctx context.Context) (*entity.ProxyRequestInfo, error) {
//	proxy, ok := ctx.Value(proxykey).(entity.ProxyRequestInfo)
//	if !ok {
//		return nil, handling.NewHttpError401(nil)
//	}
//	return &proxy, nil
//}
//
//func GetUserFromContext(ctx context.Context) (*entity.User, error) {
//	user, ok := ctx.Value(authUserKey).(entity.User)
//	if !ok {
//		return nil, handling.NewHttpError401(nil)
//	}
//	return &user, nil
//}
