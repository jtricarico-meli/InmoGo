package utils

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"time"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

type JWTMaker struct {
	secretKey string
}

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

type Maker interface {
	CreateToken(user string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

const minSecretKeySize = 32

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

func Authenticate(r *http.Request, jwt Maker) error {
	authHeader := r.Header.Get("authorization")
	if authHeader == "" {
		return errors.New("traffic not authenticated")
	}
	var credentials string
	parts := strings.Split(authHeader, " ")
	if parts[0] == "Bearer" {
		credentials = parts[1]
	}

	token, err := jwt.VerifyToken(credentials)
	if err != nil {
		return err
	}

	return token.Valid()
}
