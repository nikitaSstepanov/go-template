package jwt

import (
	"app/internal/entity"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gosuit/e"
)

type Jwt struct {
	audience   []string
	issuer     string
	accessKey  string
	refreshKey string
}

func New(options *JwtOptions) *Jwt {
	return &Jwt{
		audience:   options.Audience,
		issuer:     options.Issuer,
		accessKey:  options.AccessKey,
		refreshKey: options.RefreshKey,
	}
}

func (j *Jwt) ValidateToken(jwtString string, isRefresh bool) (*Claims, e.Error) {
	keyFunc := j.funcAccess

	if isRefresh {
		keyFunc = j.funcRefresh
	}

	token, err := jwt.ParseWithClaims(jwtString, &Claims{}, keyFunc)
	if err != nil {
		return nil, unauthErr.WithErr(err)
	}

	if !token.Valid {
		return nil, unauthErr.WithErr(err)
	}

	return token.Claims.(*Claims), nil
}

func (j *Jwt) GenerateToken(user *entity.User, expires time.Duration, isRefresh bool) (string, e.Error) {
	c := Claims{
		user.Id,
		user.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
			Issuer:    j.issuer,
			Audience:  j.audience,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	key := j.accessKey

	if isRefresh {
		key = j.refreshKey
	}

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", e.InternalErr.WithErr(err)
	}

	return tokenString, nil
}

func (j *Jwt) funcAccess(token *jwt.Token) (interface{}, error) {
	return []byte(j.accessKey), nil
}

func (j *Jwt) funcRefresh(token *jwt.Token) (interface{}, error) {
	return []byte(j.refreshKey), nil
}
