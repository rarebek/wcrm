package tokens

import (
	// "api-gateway/pkg/logger"
	// "api-gateway/internal/pkg/config
	"api-gateway/internal/pkg/config"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTHandler ...
type JWTHandler struct {
	Sub       string
	Iss       string
	Exp       string
	Iat       string
	Aud       []string
	Role      string
	SigninKey string
	Log       *log.Logger
	Token     string
	Timeot    int
}

type CustomClaims struct {
	*jwt.Token
	Sub  string   `json:"sub"`
	Exp  float64  `json:"exp"`
	Iat  float64  `json:"iat"`
	Aud  []string `json:"aud"`
	Role string   `json:"role"`
}

// GenerateAuthJWT ...
func (jwtHandler *JWTHandler) GenerateAuthJWT() (access, refresh string, err error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
		rtClaims     jwt.MapClaims
	)
	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)
	claims = accessToken.Claims.(jwt.MapClaims)
	claims["sub"] = jwtHandler.Sub
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(jwtHandler.Timeot)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = jwtHandler.Role
	claims["aud"] = jwtHandler.Aud
	access, err = accessToken.SignedString([]byte(config.NewConfig().SigningKey))
	if err != nil {
		log.Println("error generating access token", err)
		return
	}

	rtClaims = refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = jwtHandler.Sub
	refresh, err = refreshToken.SignedString([]byte(jwtHandler.SigninKey))
	if err != nil {
		log.Println("error generating refresh token", err)
		return
	}
	return
}

// ExtractClaims ...
func (jwtHandler *JWTHandler) ExtractClaims() (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	token, err = jwt.Parse(jwtHandler.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtHandler.SigninKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		log.Println("invalid jwt token")
		return nil, err
	}
	return claims, nil
}

// ExtractClaim extracts claims from given token
func ExtractClaim(tokenStr string, signinigKey []byte) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return signinigKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, err
	}
	return claims, nil
}
