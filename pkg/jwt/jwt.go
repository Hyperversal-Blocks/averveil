package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const PublicKeyUser = "___public__key_user"

var (
	ErrJWTExpired = errors.New("token expired")
)

type Token struct {
	Token string `json:"token"`
}

type DecodeJWTClaims struct {
	ExpiredAt float64 `json:"expiredAt"`
	UID       string  `json:"uid"`
}

// Wrapper wraps the signing key and the issuer
type wrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JWT interface {
	GenerateToken(publicKey, secret string) (string, error)
	ValidateToken(token, secret string) (*Claim, error)
}

func New(secret, issuer string, expiry int64) JWT {
	return &wrapper{
		SecretKey:       secret,
		Issuer:          issuer,
		ExpirationHours: expiry,
	}
}

// Claim adds email as a claim to the token
type Claim struct {
	PublicKey string
	jwt.RegisteredClaims
}

// GenerateToken generates a jwt token
func (w *wrapper) GenerateToken(publicKey, secret string) (string, error) {
	claims := &Claim{
		PublicKey: publicKey,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours))},
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(w.SecretKey))
	if err != nil {
		return "", fmt.Errorf("code signing token: %w", err)
	}

	return signedToken, nil
}

var ErrInvalidToken = errors.New("invalid token")

// ValidateToken validates the jwt token
func (w *wrapper) ValidateToken(token, secret string) (*Claim, error) {
	parsed, err := jwt.ParseWithClaims(token, &Claim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := parsed.Claims.(*Claim)
	if !ok {
		return nil, fmt.Errorf("failed to parse claims: %w", ErrInvalidToken)
	}

	if claims.ExpiresAt.Unix() < time.Now().Local().Unix() {
		return nil, ErrJWTExpired
	}

	return claims, nil
}
