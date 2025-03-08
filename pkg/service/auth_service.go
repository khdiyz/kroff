package service

import (
	"errors"
	"kroff/config"
	"kroff/pkg/models"
	"kroff/pkg/repository"
	"kroff/utils/response"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
)

type authService struct {
	repos *repository.Repository
	cfg   *config.Config
}

func NewAuthService(repos *repository.Repository, cfg *config.Config) *authService {
	return &authService{
		repos: repos,
		cfg:   cfg,
	}
}

type jwtCustomClaim struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (s *authService) createToken(username string, expiresAt time.Time) (string, error) {
	claims := &jwtCustomClaim{
		Username: username,
		Role:     "admin",
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return "", response.ServiceError(err, codes.Internal)
	}

	return token, nil
}

func (s *authService) ParseToken(token string) (*jwtCustomClaim, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &jwtCustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(s.cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := jwtToken.Claims.(*jwtCustomClaim)
	if !ok {
		return nil, errors.New("token claims are not of type *jwtCustomClaim")
	}

	return claims, nil
}

func (s *authService) Login(request models.LoginRequest) (string, error) {
	if request.Username != s.cfg.AdminUsername || request.Password != s.cfg.AdminPassword {
		return "", response.ServiceError(errors.New("invalid username or password"), codes.Unauthenticated)
	}

	return s.createToken(request.Username, time.Now().Add(time.Duration(s.cfg.JWTAccessExpirationHours)*time.Hour))
}
