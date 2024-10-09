package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	e "github.com/nikitaSstepanov/tools/error"
)

type Jwt struct {
	audience   []string
	issuer     string
	accessKey  string
	refreshKey string
}

func NewJwt(options *JwtOptions) *Jwt {
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
		return nil, unauthErr
	}

	if !token.Valid {
		return nil, unauthErr
	}

	return token.Claims.(*Claims), nil
}

func (j *Jwt) GenerateToken(id uint64, role string, expires time.Duration, isRefresh bool) (string, e.Error) {
	c := Claims{
		id,
		role,
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
		return "", internalErr
	}

	return tokenString, nil
}

func (j *Jwt) funcAccess(token *jwt.Token) (interface{}, error) {
	return []byte(j.accessKey), nil
}

func (j *Jwt) funcRefresh(token *jwt.Token) (interface{}, error) {
	return []byte(j.refreshKey), nil
}
