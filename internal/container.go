package internal

import (
	"github.com/dewadg/typix-server/internal/repositories"
	"github.com/dewadg/typix-server/internal/services"
)

type Container struct {
	MongoUserRepository repositories.UserRepository
	MongoRaceRepository repositories.RaceRepository
	HTTPWordRepository  repositories.WordRepository

	UserService services.UserService
	RaceService services.RaceService
}
