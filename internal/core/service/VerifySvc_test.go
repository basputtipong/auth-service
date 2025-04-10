package service_test

import (
	"auth-service/internal/core/domain"
	"auth-service/internal/core/port"
	"auth-service/internal/core/port/mocks"
	"auth-service/internal/core/service"
	"testing"

	liberror "github.com/basputtipong/library/error"
	"github.com/stretchr/testify/assert"
)

func TestVerifySvc_Execute(t *testing.T) {
	userTest := "user-test"
	t.Run("Should_Pass_When_Verify", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepo)

		mockRepoRes := port.UsersWithGreeting{
			UserId:   "user-test",
			Name:     "test",
			Passcode: "$2a$10$3fYp0BLAuRodsgMe6amWBeGAvBC3kxPT0TO.CJPrVd7vctjIcyNeO",
			Greeting: "hello",
		}
		mockRepo.On("GetGreetingByUserId", userTest).Return(mockRepoRes, nil)

		svc := service.NewVerifySvc(mockRepo)

		res, err := svc.Execute(domain.VerifySvcReq{
			UserId:   userTest,
			Passcode: "111111",
		})

		expected := domain.VerifySvcRes{
			Name:        "test",
			GreetingMsg: "hello",
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_No_UserId", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepo)
		svc := service.NewVerifySvc(mockRepo)

		res, err := svc.Execute(domain.VerifySvcReq{
			UserId:   "",
			Passcode: "",
		})

		expected := domain.VerifySvcRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Passcode_Not_match", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepo)

		mockRepoRes := port.UsersWithGreeting{
			UserId:   "user-test",
			Name:     "test",
			Passcode: "$2a$10$3fYp0BLAuRodsgMe6amWBeGAvBC3kxPT0TO.CJPrVd7vctjIcyNeO",
			Greeting: "hello",
		}
		mockRepo.On("GetGreetingByUserId", userTest).Return(mockRepoRes, nil)

		svc := service.NewVerifySvc(mockRepo)

		res, err := svc.Execute(domain.VerifySvcReq{
			UserId:   userTest,
			Passcode: "123456",
		})

		expected := domain.VerifySvcRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Fail_Get_Greeting", func(t *testing.T) {
		mockRepo := new(mocks.UsersRepo)

		mockRepoRes := port.UsersWithGreeting{}
		mockRepo.On("GetGreetingByUserId", userTest).Return(mockRepoRes, liberror.ErrorInternalServerError("", ""))

		svc := service.NewVerifySvc(mockRepo)

		res, err := svc.Execute(domain.VerifySvcReq{
			UserId:   userTest,
			Passcode: "111111",
		})

		expected := domain.VerifySvcRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})
}
