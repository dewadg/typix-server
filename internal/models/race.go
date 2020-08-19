package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Race struct {
	ID             primitive.ObjectID `bson:"_id"`
	Words          []string
	ParticipantIDs []primitive.ObjectID
	Participants   []User
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time
}

type RaceResult struct {
	ID            primitive.ObjectID `bson:"_id"`
	RaceID        primitive.ObjectID
	ParticipantID primitive.ObjectID
	Participant   User
	Position      int
	WPM           int
	Correct       int
	Wrong         int
}
