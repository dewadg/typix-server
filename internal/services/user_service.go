package services

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/dewadg/typix-server/internal/models"
	"github.com/dewadg/typix-server/internal/repositories"
)

type userService struct {
	mongoUserRepository repositories.UserRepository
}

func NewUserService(mongoUserRepository repositories.UserRepository) UserService {
	return &userService{
		mongoUserRepository: mongoUserRepository,
	}
}

func (svc *userService) Find(ctx context.Context, id primitive.ObjectID) (models.User, error) {
	return svc.mongoUserRepository.Find(ctx, id)
}
