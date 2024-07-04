package mocking

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type AuthMiddlewareMock struct {
	mock.Mock
}

func (m *AuthMiddlewareMock) CheckToken(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
