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

func TestBannerSvc_Execute(t *testing.T) {
	userTest := "user-test"
	t.Run("Should_Pass_When_Get_Banner", func(t *testing.T) {
		mockRepo := new(mocks.BannerRepo)

		mockRepoRes := port.Banner{
			BannerId:    "123",
			Title:       "title",
			Description: "desc",
			Image:       "image",
		}
		mockRepo.On("GetByUserId", userTest).Return(mockRepoRes, nil)

		svc := service.NewBannerSvc(mockRepo)

		res, err := svc.Execute(domain.BannerSvcReq{
			UserId: userTest,
		})

		expected := domain.BannerSvcRes{
			BannerId:    "123",
			Title:       "title",
			Description: "desc",
			Image:       "image",
		}

		assert.NoError(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_No_UserID", func(t *testing.T) {
		mockRepo := new(mocks.BannerRepo)
		svc := service.NewBannerSvc(mockRepo)

		res, err := svc.Execute(domain.BannerSvcReq{
			UserId: "",
		})

		expected := domain.BannerSvcRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Should_Error_When_Get_Banner", func(t *testing.T) {
		mockRepo := new(mocks.BannerRepo)

		mockRepoRes := port.Banner{}
		mockRepo.On("GetByUserId", userTest).Return(mockRepoRes, liberror.ErrorInternalServerError("", ""))

		svc := service.NewBannerSvc(mockRepo)

		res, err := svc.Execute(domain.BannerSvcReq{
			UserId: userTest,
		})

		expected := domain.BannerSvcRes{}

		assert.Error(t, err)
		assert.Equal(t, expected, res)
		mockRepo.AssertExpectations(t)
	})
}
