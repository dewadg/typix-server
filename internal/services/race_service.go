package services

import (
	"context"

	"github.com/dewadg/typix-server/internal/models"
	"github.com/dewadg/typix-server/internal/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type raceService struct {
	mongoRaceRepository repositories.RaceRepository
	mongoUserRepository repositories.UserRepository
}

func NewRaceService(mongoRaceRepository repositories.RaceRepository, mongoUserRepository repositories.UserRepository) RaceService {
	return &raceService{
		mongoRaceRepository: mongoRaceRepository,
		mongoUserRepository: mongoUserRepository,
	}
}

func (svc *raceService) Create(ctx context.Context, data models.Race) (models.Race, error) {
	race, err := svc.mongoRaceRepository.Create(ctx, data)
	if err != nil {
		return models.Race{}, err
	}
	return svc.Find(ctx, race.ID)
}

func (svc *raceService) Find(ctx context.Context, id primitive.ObjectID) (models.Race, error) {
	race, err := svc.mongoRaceRepository.Find(ctx, id)
	if err != nil {
		return models.Race{}, err
	}

	participants, err := svc.mongoUserRepository.GetByIDs(ctx, race.ParticipantIDs)
	if err != nil {
		return models.Race{}, err
	}
	race.Participants = participants

	return race, nil
}
