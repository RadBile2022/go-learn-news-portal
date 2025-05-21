package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/constant"
	"github.com/RadBile2022/go-learn-news-portal/library/v1/handling"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/spf13/viper"
)

type payloadType int

const (
	PayloadTypeAccess payloadType = iota
	PayloadTypeRefresh
)

type Payload struct {
	ID          int64       `json:"id"`
	Role        string      `json:"role"`
	Permissions uint        `json:"permissions"`
	SessionID   string      `json:"sid"`
	InstituteID uuid.UUID   `json:"institute_id"`
	LoginAs     string      `json:"login_as"`
	Type        payloadType `json:"type"`
	jwt.StandardClaims
}

type RefreshPayload struct {
	Payload
}

type ResetPasswordPayload struct {
	Email string `json:"email"`
	Key   string `json:"key"`
	jwt.StandardClaims
}

func GenerateAccessToken(p *Payload) (string, error) {
	p.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	return createToken(p)
}

func GenerateRefreshToken(p *RefreshPayload) (string, error) {
	p.ExpiresAt = time.Now().Add(24 * time.Hour).Unix()
	return createRefreshToken(p)
}

func createRefreshToken(p *RefreshPayload) (string, error) {
	p.Type = PayloadTypeRefresh
	p.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	jwtKey := viper.GetString("JWT_KEY")
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", handling.NewHttpError(err, http.StatusInternalServerError, "Internal server error", constant.ERR_INTERNAL_SERVER_ERROR)
	}
	return tokenString, nil
}

func createToken(p *Payload) (string, error) {
	p.Type = PayloadTypeAccess
	p.IssuedAt = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, p)
	jwtKey := viper.GetString("JWT_KEY")
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", handling.NewHttpError(err, http.StatusInternalServerError, "Internal server error", constant.ERR_INTERNAL_SERVER_ERROR)
	}
	return tokenString, nil
}

func GenerateResetPasswordToken(email string, key string) (string, error) {
	expTime := time.Now().Add(5 * time.Minute)
	return createResetPasswordToken(email, expTime, key)
}

func ParseResetPasswordToken(tokenString string) (*ResetPasswordPayload, error) {
	claims := &ResetPasswordPayload{}
	jwtKey := []byte(viper.GetString("JWT_KEY"))
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, handling.NewHttpError(err, http.StatusBadRequest, "reset password token is expired", constant.ERR_TOKEN_EXPIRED)
	}
	return claims, nil
}

func createResetPasswordToken(email string, expTime time.Time, key string) (string, error) {
	iat := time.Now()
	payload := &ResetPasswordPayload{
		Email: email,
		Key:   key,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  iat.Unix(),
			ExpiresAt: expTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	jwtKey := viper.GetString("JWT_KEY")
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", handling.NewHttpError(err, http.StatusInternalServerError, "Internal server error", constant.ERR_INTERNAL_SERVER_ERROR)
	}
	return tokenString, nil
}

func getJWTKeyFunc() jwt.Keyfunc {
	jwtKey := viper.GetString("JWT_KEY")
	return func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	}
}

func parseBearerToken(r *http.Request) (string, error) {
	bearerToken := r.Header.Get("Authorization")
	if bearerToken == "" {
		return "", handling.NewHttpError401(nil)
	}

	var tokenString string
	if len(strings.Split(bearerToken, " ")) == 2 {
		tokenString = strings.Split(bearerToken, " ")[1]
	}
	return tokenString, nil
}

func validateToken(tokenStr string) (*Payload, error) {
	claims := &Payload{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, getJWTKeyFunc())

	if err != nil || !token.Valid {
		return nil, handling.NewHttpError401(err)
	}

	//if claims.Role != "integration" && claims.SessionID == "" {
	//	return nil, handling.NewHttpError401(err)
	//}

	//if claims.Type != PayloadTypeAccess {
	//	return nil, handling.NewHttpError401(err)
	//}

	return claims, nil
}

func GenerateIntegrationApiKeyToken(p *Payload) (string, error) {
	p.ExpiresAt = time.Now().Add(24 * time.Hour).Unix() // 1 day expiration
	return createToken(p)
}
