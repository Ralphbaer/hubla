package jwt

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

// PartnerMongoRepository represents a MongoDB implementation of PartnerRepository interface
type JWTAuth struct {
	AccessTokenExpiresIn   time.Duration
	AccessTokenPrivateKey  string
	RefreshTokenExpiresIn  time.Duration
	RefreshTokenPrivateKey string
}

// NewJWT creates a instance of Config
func NewJWTAuth(jwt *JWTAuth) *JWTAuth {
	return &JWTAuth{
		AccessTokenExpiresIn:   jwt.AccessTokenExpiresIn,
		AccessTokenPrivateKey:  jwt.AccessTokenPrivateKey,
		RefreshTokenExpiresIn:  jwt.RefreshTokenExpiresIn,
		RefreshTokenPrivateKey: jwt.RefreshTokenPrivateKey,
	}
}

func (j *JWTAuth) CreateAccessToken(payload interface{}) (string, error) {
	return j.createToken(j.AccessTokenExpiresIn, payload, j.AccessTokenPrivateKey)
}

func (j *JWTAuth) CreateRefreshToken(payload interface{}) (string, error) {
	return j.createToken(j.RefreshTokenExpiresIn, payload, j.RefreshTokenPrivateKey)
}

func (j *JWTAuth) ValidateToken(token string) (interface{}, error) {
	return j.validateToken(token, j.RefreshTokenPrivateKey)
}

func (j *JWTAuth) createToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = payload
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

func (j *JWTAuth) validateToken(token string, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return "", fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims["sub"], nil
}