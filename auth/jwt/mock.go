package jwt

import (
	"time"

	"github.com/gofiber/fiber/v2"
	jwtLib "github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"github.com/wisaitas/share-pkg/cache/redis"
)

type MockJwt struct {
	mock.Mock
}

func NewMockJwt() *MockJwt {
	return &MockJwt{}
}

func (m *MockJwt) Generate(claims Claims, secret string) (string, error) {
	args := m.Called(claims, secret)
	return args.String(0), args.Error(1)
}

func (m *MockJwt) Parse(tokenString string, claims jwtLib.Claims, secret string) (jwtLib.Claims, error) {
	args := m.Called(tokenString, claims, secret)
	return args.Get(0).(jwtLib.Claims), args.Error(1)
}

func (m *MockJwt) ExtractTokenFromHeader(c *fiber.Ctx) (string, error) {
	args := m.Called(c)
	return args.String(0), args.Error(1)
}

func (m *MockJwt) ValidateToken(tokenString string, claims jwtLib.Claims, secret string) error {
	args := m.Called(tokenString, claims, secret)
	return args.Error(0)
}

func (m *MockJwt) CreateStandardClaims(id string, expireTime time.Duration) StandardClaims {
	args := m.Called(id, expireTime)
	return args.Get(0).(StandardClaims)
}

func (m *MockJwt) AuthAccessToken(c *fiber.Ctx, redis redis.Redis, jwt Jwt, secret string) error {
	args := m.Called(c, redis, jwt, secret)
	return args.Error(0)
}

func (m *MockJwt) AuthRefreshToken(c *fiber.Ctx, redis redis.Redis, jwt Jwt, secret string) error {
	args := m.Called(c, redis, jwt, secret)
	return args.Error(0)
}

func (m *MockJwt) GenerateToken(data map[string]interface{}, exp int64, secret string) (string, error) {
	args := m.Called(data, exp, secret)
	return args.String(0), args.Error(1)
}
