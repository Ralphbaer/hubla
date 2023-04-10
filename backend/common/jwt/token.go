package jwt

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/Ralphbaer/hubla/backend/common/hlog"
	commonHTTP "github.com/Ralphbaer/hubla/backend/common/net/http"
)

// Auth is a struct containing public and private keys for
// JWT access token authentication.
type Auth struct {
	AccessTokenPublicKey  string
	AccessTokenPrivateKey string
}

// NewJWTAuth creates a new JWTAuth instance with the given access token public key.
// Returns a pointer to the Auth struct.
func NewJWTAuth(accessTokenPublicKey string) *Auth {
	return &Auth{
		AccessTokenPublicKey: accessTokenPublicKey,
	}
}

// CreateAccessToken generates a JWT access token with the given payload.
// Returns the token as a string and an error if there's any issue.
func (j *Auth) CreateAccessToken(payload interface{}) (string, error) {
	return j.createToken(60*time.Minute, payload, j.AccessTokenPrivateKey)
}

// ValidateToken verifies the provided JWT access token using the public key.
// Returns an error if the validation fails.
func (j *Auth) ValidateToken(token string) error {
	return j.validateToken(token, j.AccessTokenPublicKey)
}

func (j *Auth) createToken(ttl time.Duration, payload interface{}, privateKey string) (string, error) {
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

func getTokenHeader(r *http.Request) string {
	splitToken := strings.Split(r.Header.Get("Authorization"), "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}

// Protect protects any endpoint using JWT tokens
func (j *Auth) Protect(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := hlog.NewLoggerFromContext(r.Context())
		l.Debug("JWTMiddleware:Protect")

		l.Debug("Read token from header")
		tokenString := getTokenHeader(r)

		if len(tokenString) == 0 {
			commonHTTP.Unauthorized(w, "Must provider a token")
			return
		}
		if err := j.ValidateToken(tokenString); err != nil {
			commonHTTP.Unauthorized(w, err.Error())
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (j *Auth) validateToken(token string, publicKey string) error {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	if _, ok := parsedToken.Claims.(jwt.MapClaims); !ok || !parsedToken.Valid {
		return fmt.Errorf("validate: invalid token")
	}

	return nil
}
