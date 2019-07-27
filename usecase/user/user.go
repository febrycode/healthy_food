package ucuser

import (
	"context"
	"time"

	"github.com/fsetiawan29/healthy_food/domain/image"
	"github.com/fsetiawan29/healthy_food/domain/user"
	"github.com/fsetiawan29/healthy_food/models"
	"github.com/fsetiawan29/healthy_food/util"
)

type Usecase struct {
	userRepository  user.Repository
	contextTimeout  time.Duration
	imageRepository image.Repository
}

func NewUserUsecase(u user.Repository, timeout time.Duration, i image.Repository) user.Usecase {
	return &Usecase{
		userRepository:  u,
		contextTimeout:  timeout,
		imageRepository: i,
	}
}

func (uc *Usecase) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func (uc *Usecase) GetUserByUserID(ctx context.Context, userID int64) (models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userRepository.GetByUserID(ctx, userID)
	if err != nil {
		return models.User{}, err
	}

	return res, nil
}

func (uc *Usecase) CreateUser(ctx context.Context, userData *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userData.CreatedAt = util.GetTimeNow()

	err := uc.userRepository.CreateUser(ctx, userData)
	if err != nil {
		return err
	}

	return nil
}

func (uc *Usecase) UpdateUser(ctx context.Context, userData *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userData.UpdatedAt = util.GetTimeNow()

	err := uc.userRepository.UpdateUser(ctx, userData)
	if err != nil {
		return err
	}

	if userData.AvatarURL == "" {
		return nil
	}

	imageData, err := uc.imageRepository.GetImageByName(ctx, userData.AvatarURL)
	if err != nil {
		return err
	}

	if imageData.ID <= 0 {
		return nil
	}

	err = uc.imageRepository.UpdateImage(ctx, &models.Image{
		ID:            imageData.ID,
		ReferenceType: 1,
		ReferenceID:   userData.ID,
		Name:          imageData.Name,
		UpdatedAt:     util.GetTimeNow(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (uc *Usecase) GetListUser(ctx context.Context) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.userRepository.GetListUser(ctx)
	if err != nil {
		return []models.User{}, err
	}

	return res, nil
}
