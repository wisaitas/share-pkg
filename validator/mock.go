package validator

import (
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

type MockValidator struct {
	mock.Mock
}

func NewMockValidator() *MockValidator {
	return &MockValidator{}
}

func (m *MockValidator) ValidateStruct(data any) error {
	args := m.Called(data)
	return args.Error(0)
}

func (m *MockValidator) ValidateCommonJSONBody(c *fiber.Ctx, data any) error {
	args := m.Called(c, data)
	return args.Error(0)
}

func (m *MockValidator) ValidateCommonParam(c *fiber.Ctx, data any) error {
	args := m.Called(c, data)
	return args.Error(0)
}

func (m *MockValidator) ValidateCommonFormBody(c *fiber.Ctx, data any) error {
	args := m.Called(c, data)
	return args.Error(0)
}

func (m *MockValidator) ValidateImageFiles(files []*multipart.FileHeader, maxFileSize int64) error {
	args := m.Called(files, maxFileSize)
	return args.Error(0)
}

func (m *MockValidator) ValidateCommonQueryParam(c *fiber.Ctx, data any) error {
	args := m.Called(c, data)
	return args.Error(0)
}
