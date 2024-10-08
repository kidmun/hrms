package usecases

import (
	"context"
	"hrms/internal/models"
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type userUsecase struct {
	userRepository models.UserRepository
	contextTimeout  time.Duration
}

func NewUserUsecase(userRepository models.UserRepository, contextTimeout  time.Duration) models.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: contextTimeout,
	}
}
func (uc *userUsecase) CreateUser(ctx context.Context, user *models.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.CreateUser(ctx, user)
}


func (uc *userUsecase) GetUsers(ctx context.Context) ([]*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.GetUsers(ctx)
}


func (uc *userUsecase) GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.GetUserByID(ctx, id)
}

func (uc *userUsecase) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.GetUserByUsername(ctx, username)
}

func (uc *userUsecase) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {	
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.GetUserByEmail(ctx, email)
}

func (uc *userUsecase) UpdateUser(ctx context.Context, user *models.User) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.UpdateUser(ctx, user)
}

func (uc *userUsecase) DeleteUser(ctx context.Context, id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()
	return uc.userRepository.DeleteUser(ctx, id)
}