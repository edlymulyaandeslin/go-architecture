package service

import (
	"clean-code-app-laundry/config"
	"clean-code-app-laundry/model"
	"clean-code-app-laundry/model/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(payload model.User) (dto.LoginResponseDto, error)
	VerifyToken(token string) (jwt.MapClaims, error)
}

type jwtService struct {
	config config.SecurityConfig
}

func (j *jwtService) GenerateToken(payload model.User) (dto.LoginResponseDto, error) {
	claims := dto.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    j.config.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.config.Durasi * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserId: payload.Id,
		Role:   payload.Role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(j.config.Key))
	if err != nil {
		return dto.LoginResponseDto{}, err
	}
	return dto.LoginResponseDto{Token: ss}, nil
}

func (j *jwtService) VerifyToken(token string) (jwt.MapClaims, error) {
	panic("unimplemented")
}

func NewJwtService(cg config.SecurityConfig) JwtService {
	return &jwtService{config: cg}
}
