package service_test

import (
	"auth-service/internal/core/domain"
	"auth-service/internal/core/port"
	"auth-service/internal/core/port/mocks"
	"auth-service/internal/core/service"
	"errors"
	"testing"

	liberror "github.com/basputtipong/library/error"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockJWTGenerator struct{}

func (m *mockJWTGenerator) Generate(userID string) (string, error) {
	return "token.test", nil
}

type mockJWTGeneratorError struct{}

func (m *mockJWTGeneratorError) Generate(userID string) (string, error) {
	return "", errors.New("mock generate error")
}

func TestLoginService_Execute(t *testing.T) {
	userTest := "user-test"
	passcode := "111111"

	t.Run("Should_Pass_When_Login", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepo)
		mockJWT := &mockJWTGenerator{}

		mockRepo.On("UpdatePasscodeByUserId", mock.MatchedBy(func(req port.UsersRepoReq) bool {
			return req.UserId == userTest && req.Passcode != ""
		})).Return(nil)

		svc := service.NewLoginSvc(mockRepo, mockJWT)
		res, err := svc.Execute(domain.LoginSvcReq{
			UserId:   userTest,
			Passcode: passcode,
		})

		expected := "token.test"
		assert.NoError(t, err)
		assert.Equal(t, expected, res.AuthToken)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_GenerateJWT_Error", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepo)
		mockJWT := &mockJWTGeneratorError{}

		mockRepo.On("UpdatePasscodeByUserId", mock.Anything).Return(nil)

		svc := service.NewLoginSvc(mockRepo, mockJWT)
		_, err := svc.Execute(domain.LoginSvcReq{
			UserId:   userTest,
			Passcode: passcode,
		})

		assert.Error(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Update_Passcode", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepo)
		mockJWT := &mockJWTGenerator{}

		mockRepo.On("UpdatePasscodeByUserId", mock.MatchedBy(func(req port.UsersRepoReq) bool {
			return req.UserId == userTest && req.Passcode != ""
		})).Return(liberror.ErrorInternalServerError("", ""))

		svc := service.NewLoginSvc(mockRepo, mockJWT)
		res, err := svc.Execute(domain.LoginSvcReq{
			UserId:   userTest,
			Passcode: passcode,
		})

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_No_UserId", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepo)
		mockJWT := &mockJWTGenerator{}

		svc := service.NewLoginSvc(mockRepo, mockJWT)
		res, err := svc.Execute(domain.LoginSvcReq{
			UserId:   "",
			Passcode: "",
		})

		assert.Error(t, err)
		assert.Nil(t, res)
		mockRepo.AssertExpectations(t)
	})
}
